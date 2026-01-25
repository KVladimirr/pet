package grpc

import (
	"context"
	"tasker/internal/application/usecase"
	pb "tasker/internal/task/pb"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/types/known/timestamppb"
)

type TaskService struct {
	pb.UnimplementedTaskServiceServer

	CreateTaskUC		*usecase.CreateTaskUsecase
	GetAllTasksUC		*usecase.GetAllTasksUsecase
	GetTaskByIDUC 		*usecase.GetTaskByIDUsecase
	UpdateDeadlineUC 	*usecase.UpdateTaskDeadlineUsecase
	UpdateDescriptionUC *usecase.UpdateTaskDescriptionUsecase
	UpdateStatusUC 		*usecase.UpdateTaskStatusUsecase
	UpdateTitleUC 		*usecase.UpdateTaskTitleUsecase
}

func NewTaskService(
	createTaskUC *usecase.CreateTaskUsecase,
	getAllTasksUC *usecase.GetAllTasksUsecase,
	getTaskByIDUC *usecase.GetTaskByIDUsecase,
	updateDeadlineUC *usecase.UpdateTaskDeadlineUsecase,
	updateDescriptionUC *usecase.UpdateTaskDescriptionUsecase,
	updateStatusUC *usecase.UpdateTaskStatusUsecase,
	updateTitleUC *usecase.UpdateTaskTitleUsecase,
) (*TaskService, error) {
	return &TaskService{
		CreateTaskUC: createTaskUC,
		GetTaskByIDUC: getTaskByIDUC,
		GetAllTasksUC: getAllTasksUC,
		UpdateDeadlineUC: updateDeadlineUC,
		UpdateDescriptionUC: updateDescriptionUC,
		UpdateStatusUC: updateStatusUC,
		UpdateTitleUC: updateTitleUC,
	}, nil
}

func (t *TaskService) CreateTask(ctx context.Context, req *pb.CreateTaskRequest) (*pb.TaskResponse, error) {
	if err := CreateTaskValidation(req); err != nil {
		return nil, status.Errorf(codes.InvalidArgument, err.Error())
	}

	cmd := usecase.CreateTaskDTO{
        Title:       req.Title,
        Description: req.Description,
        Deadline:    req.Deadline.AsTime(),
    }

	task, err := t.CreateTaskUC.Execute(ctx, &cmd)
	if err != nil {
		return nil, status.Errorf(codes.Internal, "failed to create task: %v", err)
	}

	return &pb.TaskResponse{
		Task: &pb.Task{
			Id: task.ID.String(),
			Title: task.Title,
			Description: task.Description,
			Status: string(task.Status),
			Deadline: timestamppb.New(task.Deadline),
			CreatedAt: timestamppb.New(task.CreatedAt),
			UpdatedAt: timestamppb.New(task.UpdatedAt),
		},
	}, nil
}

func (t *TaskService) UpdateTask(ctx context.Context, req *pb.UpdateTaskRequest) (*pb.TaskResponse, error) {
	if req.Id != "" {
		
	}
}