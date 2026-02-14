package mocks

import (
	"context"
	"tasker/internal/domain"

	"github.com/google/uuid"
	"github.com/stretchr/testify/mock"
)

type MockTestifyRepo struct {
	mock.Mock
}

func (m *MockTestifyRepo) Save(ctx context.Context, task *domain.Task) error {
	args := m.Called(ctx, task) // записывает в объект мока: название метода, кол-во вызовов, аргументы. Возврвщает новый Arguments, где содержатся только выходные аргументы
	return args.Error(0) // возвращает первый из выходных параметров, которые задаются в .Return()
}

func (m *MockTestifyRepo) GetByID(ctx context.Context, id uuid.UUID) (*domain.Task, error) {
	args := m.Called(ctx, id)
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}
	return args.Get(0).(*domain.Task), args.Error(1)
}

func (m *MockTestifyRepo) GetAll(ctx context.Context, limit uint, offset uint) ([]*domain.Task, error) {
	args := m.Called(ctx, limit, offset)
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}
	return args.Get(0).([]*domain.Task), args.Error(1)
}

func (m *MockTestifyRepo) Delete(ctx context.Context, id uuid.UUID) error {
	args := m.Called(ctx, id)
	return args.Error(0)
}
