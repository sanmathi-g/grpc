package main

import (
	"context"
	"log"
	"time"

	pb "github.com/sanmathi-g/grpcapi/proto"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func main() {
	// Set up a connection to the gRPC server
	conn, err := grpc.Dial("localhost:50051", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("Failed to connect to the server: %v", err)
	}
	defer conn.Close()

	// Create a gRPC client
	client := pb.NewMovieServiceClient(conn)

	// Set a timeout for the requests
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

}
