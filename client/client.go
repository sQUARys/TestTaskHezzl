package main

import (
	"fmt"
	pb "github.com/sQUARys/TestTaskHezzl/proto"
	"golang.org/x/net/context"
	"google.golang.org/grpc"
	"google.golang.org/grpc/grpclog"
)

func main() {
	opts := []grpc.DialOption{
		grpc.WithInsecure(),
	}

	conn, err := grpc.Dial("localhost:8080", opts...)

	if err != nil {
		grpclog.Fatalf("fail to dial: %v", err)
	}

	defer conn.Close()

	client := pb.NewUserServiceClient(conn)
	ctx := context.Background()

	createRequest := &pb.CreateUserRequest{
		User: &pb.User{
			Id:   2,
			Name: "John",
		},
	}

	//deleteRequest := &pb.DeleteUserRequest{
	//	Id: 2,
	//}

	listRequest := &pb.ListUserRequest{}

	createResp, err := client.CreateUser(ctx, createRequest)
	//deleteResp, err := client.DeleteUser(ctx, deleteRequest)
	listResp, err := client.ListUser(ctx, listRequest)

	if err != nil {
		grpclog.Fatalf("fail to dial: %v", err)
	}

	fmt.Println(createResp, listResp)
}
