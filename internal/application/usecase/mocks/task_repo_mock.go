package mocks

import (
	"context"
	"tasker/internal/domain"

	"github.com/google/uuid"
)

type MockRepository struct {
	saveWasCalled	bool
	saveArgs		*domain.Task
	saveError		error

	getByIDWasCalled	bool
	getByIDArgs			uuid.UUID
	getByIDError		error
	
	getAllWasCalled	bool
	getAllArgsLimit	uint	
	getAllArgsOffset uint
	getAllError		error

	deleteWasCalled	bool
	deleteArgs		uuid.UUID
	deleteError		error

}

func (m *MockRepository) Save(ctx context.Context, task *domain.Task) error {
	m.saveWasCalled = true
	m.saveArgs = task
	return m.saveError
}

func (m *MockRepository) GetByID(ctx context.Context, id uuid.UUID) (*domain.Task, error) {
	m.getByIDWasCalled = true
	m.getByIDArgs = id
	return &domain.Task{}, m.getByIDError
}

func (m *MockRepository) GetAll(ctx context.Context, limit uint, offset uint) ([]*domain.Task, error) {
	m.getAllWasCalled = true
	m.getAllArgsLimit = limit
	m.getAllArgsOffset = offset
	return []*domain.Task{
		{},
	}, m.getAllError
}

func (m *MockRepository) Delete(ctx context.Context, id uuid.UUID) error {
	m.deleteWasCalled = true
	m.deleteArgs = id
	return m.deleteError
}
