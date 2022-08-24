package main

import (
	"fmt"
	"github.com/sQUARys/TestTaskHezzl/cache"
	pb "github.com/sQUARys/TestTaskHezzl/proto"
	"github.com/sQUARys/TestTaskHezzl/repositories"
	"github.com/sQUARys/TestTaskHezzl/user"
	"golang.org/x/net/context"
	"google.golang.org/grpc"
	"google.golang.org/grpc/grpclog"
	"net"
)

type Server struct {
	pb.UnimplementedUserServiceServer
	repo  Repository
	cache Cache
}

type Repository interface {
	AddUser(user *pb.User) error
	DeleteUser(id int32) error
}

type Cache interface {
	Set(user user.User)
	Get(key string) user.User
	GetAll() []user.User
}

func main() {
	listener, err := net.Listen("tcp", "localhost:8080")
	if err != nil {
		grpclog.Fatalf("failed to listen: %v", err)
	}
	options := []grpc.ServerOption{}
	server := grpc.NewServer(options...)

	db := repositories.New()
	c := cache.New()

	pb.RegisterUserServiceServer(server, &Server{
		repo:  db,
		cache: c,
	})
	server.Serve(listener)
}

func (s *Server) CreateUser(ctx context.Context, request *pb.CreateUserRequest) (*pb.CreateUserResponse, error) {
	s.repo.AddUser(request.User)

	user := user.User{
		Id:   request.User.Id,
		Name: request.User.Name,
	}

	s.cache.Set(user)

	return &pb.CreateUserResponse{
		User: request.User,
	}, nil

}

func (s *Server) DeleteUser(ctx context.Context, request *pb.DeleteUserRequest) (*pb.DeleteUserResponse, error) {
	s.repo.DeleteUser(request.Id)
	fmt.Println("GOT : ", request.Id)

	return nil, nil
}

func (s *Server) ListUser(request *pb.ListUserRequest, server pb.UserService_ListUserServer) error {
	s.cache.Set(user.User{Id: 1, Name: "OLEG"})
	s.cache.Set(user.User{Id: 2, Name: "IGOR"})

	fmt.Println(s.cache.GetAll())
	return nil
}
