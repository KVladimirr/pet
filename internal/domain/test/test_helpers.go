package test

import (
	"tasker/internal/domain"
	"time"

	"github.com/google/uuid"
)


type TaskBuilder struct {
	task *domain.Task
}

func NewTaskBuilder() *TaskBuilder {
	task, _ := domain.NewTask(
		"Base title",
		"Base Description",
		time.Now().Add(time.Hour),
	)

	return &TaskBuilder{
		task: task,
	}
}

func (t *TaskBuilder) Title(title string) *TaskBuilder {
	t.task.Title = title
	return t
}

func (t *TaskBuilder) Description(description string) *TaskBuilder {
	t.task.Description = description
	return t
}

func (t *TaskBuilder) ID(id uuid.UUID) *TaskBuilder {
	t.task.ID = id
	return t
}

func (t *TaskBuilder) Status(status domain.TaskStatus) *TaskBuilder {
	t.task.Status = status
	return t
}

func (t *TaskBuilder) Deadline(deadline time.Time) *TaskBuilder {
	t.task.Deadline = deadline
	return t
}

func (t *TaskBuilder) CreatedAt(created_at time.Time) *TaskBuilder {
	t.task.CreatedAt = created_at
	return t
}

func (t *TaskBuilder) UpdatedAt(updated_at time.Time) *TaskBuilder {
	t.task.UpdatedAt = updated_at
	return t
}

