package gateway

import (
	"context"
	"net/http"
	pb "tasker/internal/task/pb"
	"time"

	"github.com/gin-gonic/gin"
	"google.golang.org/protobuf/types/known/timestamppb"
)

type Gateway struct {
	Client pb.TaskServiceClient
}

func New(client pb.TaskServiceClient) *Gateway {
	return &Gateway{Client: client}
}

// @Summary      Создание задачи
// @Description  Создает задачу
// @Tags         task
// @Accept       json
// @Produce      json
// @Param        request   body	gateway.CreateTaskRequest true "Данные для создания задачи"
// @Success      200  {object}  gateway.CreateTaskResponse
// @Failure      400  {object}  gateway.ErrorResponse
// @Failure      500  {object}  gateway.ErrorResponse
// @Router       /task [post]
func (g *Gateway) CreateTaskHandler(c *gin.Context) {
	var req CreateTaskRequest

    if err := c.ShouldBindJSON(&req); err != nil {
        c.JSON(http.StatusBadRequest, ErrorResponse{Error: "invalid request - " + err.Error()})
        return
    }

	deadline, err := time.Parse(time.RFC3339, req.Deadline)
	if err != nil {
		c.JSON(http.StatusBadRequest, ErrorResponse{Error: "invalid date format"})
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

func (g *Gateway) GetTaskHandler(c *gin.Context) {
	var reqQuery struct {
		Id string `form:"id"`
	}

	if err := c.ShouldBindQuery(&reqQuery); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid request"})
		return
	}

	grpcReq := &pb.GetTaskRequest{
		Id: reqQuery.Id,
	}

	resp, err := g.Client.GetTask(context.Background(), grpcReq)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, resp)

}

func (g *Gateway) ListTasksHandler(c *gin.Context) {
	grpcReq := &pb.ListTasksRequest{}

	resp, err := g.Client.ListTasks(context.Background(), grpcReq)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, resp)
}

func (g *Gateway) UpdateTaskHandler(c *gin.Context) {
	var reqQuery struct {
		Id string `form:"id"`
	}
	
	var reqBody struct {
		Status string `json:"status"`
	}

	if err := c.ShouldBindQuery(&reqQuery); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid request"})
		return
	}
	
	if err := c.ShouldBindJSON(&reqBody); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid request"})
		return
	}

	grpcReq := &pb.UpdateTaskRequest{
		Id: reqQuery.Id,
		Status: reqBody.Status,
	}

	resp, err := g.Client.UpdateTask(context.Background(), grpcReq)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, resp)
}

func (g *Gateway) DeleteTaskHandler(c *gin.Context) {
	var reqQuery struct {
		Id string `form:"id"`
	}

	if err := c.ShouldBindQuery(&reqQuery); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid request"})
		return
	}

	grpcReq := &pb.DeleteTaskRequest{
		Id: reqQuery.Id,
	}

	resp, err := g.Client.DeleteTask(context.Background(), grpcReq)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, resp)
}