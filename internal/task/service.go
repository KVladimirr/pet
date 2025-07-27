package task

import (
	"context"
	"fmt"
	"log"

	pb "tasker/internal/task/pb"

	"github.com/google/uuid"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/types/known/timestamppb"
)

type Server struct {
	pb.UnimplementedTaskServiceServer
	Store *TaskStore
}

func (s *Server) CreateTask(ctx context.Context, req *pb.CreateTaskRequest) (*pb.TaskResponse, error) {
	log.Printf("CreateTask called: %+v", req)

	if err := CreateTaskValidation(req); err != nil {
		return nil, status.Errorf(codes.InvalidArgument, err.Error())
	}

	task := &pb.Task{
		Id: uuid.New().String(),
		Title: req.Title,
		Description: req.Description,
		Status: "TODO",
		Deadline: req.Deadline,
		CreatedAt: timestamppb.Now(),
		UpdatedAt: timestamppb.Now(),
	}
	s.Store.Create(task)

	return &pb.TaskResponse{Task: task}, nil
}

func (s *Server) GetTask(ctx context.Context, req *pb.GetTaskRequest) (*pb.TaskResponse, error) {
	log.Printf("GetTask called: %+v", req)

	task, ok := s.Store.Get(req.Id)
	
	if !ok {
		return nil, fmt.Errorf("task with id %q not found", req.Id)
	}

	return &pb.TaskResponse{Task: task}, nil

}

func (s *Server) ListTasks(ctx context.Context, req *pb.ListTasksRequest) (*pb.ListTasksResponse, error) {
	log.Printf("ListTasks called: %+v", req)

	tasks := s.Store.List()

	return &pb.ListTasksResponse{Task: tasks}, nil
}

func (s *Server) UpdateTask(ctx context.Context, req *pb.UpdateTaskRequest) (*pb.TaskResponse, error) {
	log.Printf("UpdateTask called: %+v", req)

	task, ok := s.Store.Get(req.Id)
	if !ok {
		return nil, fmt.Errorf("task with id %q not found", req.Id)
	}

	s.Store.UpdateStatus(req.Id, req.Status)

	return &pb.TaskResponse{Task: task}, nil
}

func (s *Server) DeleteTask(ctx context.Context, req *pb.DeleteTaskRequest) (*pb.DeleteTaskResponse, error) {
	log.Printf("DeleteTask called: %+v", req)

	_, ok := s.Store.Get(req.Id)
	if !ok {
		return nil, fmt.Errorf("task with id %q not found", req.Id)
	}

	res := s.Store.Delete(req.Id)

	return &pb.DeleteTaskResponse{Success: res}, nil

}