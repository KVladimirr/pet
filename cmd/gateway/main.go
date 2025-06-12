package main

import (
	"context"
	"log"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"google.golang.org/grpc"
	"google.golang.org/protobuf/types/known/timestamppb"

	pb "tasker/internal/task/pb"
)

func main() {
	conn, err := grpc.NewClient("localhost:50051")
	if err != nil {
		log.Fatalf("Failed to connect to grpc: %v", err)
	}
	defer conn.Close()

	client := pb.NewTaskServiceClient(conn)

	router := gin.Default()

	router.POST("/task", RestCreateTask())
}



/*
Улучшеный код 

cmd/gateway/main.go

package main

import (
    "log"

    "github.com/gin-gonic/gin"
    "google.golang.org/grpc"

    pb "tasker/internal/task/pb"
    "tasker/internal/gateway"
)

func main() {
    conn, err := grpc.Dial("localhost:50051", grpc.WithInsecure())
    if err != nil {
        log.Fatalf("failed to connect to gRPC: %v", err)
    }
    defer conn.Close()

    client := pb.NewTaskServiceClient(conn)
    g := gateway.New(client)

    router := gin.Default()

    router.POST("/tasks", g.CreateTaskHandler)

    // остальные роуты тоже можно вынести и подключить аналогично

    log.Println("REST API Gateway listening on :8080")
    router.Run(":8080")
}
*/





/*
Изначально все должно было быть в одном файле, но лучше разнести, код ниже не актуален по расположению

package main

import (
    "context"
    "log"
    "net/http"
    "time"

    "github.com/gin-gonic/gin"
    "google.golang.org/grpc"
    "google.golang.org/protobuf/types/known/timestamppb"

    pb "tasker/internal/task/pb"
)

func main() {
    conn, err := grpc.Dial("localhost:50051", grpc.WithInsecure())
    if err != nil {
        log.Fatalf("failed to connect to gRPC: %v", err)
    }
    defer conn.Close()

    client := pb.NewTaskServiceClient(conn)

    router := gin.Default()

    router.POST("/tasks", func(c *gin.Context) {
        var req struct {
            Title       string json:"title"
            Description string json:"description"
            Deadline    string json:"deadline" // RFC3339 формат
        }

        if err := c.ShouldBindJSON(&req); err != nil {
            c.JSON(http.StatusBadRequest, gin.H{"error": "invalid request"})
            return
        }

        deadline, err := time.Parse(time.RFC3339, req.Deadline)
        if err != nil {
            c.JSON(http.StatusBadRequest, gin.H{"error": "invalid deadline format"})
            return
        }

        grpcReq := &pb.CreateTaskRequest{
            Title:       req.Title,
            Description: req.Description,
            Deadline:    timestamppb.New(deadline),
        }

        resp, err := client.CreateTask(context.Background(), grpcReq)
        if err != nil {
            c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
            return
        }

        c.JSON(http.StatusOK, resp.Task)
    })

    router.GET("/tasks/:id", func(c *gin.Context) {
        id := c.Param("id")

        resp, err := client.GetTask(context.Background(), &pb.GetTaskRequest{Id: id})
        if err != nil {
            c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
            return
        }

        c.JSON(http.StatusOK, resp.Task)
    })

    router.GET("/tasks", func(c *gin.Context) {
        resp, err := client.ListTasks(context.Background(), &pb.ListTasksRequest{})
        if err != nil {
            c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
            return
        }

        c.JSON(http.StatusOK, resp.Task)
    })

    router.PUT("/tasks/:id", func(c *gin.Context) {
        id := c.Param("id")

        var req struct {
            Status string json:"status"
        }
        if err := c.ShouldBindJSON(&req); err != nil {
            c.JSON(http.StatusBadRequest, gin.H{"error": "invalid request"})
            return
        }

        resp, err := client.UpdateTask(context.Background(), &pb.UpdateTaskRequest{
            Id:     id,
            Status: req.Status,
        })
        if err != nil {
            c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
            return
        }

        c.JSON(http.StatusOK, resp.Task)
    })

    router.DELETE("/tasks/:id", func(c *gin.Context) {
        id := c.Param("id")

        resp, err := client.DeleteTask(context.Background(), &pb.DeleteTaskRequest{Id: id})
        if err != nil || !resp.Success {
            c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to delete task"})
            return
        }

        c.JSON(http.StatusOK, gin.H{"message": "task deleted"})
    })

    log.Println("REST API Gateway listening on :8080")
    router.Run(":8080")
}

*/