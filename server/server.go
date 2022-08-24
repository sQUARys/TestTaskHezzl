package main

import (
	"github.com/sQUARys/TestTaskHezzl/cache"
	pb "github.com/sQUARys/TestTaskHezzl/proto"
	"github.com/sQUARys/TestTaskHezzl/repositories"
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
	SetUser(user pb.User)
	GetUser(key string) pb.User
	DeleteUser(key string)
	GetUsers() []*pb.User
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

	s.cache.SetUser(*request.User)

	return &pb.CreateUserResponse{
		User: request.User,
	}, nil

}

func (s *Server) DeleteUser(ctx context.Context, request *pb.DeleteUserRequest) (*pb.DeleteUserResponse, error) {
	s.repo.DeleteUser(request.Id)
	return nil, nil
}

func (s *Server) ListUser(ctx context.Context, request *pb.ListUserRequest) (*pb.ListUserResponse, error) {
	Users := s.cache.GetUsers()
	return &pb.ListUserResponse{
		User: Users,
	}, nil
}
