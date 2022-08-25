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

	//Special request to Create New User
	createRequest := &pb.CreateUserRequest{
		User: &pb.User{
			Id:   3,
			Name: "Maks",
		},
	}

	//Special request to Delete User
	deleteRequest := &pb.DeleteUserRequest{
		Name: "Maks",
	}

	//Special request for Get List Users
	listRequest := &pb.ListUserRequest{}

	createResp, err := client.CreateUser(ctx, createRequest)
	listResp, err := client.ListUser(ctx, listRequest)
	deleteResp, err := client.DeleteUser(ctx, deleteRequest)
	listResp1, err := client.ListUser(ctx, listRequest)

	if err != nil {
		grpclog.Fatalf("fail to dial: %v", err)
	}

	fmt.Println("CreteResp : ", createResp)
	fmt.Println("ListResp before delete : ", listResp)
	fmt.Println("DeleteResp: ", deleteResp)
	fmt.Println("ListResp after delete : ", listResp1)

}
