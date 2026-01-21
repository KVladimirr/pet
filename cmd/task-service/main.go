package main

import (
	"context"
	"log"
	"net"
	"tasker/config"
	"tasker/internal/storage"
	"tasker/internal/storage/migrations"

	"google.golang.org/grpc"

	"tasker/internal/task"
	"tasker/internal/task/pb"
)

const grpcPort = ":50051"

func main() {
	ctx := context.Background()

	if err := migrations.Migration("internal/storage/migrations", config.PG_CONNECTION_STRING); err != nil {
		log.Fatalf("failed to make migrations: %v", err)
	}

	pg, err := storage.NewPostgres(ctx, config.PG_CONNECTION_STRING)
	if err != nil {
		log.Fatalf("failed to connect to pg: %v", err)
	}
	defer pg.Close(ctx)

	lis, err := net.Listen("tcp", grpcPort)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	server := grpc.NewServer()
	db := storage.NewTaskStore(pg)

	taskpb.RegisterTaskServiceServer(server, &task.Server{Store: db})

	if err := server.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
	
	log.Printf("gRPC server is running on %s", grpcPort)
}