package task

import (
	"fmt"
	"tasker/config"
	pb "tasker/internal/task/pb"
	"time"
)


func CreateTaskValidation(req *pb.CreateTaskRequest) error {
	if len(req.Title) > config.TITLE_LEN {
		return fmt.Errorf("title length must be <= %d", config.TITLE_LEN)
	}

	if len(req.Description) > config.DESCRIPTION_LEN {
		return fmt.Errorf("description length must be <= %d", config.DESCRIPTION_LEN)
	}

	deadline := req.Deadline.AsTime()
	now := time.Now()
	if deadline.Before(now) {
		return fmt.Errorf("deadline must be in future")
	}

	return nil
}