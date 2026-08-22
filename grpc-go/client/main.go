package main

import (
	"context"
	"fmt"
	"log"
	"time"

	userpb "grpc-go/proto"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func main() {
	// establishing connection to the server over HTTP/2
	conn, err := grpc.NewClient("localhost:50051", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("Failed to connect: %v", err)
	}
	defer conn.Close()

	client := userpb.NewUserServiceClient(conn)

	// Context with timeout to avoid hanging network calls
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	// calling the remote RPC method as if it's a local Go function (that's what gRPC does) so this is basically the client stub
	res, err := client.GetUser(ctx, &userpb.GetUserRequest{Id: 101})
	if err != nil {
		log.Fatalf("RPC failed: %v", err)
	}

	fmt.Printf("Response Received:\nID: %d\nName: %s\nEmail: %s\nActive: %t\n",
		res.GetId(), res.GetName(), res.GetEmail(), res.GetIsActive())
}
