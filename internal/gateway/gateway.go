package gateway

import (
	"context"
	"fmt"
	"net/http"
	pb "tasker/internal/task/pb"

	// "time"

	"github.com/gin-gonic/gin"
	"google.golang.org/protobuf/types/known/timestamppb"
)

type Gateway struct {
	Client pb.TaskServiceClient
}

func New(client pb.TaskServiceClient) *Gateway {
	return &Gateway{Client: client}
}

func (g *Gateway) CreateTaskHandler(c *gin.Context) {
    var req struct {
        Title       string `json:"title"`
        Description string `json:"description"`
        Deadline    *timestamppb.Timestamp `json:"deadline"` // RFC3339 формат
    }

    if err := c.ShouldBindJSON(&req); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": fmt.Sprintf("invalif request - %v", err)})
        return
    }

    // deadline, err := time.Parse(time.RFC3339, req.Deadline) // по хорошему проверка должна быть не здесь
    // if err != nil {
    //     c.JSON(http.StatusBadRequest, gin.H{"error": "invalid deadline format"})
    //     return
    // }

    grpcReq := &pb.CreateTaskRequest{
        Title:       req.Title,
        Description: req.Description,
        Deadline:    req.Deadline,
    }

    resp, err := g.Client.CreateTask(context.Background(), grpcReq)
    if err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
        return
    }

    c.JSON(http.StatusOK, resp.Task)
}

func (g *Gateway) GetTaskHandler(c *gin.Context) {
	var req struct {
		Id string `json:"id"`
	}

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid request"})
		return
	}

	grpcReq := &pb.GetTaskRequest{
		Id: req.Id,
	}

	resp, err := g.Client.GetTask(context.Background(), grpcReq)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
	}

	c.JSON(http.StatusOK, resp)

}

func (g *Gateway) ListTasksHandler(c *gin.Context) {
	grpcReq := &pb.ListTasksRequest{}

	resp, err := g.Client.ListTasks(context.Background(), grpcReq)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
	}

	c.JSON(http.StatusOK, resp)
}

func (g *Gateway) UpdateTaskHandler(c *gin.Context) {
	var req struct {
		Id string `json:"id"`
		Status string `json:"status"`
	}

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid request"})
		return
	}

	grpcReq := &pb.UpdateTaskRequest{
		Id: req.Id,
		Status: req.Status,
	}

	resp, err := g.Client.UpdateTask(context.Background(), grpcReq)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
	}

	c.JSON(http.StatusOK, resp)
}

func (g *Gateway) DeleteTaskHandler(c *gin.Context) {
	var req struct {
		Id string `json:"id"`
	}

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid request"})
		return
	}

	grpcReq := &pb.DeleteTaskRequest{
		Id: req.Id,
	}

	resp, err := g.Client.DeleteTask(context.Background(), grpcReq)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
	}

	c.JSON(http.StatusOK, resp)
}