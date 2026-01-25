package usecase

import (
	"context"
	"tasker/internal/infrastracture/repository"

	"github.com/google/uuid"
)


type UpdateTaskTitleUsecase struct {
	repo repository.TaskRepository
}

type UpdateTaskTitleDTO struct {
	TaskID uuid.UUID
	NewTitle string
}

func (u *UpdateTaskTitleUsecase) Execute(ctx context.Context, cmd *UpdateTaskTitleDTO) error {
	task, err := u.repo.GetByID(ctx, cmd.TaskID)
	if err != nil {
		return err
	}

	if err = task.UpdateTitle(cmd.NewTitle); err != nil {
		return err
	}

	if err = u.repo.Save(ctx, task); err != nil {
		return err
	}

	return nil
}