package usecase

import (
	"context"
	"tasker/internal/domain"

	"github.com/google/uuid"
)


type UpdateTaskStatusUsecase struct {
	repo TaskRepository
}

func NewUpdateTaskStatusUsecase(repo TaskRepository) *UpdateTaskStatusUsecase {
	return &UpdateTaskStatusUsecase{repo: repo}
}

type UpdateTaskStatusDTO struct {
	TaskID uuid.UUID
	NewStatus domain.TaskStatus
}

func (u *UpdateTaskStatusUsecase) Execute(ctx context.Context, cmd *UpdateTaskStatusDTO) error {
	task, err := u.repo.GetByID(ctx, cmd.TaskID)
	if err != nil {
		return err
	}

	if err = task.UpdateStatus(cmd.NewStatus); err != nil {
		return err
	}

	if err = u.repo.Save(ctx, task); err != nil {
		return err
	}

	return nil
}