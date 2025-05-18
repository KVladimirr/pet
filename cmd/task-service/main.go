package main

import (
	"log"
	"net"

	"google.golang.org/grpc"
	
	"tasker/internal/task/pb"
	"tasker/internal/task"
)

const grpcPort = ":50051"

func main() {
	lis, err := net.Listen("tcp", grpcPort)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	server := grpc.NewServer()
	taskpb.RegisterTaskServiceServer(server, &task.Server{})

	log.Printf("gRPC server is running on %s", grpcPort)
	if err := server.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}