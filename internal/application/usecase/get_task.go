package usecase

import (
	"context"
	"errors"
	"tasker/internal/domain"

	"github.com/google/uuid"
)

var NilIDError = errors.New("nil ID")


type GetTaskByIDUsecase struct {
	repo TaskRepository
}

func NewGetTaskByIDUsecase(repo TaskRepository) *GetTaskByIDUsecase {
	return &GetTaskByIDUsecase{repo: repo}
}

type GetTaskByIDDTO struct {
	TaskID uuid.UUID
}

func (u *GetTaskByIDUsecase) Execute(ctx context.Context, cmd *GetTaskByIDDTO) (*domain.Task, error) {
	if cmd == nil {
		return nil, NilDTOError
	}

	if cmd.TaskID == uuid.Nil {
		return nil, NilIDError
	}

	return u.repo.GetByID(ctx, cmd.TaskID)
}