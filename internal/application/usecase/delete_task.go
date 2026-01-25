package usecase

import (
	"context"


	"github.com/google/uuid"
)


type DeleteTaskUsecase struct {
	repo TaskRepository
}

type DeleteTaskDTO struct {
	TaskID uuid.UUID
}

func (c *DeleteTaskUsecase) Execute(ctx context.Context, cmd *DeleteTaskDTO) error {
	if err := c.repo.Delete(ctx, cmd.TaskID); err != nil {
		return err
	}

	return nil
}