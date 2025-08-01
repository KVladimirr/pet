package main

import (
	"log"
	"net"

	"google.golang.org/grpc"
	
	"tasker/internal/task/pb"
	"tasker/internal/task"
	"tasker/internal/storage"
)

const grpcPort = ":50051"

func main() {
	lis, err := net.Listen("tcp", grpcPort)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	server := grpc.NewServer()
	taskStore := storage.NewTaskStore()

	taskpb.RegisterTaskServiceServer(server, &task.Server{Store: taskStore})

	log.Printf("gRPC server is running on %s", grpcPort)
	if err := server.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}