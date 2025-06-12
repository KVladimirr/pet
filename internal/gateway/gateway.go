package gateway_service

import ()

type Gateway struct {}

func (g *Gateway) CreateTask() {}


/*
Здесь будет логика обработчиков гет/пост/пут и тд

internal/gateway/gateway.go

package gateway

import (
    "context"
    "net/http"
    "time"

    "github.com/gin-gonic/gin"
    "google.golang.org/protobuf/types/known/timestamppb"

    pb "tasker/internal/task/pb"
)

// Gateway хранит gRPC-клиент
type Gateway struct {
    Client pb.TaskServiceClient
}

// New создает новый Gateway
func New(client pb.TaskServiceClient) *Gateway {
    return &Gateway{Client: client}
}

// Handler для создания задачи
func (g *Gateway) CreateTaskHandler(c *gin.Context) {
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

    resp, err := g.Client.CreateTask(context.Background(), grpcReq)
    if err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
        return
    }

    c.JSON(http.StatusOK, resp.Task)
}


*/