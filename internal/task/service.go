package task

import (
	"context"
	"log"

	pb "tasker/internal/task/pb"
)

type Server struct {
	pb.UnimplementedTaskServiceServer
}

func (s *Server) CreateTask(ctx context.Context, req *pb.CreateTaskRequest) (*pb.TaskResponse, error) {
	log.Printf("CreateTask called: %+v", req)

	return &pb.TaskResponse{
		Task: &pb.Task{
			Id: "1",
			Title: req.GetTitle(),
			Description: req.GetDescription(),
		},
	}, nil
}