package main

import (
	"context"
	"fmt"
	"log"
	"net"

	pb "github.com/satyxm/microservice-example/proto"
	"google.golang.org/grpc"
)

// Define the gRPC server structure
type AdderServer struct {
	pb.UnimplementedAdderServer
}

// Implement the Add function
func (s *AdderServer) Add(ctx context.Context, req *pb.AddRequest) (*pb.AddResponse, error) {
	result := req.A + req.B
	return &pb.AddResponse{Result: result}, nil
}

func main() {
	// Create a TCP listener on port 50051
	listener, err := net.Listen("tcp", ":50051")
	if err != nil {
		log.Fatalf("Failed to listen: %v", err)
	}

	// Create a new gRPC server
	server := grpc.NewServer()
	pb.RegisterAdderServer(server, &AdderServer{})

	fmt.Println("gRPC Server is running on port 50051...")
	if err := server.Serve(listener); err != nil {
		log.Fatalf("Failed to serve: %v", err)
	}
}
