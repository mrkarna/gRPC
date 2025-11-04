package main

import (
	"context"
	"fmt"
	"google.golang.org/grpc"
	pb "grpc-hello/pb"
	"log"
	"net"
)

type helloServer struct {
	pb.UnimplementedHelloServiceServer
}

func (s *helloServer) SayHello(ctx context.Context, req *pb.HelloRequest) (*pb.HelloResponse, error) {
	message := fmt.Sprintf("Hello %s", req.Name)
	return &pb.HelloResponse{Message: message}, nil
}

func main() {
	// Listen on TCP port 8080
	lis, err := net.Listen("tcp", ":8080")
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	// Create a new gRPC server
	grpcServer := grpc.NewServer()
	pb.RegisterHelloServiceServer(grpcServer, &helloServer{})

	log.Println("grpc server is running on port 8080...")

	if err := grpcServer.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
