package gateway

import (
	"context"
	"net/http"
	gatewayhelpers "tasker/internal/helpers/gateway_helpers"
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
// @Param        body   body	gateway.CreateTaskRequest true "Данные для создания задачи"
// @Success      200  {object}  gateway.TaskResponse
// @Failure      400  {object}  gateway.ErrorResponse
// @Failure      500  {object}  gateway.ErrorResponse
// @Router       /task [post]
func (g *Gateway) CreateTaskHandler(c *gin.Context) {
	var req CreateTaskRequest

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, ErrorResponse{Error: "invalid JSON: " + err.Error()})
		return
	}

	deadline, err := time.Parse(time.RFC3339, req.Deadline)
	if err != nil {
		c.JSON(http.StatusBadRequest, ErrorResponse{Error: "invalid date format: " + err.Error()})
		return
	}

	grpcReq := &pb.CreateTaskRequest{
        Title:       req.Title,
        Description: req.Description,
        Deadline:    timestamppb.New(deadline),
    }

	if err := gatewayhelpers.Validate(c, grpcReq); err != nil {
		c.JSON(http.StatusBadRequest, ErrorResponse{Error: "invalid request: " + err.Error()})
		return
	}

    resp, err := g.Client.CreateTask(context.Background(), grpcReq)
    if err != nil {
		c.JSON(http.StatusInternalServerError, ErrorResponse{Error: err.Error()})
        return
    }

    c.JSON(http.StatusOK, resp.Task)
}

// @Summary      Получение задачи
// @Description  Запрос задачи по id
// @Tags         task
// @Accept       json
// @Produce      json
// @Param        Id   query	gateway.GetTaskRequest true "Данные для получения задачи"
// @Success      200  {object}  gateway.TaskResponse
// @Failure      400  {object}  gateway.ErrorResponse
// @Failure      500  {object}  gateway.ErrorResponse
// @Router       /task [get]
func (g *Gateway) GetTaskHandler(c *gin.Context) {
	var reqQuery GetTaskRequest

	if err := c.ShouldBindQuery(&reqQuery); err != nil {
		c.JSON(http.StatusBadRequest, ErrorResponse{Error: "invalid JSON: " + err.Error()})
		return
	}

	grpcReq := &pb.GetTaskRequest{
		Id: reqQuery.Id,
	}

	if err := gatewayhelpers.Validate(c, grpcReq); err != nil {
		c.JSON(http.StatusBadRequest, ErrorResponse{Error: "invalid request: " + err.Error()})
		return
	}

	resp, err := g.Client.GetTask(context.Background(), grpcReq)
	if err != nil {
		c.JSON(http.StatusInternalServerError, ErrorResponse{Error: err.Error()})
		return
	}

	c.JSON(http.StatusOK, resp)

}

// @Summary      Получение списка задач
// @Description  Запрос получения всех задач
// @Tags         task
// @Accept       json
// @Produce      json
// @Success      200  {array}  gateway.TaskResponse
// @Failure      500  {object}  gateway.ErrorResponse
// @Router       /tasks [get]
func (g *Gateway) ListTasksHandler(c *gin.Context) {
	grpcReq := &pb.ListTasksRequest{}

	resp, err := g.Client.ListTasks(context.Background(), grpcReq)
	if err != nil {
		c.JSON(http.StatusInternalServerError, ErrorResponse{Error: err.Error()})
		return
	}

	c.JSON(http.StatusOK, resp)
}

// @Summary      Обновление задачи
// @Description  Обновляет статус задачи по ее id
// @Tags         task
// @Accept       json
// @Produce      json
// @Param        id		query	string true "Id задачи"
// @Param        body   body	gateway.UpdateTaskRequestBody true "Статус задачи"
// @Success      200  {object}  gateway.TaskResponse
// @Failure      400  {object}  gateway.ErrorResponse
// @Failure      500  {object}  gateway.ErrorResponse
// @Router       /task [put]
func (g *Gateway) UpdateTaskHandler(c *gin.Context) {
	var reqQuery UpdateTaskRequestQuery
	
	var reqBody UpdateTaskRequestBody

	if err := c.ShouldBindQuery(&reqQuery); err != nil {
		c.JSON(http.StatusBadRequest, ErrorResponse{Error: "invalid JSON: " + err.Error()})
		return
	}
	
	if err := c.ShouldBindJSON(&reqBody); err != nil {
		c.JSON(http.StatusBadRequest, ErrorResponse{Error: "invalid JSON: " + err.Error()})
		return
	}

	grpcReq := &pb.UpdateTaskRequest{
		Id: reqQuery.Id,
		Status: &reqBody.Status,
	}

	if err := gatewayhelpers.Validate(c, grpcReq); err != nil {
		c.JSON(http.StatusBadRequest, ErrorResponse{Error: "invalid request: " + err.Error()})
		return
	}

	resp, err := g.Client.UpdateTask(context.Background(), grpcReq)
	if err != nil {
		c.JSON(http.StatusInternalServerError, ErrorResponse{Error: err.Error()})
		return
	}

	c.JSON(http.StatusOK, resp)
}

// @Summary      Удаление задачи
// @Description  Удаляет задачу по ее id
// @Tags         task
// @Accept       json
// @Produce      json
// @Param        Id		query	gateway.DeleteTaskRequest true "Id задачи"
// @Success      200  {object}  gateway.TaskResponse
// @Failure      400  {object}  gateway.ErrorResponse
// @Failure      500  {object}  gateway.ErrorResponse
// @Router       /task [delete]
func (g *Gateway) DeleteTaskHandler(c *gin.Context) {
	var reqQuery DeleteTaskRequest

	if err := c.ShouldBindQuery(&reqQuery); err != nil {
		c.JSON(http.StatusBadRequest, ErrorResponse{Error: "invalid JSON: " + err.Error()})
		return
	}

	grpcReq := &pb.DeleteTaskRequest{
		Id: reqQuery.Id,
	}

	if err := gatewayhelpers.Validate(c, grpcReq); err != nil {
		c.JSON(http.StatusBadRequest, ErrorResponse{Error: "invalid request: " + err.Error()})
		return
	}

	resp, err := g.Client.DeleteTask(context.Background(), grpcReq)
	if err != nil {
		c.JSON(http.StatusInternalServerError, ErrorResponse{Error: err.Error()})
		return
	}

	c.JSON(http.StatusOK, resp)
}