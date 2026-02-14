package usecase

import (
	"fmt"
	"tasker/internal/domain"
	"time"

	"github.com/google/uuid"
)


func GenerateTasks(count int) []*domain.Task {
	tasks := make([]*domain.Task, count)
	for i := 0; i < count; i++ {
		now := time.Now()
		tasks[i] = &domain.Task{
            ID:          uuid.New(),
            Title:       fmt.Sprintf("Gen task %d", i),
            Description: fmt.Sprintf("Gen description %d", i),
            Status:      domain.TaskStatusTodo,
            CreatedAt:   now,
			UpdatedAt: 	 now,	
            Deadline:    now.Add(24 * time.Hour),
        }
	}
	return tasks
}