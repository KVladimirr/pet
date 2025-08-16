package main

import (
	"context"
	"fmt"
	"log"
	"tasker/config"
	"tasker/internal/storage"
)

// "net"

// "google.golang.org/grpc"

// "tasker/internal/task/pb"
// "tasker/internal/task"
// "tasker/internal/storage"

const grpcPort = ":50051"

func main() {
	ctx := context.Background()
	// lis, err := net.Listen("tcp", grpcPort)
	// if err != nil {
	// 	log.Fatalf("failed to listen: %v", err)
	// }

	// server := grpc.NewServer()
	// taskStore := storage.NewTaskStore()

	// taskpb.RegisterTaskServiceServer(server, &task.Server{Store: taskStore})

	// log.Printf("gRPC server is running on %s", grpcPort)
	// if err := server.Serve(lis); err != nil {
	// 	log.Fatalf("failed to serve: %v", err)
	// }

	pg, err := storage.NewPostgres(ctx, config.PG_CONNECTION_STRING)
	if err != nil {
		log.Fatalf("failed to connect to pg: %v", err)
	}

	var result string
	pg.Conn.QueryRow(ctx, "SELECT datname FROM pg_database;").Scan(&result)

	fmt.Println(result)

	defer pg.Close(ctx)
}