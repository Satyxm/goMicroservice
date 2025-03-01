package main

import (
	"context"
	"fmt"
	"log"

	pb "github.com/satyxm/microservice-example/proto"
	"google.golang.org/grpc"
)

func main() {
	// Connect to the gRPC server
	conn, err := grpc.Dial("localhost:50051", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("Could not connect: %v", err)
	}
	defer conn.Close()

	// Create a new gRPC client
	client := pb.NewAdderClient(conn)

	// Send a request to the gRPC server
	req := &pb.AddRequest{A: 5, B: 3}
	res, err := client.Add(context.Background(), req)
	if err != nil {
		log.Fatalf("Error calling Add: %v", err)
	}

	fmt.Printf("Sum: %d\n", res.Result)
}
