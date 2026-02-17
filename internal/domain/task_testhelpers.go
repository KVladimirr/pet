package domain

import (
	"time"

	"github.com/google/uuid"
)


type TaskBuilder struct {
	Task *Task
}

func NewTaskBuilder() *TaskBuilder {
	task, _ := NewTask(
		"Base title",
		"Base Description",
		time.Now().Add(time.Hour),
	)

	return &TaskBuilder{
		Task: task,
	}
}

func (t *TaskBuilder) Title(title string) *TaskBuilder {
	t.Task.Title = title
	return t
}

func (t *TaskBuilder) Description(description string) *TaskBuilder {
	t.Task.Description = description
	return t
}

func (t *TaskBuilder) ID(id uuid.UUID) *TaskBuilder {
	t.Task.ID = id
	return t
}

func (t *TaskBuilder) Status(status TaskStatus) *TaskBuilder {
	t.Task.Status = status
	return t
}

func (t *TaskBuilder) Deadline(deadline time.Time) *TaskBuilder {
	t.Task.Deadline = deadline
	return t
}

func (t *TaskBuilder) CreatedAt(created_at time.Time) *TaskBuilder {
	t.Task.CreatedAt = created_at
	return t
}

func (t *TaskBuilder) UpdatedAt(updated_at time.Time) *TaskBuilder {
	t.Task.UpdatedAt = updated_at
	return t
}

func DeepCopyTask(src *Task) *Task {
	return &Task{
		ID: src.ID,
		Title: src.Title,
		Description: src.Description,
		Status: src.Status,
		Deadline: src.Deadline,
		CreatedAt: src.CreatedAt,
		UpdatedAt: src.UpdatedAt,
	}
}

