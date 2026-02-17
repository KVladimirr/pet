package usecase

import (
	"context"
	"time"

	"github.com/google/uuid"
)


type UpdateTaskDeadlineUsecase struct {
	repo TaskRepository
}

func NewUpdateTaskDeadlineUsecase(repo TaskRepository) *UpdateTaskDeadlineUsecase {
	return &UpdateTaskDeadlineUsecase{repo: repo}
}

type UpdateTaskDeadlineDTO struct {
	TaskID uuid.UUID
	NewDeadline time.Time
}

func (u *UpdateTaskDeadlineUsecase) Execute(ctx context.Context, cmd *UpdateTaskDeadlineDTO) error {
	if cmd == nil {
		return NilDTOError
	}

	task, err := u.repo.GetByID(ctx, cmd.TaskID)
	if err != nil {
		return err
	}

	if err = task.UpdateDeadline(cmd.NewDeadline); err != nil {
		return err
	}

	if err = u.repo.Save(ctx, task); err != nil {
		return err
	}

	return nil
}