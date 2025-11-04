package main

import (
	"context"
	"google.golang.org/grpc"
	pb "grpc-hello/pb"
	"log"
	"time"
)

func main() {
	// Connect to the gRPC server
	conn, err := grpc.Dial("localhost:8080", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()

	// Create a stub (client)
	client := pb.NewHelloServiceClient(conn)

	// Create a context with timeout
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	// Make the RPC call
	res, err := client.SayHello(ctx, &pb.HelloRequest{Name: "Manish"})
	if err != nil {
		log.Fatalf("could not greet: %v", err)
	}

	log.Printf("Server response: %s", res.Message)
}
