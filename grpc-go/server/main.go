package main

import (
	"context"
	"fmt"
	"log"
	"net"

	userpb "grpc-go/proto"

	"google.golang.org/grpc"
)

type server struct {
	userpb.UnimplementedUserServiceServer
}

func (s *server) GetUser(ctx context.Context, req *userpb.GetUserRequest) (*userpb.User, error) {
	log.Printf("Received RPC request for User ID: %d", req.GetId())

	return &userpb.User{
		Id:       req.GetId(),
		Name:     "Alex Mercer",
		Email:    "alex@example.com",
		IsActive: true,
	}, nil
}

func main() {
	lis, err := net.Listen("tcp", ":50051")
	if err != nil {
		log.Fatalf("Failed to listen on port 50051: %v", err)
	}

	grpcServer := grpc.NewServer()
	userpb.RegisterUserServiceServer(grpcServer, &server{})

	fmt.Println("gRPC Server running on :50051...")
	if err := grpcServer.Serve(lis); err != nil {
		log.Fatalf("Failed to serve: %v", err)
	}
}
