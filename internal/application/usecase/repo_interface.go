package usecase

import (
	"context"
	"tasker/internal/domain"

	"github.com/google/uuid"
)


type TaskRepository interface {
	Save(ctx context.Context, task *domain.Task) error
    GetByID(ctx context.Context, id uuid.UUID) (*domain.Task, error)
    GetAll(ctx context.Context, limit uint, offset uint) ([]*domain.Task, error)
    Delete(ctx context.Context, id uuid.UUID) error
}

