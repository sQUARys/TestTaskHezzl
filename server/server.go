package main

import (
	"fmt"
	pb "github.com/sQUARys/TestTaskHezzl/proto"
	"github.com/sQUARys/TestTaskHezzl/repositories"
	"golang.org/x/net/context"
	"google.golang.org/grpc"
	"google.golang.org/grpc/grpclog"
	"net"
)

type Server struct {
	pb.UnimplementedUserServiceServer
}

func main() {
	listener, err := net.Listen("tcp", "127.0.0.1:8080")
	if err != nil {
		grpclog.Fatalf("failed to listen: %v", err)
	}
	options := []grpc.ServerOption{}
	server := grpc.NewServer(options...)

	pb.RegisterUserServiceServer(server, &Server{})
	server.Serve(listener)
}

func (s *Server) CreateUser(ctx context.Context, request *pb.CreateUserRequest) (*pb.CreateUserResponse, error) {
	//TODO implement me
	fmt.Println("Create : ", request)
	db := repositories.New()
	db.AddUser(request.User)
	return nil, nil

}

func (s *Server) DeleteUser(ctx context.Context, request *pb.DeleteUserRequest) (*pb.DeleteUserResponse, error) {
	//TODO implement me
	fmt.Println("GOT : ", request.Id)

	return nil, nil
}

func (s *Server) ListUser(request *pb.ListUserRequest, server pb.UserService_ListUserServer) error {
	//TODO implement me
	fmt.Println("List")

	return nil
}
