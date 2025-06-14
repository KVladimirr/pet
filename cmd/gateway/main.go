package main

import (
	"log"

	"github.com/gin-gonic/gin"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"

	"tasker/internal/gateway"
	pb "tasker/internal/task/pb"
)

func main() {
	conn, err := grpc.NewClient("localhost:50051", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("Failed to connect to grpc: %v", err)
	}
	defer conn.Close()

	client := pb.NewTaskServiceClient(conn)

	router := gin.Default()

	gateway := gateway.New(client)
	
	router.POST("/task", gateway.CreateTaskHandler)
	router.GET("/task", gateway.GetTaskHandler)
	router.GET("/tasks", gateway.ListTasksHandler)
	router.PUT("/task", gateway.UpdateTaskHandler)
	router.DELETE("/task", gateway.DeleteTaskHandler)

	log.Println("REST API Gateway listening on :8080")
    router.Run(":50052")
}