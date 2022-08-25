package controllers

import (
	"fmt"
	pb "github.com/sQUARys/TestTaskHezzl/proto"
	"golang.org/x/net/context"
	"google.golang.org/grpc"
	"google.golang.org/grpc/grpclog"
	"net"
)

type Controller struct {
	pb.UnimplementedUserServiceServer
	repo  Repository
	cache Cache
}

type Repository interface {
	AddUser(user *pb.User) error
	DeleteUser(name string) error
}

type Cache interface {
	SetUser(user pb.User) error
	GetUser(key string) (pb.User, error)
	GetUsers() ([]*pb.User, error)
	DeleteUser(name string)
}

func New(repo Repository) *Controller {
	return &Controller{
		repo:
	}
}

func (ctr *Controller) Start() {
	listener, err := net.Listen("tcp", "localhost:8080")
	if err != nil {
		grpclog.Fatalf("failed to listen: %v", err)
	}
	options := []grpc.ServerOption{}
	server := grpc.NewServer(options...)

	pb.RegisterUserServiceServer(server, &Server{
		repo:  db,
		cache: c,
	})
	server.Serve(listener)
}

func (ctr *Controller) CreateUser(ctx context.Context, request *pb.CreateUserRequest) (*pb.CreateUserResponse, error) {
	err := ctr.repo.AddUser(request.User)

	if err != nil {
		return nil, err
	}

	err = ctr.cache.SetUser(*request.User)
	if err != nil {
		return nil, err
	}

	return &pb.CreateUserResponse{
		User: request.User,
	}, nil

}

func (ctr *Controller) DeleteUser(ctx context.Context, request *pb.DeleteUserRequest) (*pb.DeleteUserResponse, error) {
	ctr.repo.DeleteUser(request.Id)
	return nil, nil
}

func (ctr *Controller) ListUser(ctx context.Context, request *pb.ListUserRequest) (*pb.ListUserResponse, error) {
	Users, err := ctr.cache.GetUsers()
	if err != nil {
		return nil, err
	}
	return &pb.ListUserResponse{
		User: Users,
	}, nil
}
