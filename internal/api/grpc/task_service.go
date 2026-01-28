package grpc

import (
	"context"
	"tasker/internal/application/usecase"
	"tasker/internal/domain"
	pb "tasker/internal/pb"

	"github.com/google/uuid"
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
	DeleteTaskUC 		*usecase.DeleteTaskUsecase
}

func NewTaskService(
	createTaskUC *usecase.CreateTaskUsecase,
	getAllTasksUC *usecase.GetAllTasksUsecase,
	getTaskByIDUC *usecase.GetTaskByIDUsecase,
	updateDeadlineUC *usecase.UpdateTaskDeadlineUsecase,
	updateDescriptionUC *usecase.UpdateTaskDescriptionUsecase,
	updateStatusUC *usecase.UpdateTaskStatusUsecase,
	updateTitleUC *usecase.UpdateTaskTitleUsecase,
	deleteTaskUC *usecase.DeleteTaskUsecase,
) *TaskService {
	return &TaskService{
		CreateTaskUC: createTaskUC,
		GetTaskByIDUC: getTaskByIDUC,
		GetAllTasksUC: getAllTasksUC,
		UpdateDeadlineUC: updateDeadlineUC,
		UpdateDescriptionUC: updateDescriptionUC,
		UpdateStatusUC: updateStatusUC,
		UpdateTitleUC: updateTitleUC,
		DeleteTaskUC: deleteTaskUC,
	}
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
	taskID, err := uuid.Parse(req.GetId())
	if err != nil {
		return nil, status.Errorf(codes.InvalidArgument, "failed to parse uuid: %v", err.Error())
	}

	task, err := t.GetTaskByIDUC.Execute(ctx, &usecase.GetTaskByIDDTO{TaskID: taskID})
	if err != nil {
		return nil, status.Errorf(codes.NotFound, "task not found: %v", err.Error())
	}

	if req.GetTitle() != "" && req.GetTitle() != task.Title {
		err = t.UpdateTitleUC.Execute(ctx, &usecase.UpdateTaskTitleDTO{TaskID: taskID, NewTitle: req.GetTitle()})
		if err != nil {
			return nil, status.Errorf(codes.Internal, "failed to update task title: %v", err.Error())
		}
	}

	if req.Description != nil {
		if req.GetDescription() != task.Description{
			err = t.UpdateDescriptionUC.Execute(ctx, &usecase.UpdateTaskDescriptionDTO{TaskID: taskID, NewDescription: req.GetDescription()})
			if err != nil {
				return nil, status.Errorf(codes.Internal, "failed to update task description: %v", err.Error())
			}
		}
	}

	if req.GetStatus() != "" && req.GetStatus() != string(task.Status) {
		err = t.UpdateStatusUC.Execute(ctx, &usecase.UpdateTaskStatusDTO{
			TaskID: taskID,
			NewStatus: domain.TaskStatus(req.GetStatus()),
		})
		if err != nil {
			return nil, status.Errorf(codes.Internal, "failed to update task status: %v", err.Error())
		}
	}

	if req.GetDeadline() != nil && !req.GetDeadline().AsTime().Equal(task.Deadline) {
		err = t.UpdateDeadlineUC.Execute(ctx, &usecase.UpdateTaskDeadlineDTO{TaskID: taskID, NewDeadline: req.GetDeadline().AsTime()})
		if err != nil {
			return nil, status.Errorf(codes.Internal, "failed to update task deadline: %v", err.Error())
		}
	}

	updatedTask, err := t.GetTaskByIDUC.Execute(ctx, &usecase.GetTaskByIDDTO{TaskID: taskID})
	if err != nil {
		return nil, status.Errorf(codes.NotFound, "updated task not found: %v", err.Error())
	}

	return &pb.TaskResponse{
		Task: &pb.Task{
			Id: updatedTask.ID.String(),
			Title: updatedTask.Title,
			Description: updatedTask.Description,
			Status: string(updatedTask.Status),
			Deadline: timestamppb.New(updatedTask.Deadline),
			CreatedAt: timestamppb.New(updatedTask.CreatedAt),
			UpdatedAt: timestamppb.New(updatedTask.UpdatedAt),
		},
	}, nil
}

func (t *TaskService) GetTask(ctx context.Context, req *pb.GetTaskRequest) (*pb.TaskResponse, error) {
	taskID, err := uuid.Parse(req.GetId())
	if err != nil {
		return nil, status.Errorf(codes.InvalidArgument, "failed to parse uuid: %v", err.Error())
	}

	cmd := &usecase.GetTaskByIDDTO{TaskID: taskID}
	
	task, err := t.GetTaskByIDUC.Execute(ctx, cmd)
	if err != nil {
		return nil, status.Errorf(codes.NotFound, "task not found: %v", err.Error())
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

func (t *TaskService) ListTasks(ctx context.Context, req *pb.ListTasksRequest) (*pb.ListTasksResponse, error) {
	select {
    case <-ctx.Done():
        return nil, status.FromContextError(ctx.Err()).Err()
    default:
    }

	cmd := &usecase.GetAllTasksDTO{Limit: uint(req.GetLimit()), Offset: uint(req.GetOffset())}

	tasks, err := t.GetAllTasksUC.Execute(ctx, cmd)
	if err != nil {
		return nil, status.Errorf(codes.Internal, "failed to list tasks: %v", err.Error())
	}

	handlerTasks := make([]*pb.Task, len(tasks))

	for i, task := range tasks {
		handlerTasks[i] = &pb.Task{
			Id: task.ID.String(),
			Title: task.Title,
			Description: task.Description,
			Status: string(task.Status),
			Deadline: timestamppb.New(task.Deadline),
			CreatedAt: timestamppb.New(task.CreatedAt),
			UpdatedAt: timestamppb.New(task.UpdatedAt),
		}
	}

	return &pb.ListTasksResponse{Task: handlerTasks}, nil
}

func (t *TaskService) DeleteTask(ctx context.Context, req *pb.DeleteTaskRequest) (*pb.DeleteTaskResponse, error) {
	taskID, err := uuid.Parse(req.GetId())
	if err != nil {
		return nil, status.Errorf(codes.InvalidArgument, "failed to parse uuid: %v", err.Error())
	}
	
	cmd := &usecase.DeleteTaskDTO{TaskID: taskID}

	if err = t.DeleteTaskUC.Execute(ctx, cmd); err != nil {
		return nil, status.Errorf(codes.Internal, "failed to delete task: %v", err.Error())
	}

	return &pb.DeleteTaskResponse{Success: true}, nil
}