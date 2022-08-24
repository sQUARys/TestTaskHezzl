package main

import (
	pb "github.com/sQUARys/TestTaskHezzl/proto"
	"golang.org/x/net/context"
	"google.golang.org/grpc"
	"log"
	"net"
)

type server struct {
}

func main() {
	srv := grpc.NewServer()
	listen, err := net.Listen("tcp", ":8000")
	if err != nil {
		log.Fatalf("Failed to listen: %v", err)
	}
	pb.RegisterUserServiceServer(srv, &server{})

	err = srv.Serve(listen)
	if err != nil {
		log.Fatalf("Failed to listen: %v", err)
	}
}

func (s server) CreateUser(ctx context.Context, request *pb.CreateUserRequest) (*pb.CreateUserResponse, error) {
	//TODO implement me
	panic("Create user")
}

func (s server) DeletePokemon(ctx context.Context, request *pb.DeleteUserRequest) (*pb.DeleteUserResponse, error) {
	//TODO implement me
	panic("Delete")
}

func (s server) ListPokemon(request *pb.ListUserRequest, pokemonServer pb.UserService_ListPokemonServer) error {
	//TODO implement me
	panic("implement me")
}

func (s server) mustEmbedUnimplementedUserServiceServer() {
	//TODO implement me
	panic("implement me")
}
