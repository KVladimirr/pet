package usecase

import (
	"context"
	"tasker/internal/infrastracture/repository"

	"github.com/google/uuid"
)


type UpdateTaskDescriptionUsecase struct {
	repo repository.TaskRepository
}

type UpdateTaskDescriptionDTO struct {
	TaskID uuid.UUID
	NewDescription string
}

func (u *UpdateTaskDescriptionUsecase) Execute(ctx context.Context, cmd *UpdateTaskDescriptionDTO) error {
	task, err := u.repo.GetByID(ctx, cmd.TaskID)
	if err != nil {
		return err
	}

	if err = task.UpdateDescription(cmd.NewDescription); err != nil {
		return err
	}

	if err = u.repo.Save(ctx, task); err != nil {
		return err
	}

	return nil
}