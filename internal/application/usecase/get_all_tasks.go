package usecase

import (
	"context"
	"tasker/internal/domain"
)


const MaxLimit = 100

type GetAllTasksUsecase struct {
	repo TaskRepository
}

func NewGetAllTasksUsecase(repo TaskRepository) *GetAllTasksUsecase {
	return &GetAllTasksUsecase{repo: repo}
}

type GetAllTasksDTO struct {
	Limit uint
	Offset uint
}

func (u *GetAllTasksUsecase) Execute(ctx context.Context, cmd *GetAllTasksDTO) ([]*domain.Task, error) {
	if cmd == nil {
		return nil, NilDTOError
	}

	if cmd.Limit == 0 {
		return []*domain.Task{}, nil
	}

	if cmd.Limit > MaxLimit {
		cmd.Limit = MaxLimit
	}

	return u.repo.GetAll(ctx, cmd.Limit, cmd.Offset)
	// Как вариант на стороне репозитория тоже сделать DTO, а не просто аргументы функции
}