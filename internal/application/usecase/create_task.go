package usecase

import (
	"context"
	"errors"
	"tasker/internal/domain"
	"time"
)

var NilDTOError = errors.New("Nil dto")

type CreateTaskUsecase struct {
	repo TaskRepository
}

func NewCreateTaskUsecase(repo TaskRepository) *CreateTaskUsecase {
	return &CreateTaskUsecase{repo: repo}
}

type CreateTaskDTO struct {
	Title string
	Description string
	Deadline time.Time
}

func (c *CreateTaskUsecase) Execute(ctx context.Context, cmd *CreateTaskDTO) (*domain.Task, error) {
	if cmd == nil {
		return nil, NilDTOError
	}

	task, err := domain.NewTask(cmd.Title, cmd.Description, cmd.Deadline)
	if err != nil {
		return nil, err
	}

	if err := c.repo.Save(ctx, task); err != nil {
		return nil, err
	}

	return task, nil
}

/*
Что мы хоти проверить в этом методе?

Валидный кейс:
- Что создалась задача с переданными полями
- Что в мок репозитории вызвался нужный метод
- Что в мок репозитории вызвался метод с нужными полями (маппинг в dto репы (пока это доменный тип))
- Что нет ошибки
*/