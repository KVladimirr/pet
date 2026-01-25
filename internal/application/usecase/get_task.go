package usecase

import (
	"context"
	"tasker/internal/infrastracture/repository"
	"tasker/internal/domain"

	"github.com/google/uuid"
)


type GetTaskByIDUsecase struct {
	repo repository.TaskRepository
}

type GetTaskByIDDTO struct {
	TaskID uuid.UUID
}

func (u *GetTaskByIDUsecase) Execute(ctx context.Context, cmd *GetTaskByIDDTO) (*domain.Task, error) {
	return u.repo.GetByID(ctx, cmd.TaskID)
}