package controllers

import (
	pb "github.com/sQUARys/TestTaskHezzl/proto"
	"github.com/sQUARys/TestTaskHezzl/services"
	"golang.org/x/net/context"
	"google.golang.org/grpc"
	"google.golang.org/grpc/grpclog"
	"net"
)

type Controller struct {
	pb.UnimplementedUserServiceServer
	Service services.Service
}

func New(service *services.Service) *Controller {
	return &Controller{
		Service: *service,
	}
}

func (ctr *Controller) Start() {
	listener, err := net.Listen("tcp", "localhost:8080")
	if err != nil {
		grpclog.Fatalf("failed to listen: %v", err)
	}
	options := []grpc.ServerOption{}
	server := grpc.NewServer(options...)

	pb.RegisterUserServiceServer(server, &Controller{
		Service: ctr.Service,
	})
	server.Serve(listener)
}

func (ctr *Controller) CreateUser(ctx context.Context, request *pb.CreateUserRequest) (*pb.CreateUserResponse, error) {
	ctr.Service.AddUser(*request.User)

	return &pb.CreateUserResponse{
		User: request.User,
	}, nil

}

func (ctr *Controller) DeleteUser(ctx context.Context, request *pb.DeleteUserRequest) (*pb.DeleteUserResponse, error) {
	ctr.Service.DeleteUser(request.Name)
	return nil, nil
}

func (ctr *Controller) ListUser(ctx context.Context, request *pb.ListUserRequest) (*pb.ListUserResponse, error) {
	Users, err := ctr.Service.Cache.GetUsers()
	if err != nil {
		return nil, err
	}
	return &pb.ListUserResponse{
		User: Users,
	}, nil
}
