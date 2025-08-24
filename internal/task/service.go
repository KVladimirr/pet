package task

import (
	"context"
	// "fmt"
	"log"

	pb "tasker/internal/task/pb"
	"tasker/internal/storage"

	"github.com/google/uuid"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/types/known/timestamppb"
)

type Server struct {
	pb.UnimplementedTaskServiceServer
	Store *storage.TaskStore
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
	s.Store.Create(ctx, task)

	return &pb.TaskResponse{Task: task}, nil
}

func (s *Server) GetTask(ctx context.Context, req *pb.GetTaskRequest) (*pb.TaskResponse, error) {
	log.Printf("GetTask called: %+v", req)

	task, err := s.Store.Get(ctx, req.Id)
	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}

	return &pb.TaskResponse{Task: task}, nil

}

func (s *Server) ListTasks(ctx context.Context, req *pb.ListTasksRequest) (*pb.ListTasksResponse, error) {
	log.Printf("ListTasks called: %+v", req)

	tasks, err := s.Store.List(ctx)
	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}

	return &pb.ListTasksResponse{Task: tasks}, nil
}

func (s *Server) UpdateTask(ctx context.Context, req *pb.UpdateTaskRequest) (*pb.TaskResponse, error) {
	log.Printf("UpdateTask called: %+v", req)

	_, err := s.Store.Get(ctx, req.Id)
	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}

	updatedTask, err := s.Store.UpdateStatus(ctx, req.Id, req.Status)
	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}

	return &pb.TaskResponse{Task: updatedTask}, nil
}

func (s *Server) DeleteTask(ctx context.Context, req *pb.DeleteTaskRequest) (*pb.DeleteTaskResponse, error) {
	log.Printf("DeleteTask called: %+v", req)

	_, err := s.Store.Get(ctx, req.Id)
	if err != nil {
		return &pb.DeleteTaskResponse{Success: false}, status.Error(codes.Internal, err.Error())
	}

	if err := s.Store.Delete(ctx, req.Id); err != nil{
		return &pb.DeleteTaskResponse{Success: false}, status.Error(codes.Internal, err.Error())
	}

	return &pb.DeleteTaskResponse{Success: true}, nil

}