package task

import (
	"context"
	"fmt"
	"log"

	pb "tasker/internal/task/pb"

	"github.com/google/uuid"
	"google.golang.org/protobuf/types/known/timestamppb"
)

type Server struct {
	pb.UnimplementedTaskServiceServer
	Store *TaskStore
}

func (s *Server) CreateTask(ctx context.Context, req *pb.CreateTaskRequest) (*pb.TaskResponse, error) {
	log.Printf("CreateTask called: %+v", req)

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