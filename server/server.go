package main

import (
	"google.golang.org/grpc"
	"log"
	"net"
	pb "github.com/sQUARys/TestTaskHezzl/proto"
)

type server struct{
}

func main() {
	srv := grpc.NewServer()
	listen, err := net.Listen("tcp", ":8000")
	if err != nil {
		log.Fatalf("Failed to listen: %v", err)
	}
	pb.
}
