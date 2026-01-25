package usecase

import (
	"context"
	"tasker/internal/domain"
)


type GetAllTasksUsecase struct {
	repo TaskRepository
}

type GetAllTasksDTO struct {
	Limit uint
	Offset uint
}

func (u *GetAllTasksUsecase) Execute(ctx context.Context, cmd *GetAllTasksDTO) ([]*domain.Task, error) {
	return u.repo.GetAll(ctx, cmd.Limit, cmd.Offset)
	// Как вариант на стороне репозитория тоже сделать DTO, а не просто аргументы функции
}