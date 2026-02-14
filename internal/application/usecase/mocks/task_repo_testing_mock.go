package mocks

import (
	"context"
	"tasker/internal/domain"

	"github.com/google/uuid"
)

type MockRepositoryForCreate struct {
	SaveWasCalled	bool
	SaveArgs		*domain.Task
	SaveError		error
}

func (m *MockRepositoryForCreate) Save(ctx context.Context, task *domain.Task) error {
	m.SaveWasCalled = true
	m.SaveArgs = task
	return m.SaveError
}

func (m *MockRepositoryForCreate) GetByID(ctx context.Context, id uuid.UUID) (*domain.Task, error) {return nil, nil}

func (m *MockRepositoryForCreate) GetAll(ctx context.Context, limit uint, offset uint) ([]*domain.Task, error) {return nil, nil}

func (m *MockRepositoryForCreate) Delete(ctx context.Context, id uuid.UUID) error {return nil}
