package main

import (
	"database/sql"
	"fmt"
	"log"
	"net"
	"os"
	"os/signal"
	"syscall"
	"tasker/config"
	"tasker/internal/application/usecase"
	"tasker/internal/infrastracture/repository/postgres"

	"time"

	"tasker/internal/pb"

	api "tasker/internal/api/grpc"

	_ "github.com/lib/pq"
	"google.golang.org/grpc"
)

const grpcPort = ":50051"

func main() {
	// if err := migrations.Migration("internal/storage/migrations", config.PG_CONNECTION_STRING); err != nil {
	// 	log.Fatalf("failed to make migrations: %v", err)
	// }

	db, err := initDB()
	if err != nil {
		log.Fatal("Failed to initialize database", err)
	}
	defer db.Close()

	taskDB := postgres.NewPostgresTaskRepository(db)

	taskService := api.NewTaskService(
		usecase.NewCreateTaskUsecase(taskDB),
		usecase.NewGetAllTasksUsecase(taskDB),
		usecase.NewGetTaskByIDUsecase(taskDB),
		usecase.NewUpdateTaskDeadlineUsecase(taskDB),
		usecase.NewUpdateTaskDescriptionUsecase(taskDB),
		usecase.NewUpdateTaskStatusUsecase(taskDB),
		usecase.NewUpdateTaskTitleUsecase(taskDB),
		usecase.NewDeleteTaskUsecase(taskDB),
	)
	
	lis, err := net.Listen("tcp", grpcPort)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	
	server := grpc.NewServer()
	taskpb.RegisterTaskServiceServer(server, taskService)

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)

	go func() {
		<-quit

		log.Println("Shutting down server")
		server.GracefulStop()
	}()
	
	log.Printf("gRPC server is running on %s", grpcPort)

	if err := server.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}

func initDB() (*sql.DB, error) {
	connStr := config.PG_CONNECTION_STRING

	db, err := sql.Open("postgres", connStr)
	if err != nil {
		return nil, fmt.Errorf("failed to open DB: %v", err)
	}

	db.SetConnMaxLifetime(time.Duration(config.CONN_MAX_LIFETIME) * time.Second)

	log.Println("Database connection established successfully")
	return db, nil
}
