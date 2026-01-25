package usecase

import (
	"context"
	"tasker/internal/infrastracture/repository"
	"tasker/internal/domain"
	"time"
)


type CreateTaskUsecase struct {
	repo repository.TaskRepository
}

type CreateTaskDTO struct {
	Title string
	Description string
	Deadline time.Time
}

func (c *CreateTaskUsecase) Execute(ctx context.Context, cmd *CreateTaskDTO) (*domain.Task, error) {
	task, err := domain.NewTask(cmd.Title, cmd.Description, cmd.Deadline)
	if err != nil {
		return nil, err
	}

	if err := c.repo.Save(ctx, task); err != nil {
		return nil, err
	}

	return task, nil
}