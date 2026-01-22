package domain

import (
	"errors"
	"time"

	"github.com/google/uuid"
)

type Task struct {
	id uuid.UUID
	title string
	description string
	status TaskStatus
	deadline time.Time
	createdAt time.Time
	updatedAt time.Time
}

type TaskStatus string 

const (
	TaskStatusTodo			TaskStatus = "TODO"
	TaskStatusInProgress	TaskStatus = "INPROGRESS"
	TaskStatusDone			TaskStatus = "DONE"
	TaskStatusCanceled		TaskStatus = "CANCELED"
	TaskStatusHold 			TaskStatus = "HOLD"
)

var (
    ErrCannotEditTask = errors.New("cannot change finished task")
	ErrEmptyTitle = errors.New("title cannot be empty")
	ErrOldDeadline = errors.New("deadline cannot be in the past")
	ErrInvalidStatus = errors.New("invalid status")
)

func NewTask(title string, description string, deadline time.Time) (*Task, error) {
    if title == "" {
        return nil, errors.New("title cannot be empty")
    }

	now := time.Now()

	if deadline.Before(now) {
		return nil, ErrOldDeadline
	}

	return &Task{
		id: uuid.New(),
		title: title,
		description: description,
		status: TaskStatusTodo,
		deadline: deadline,
		createdAt: now,
		updatedAt: now,
	}, nil
}

func (t *Task) GetID() uuid.UUID { return t.id}
func (t *Task) GetTitle() string { return t.title}
func (t *Task) GetDescription() string { return t.description}
func (t *Task) GetStatus() TaskStatus { return t.status}
func (t *Task) GetDeadline() time.Time { return t.deadline}
func (t *Task) GetCreatedAt() time.Time { return t.createdAt}
func (t *Task) GetUpdatedAt() time.Time { return t.updatedAt}

func (t *Task) UpdateStatus(newStatus TaskStatus) error {
	if !t.CanEdit() {
		return ErrCannotEditTask
	}

	switch newStatus{
	case TaskStatusTodo, TaskStatusInProgress, TaskStatusHold, TaskStatusCanceled, TaskStatusDone:
		t.status = newStatus
		t.updatedAt = time.Now()
		return nil
	default:
		return ErrInvalidStatus
	}
}

func (t *Task) UpdateTitle(newTitle string) error {
	if !t.CanEdit() {
		return ErrCannotEditTask
	}

	if newTitle == "" {
		return ErrEmptyTitle
	}

	t.title = newTitle
	t.updatedAt = time.Now()
	return nil
}

func (t *Task) UpdateDescription(newDescription string) error {
	if !t.CanEdit() {
		return ErrCannotEditTask
	}

	t.description = newDescription
	t.updatedAt = time.Now()
	return nil
}

func (t *Task) UpdateDeadline(newDeadline time.Time) error {
	if !t.CanEdit() {
		return ErrCannotEditTask
	}

	t.deadline = newDeadline
	t.updatedAt = time.Now()
	return nil
}

func (t *Task) IsOverdue() bool {
	return time.Now().After(t.deadline)
}

func (t *Task) IsActive() bool {
	return (t.status == TaskStatusTodo || t.status == TaskStatusInProgress) && !t.IsOverdue()
}

func (t *Task) IsFinished() bool {
	return (t.status == TaskStatusCanceled || t.status == TaskStatusDone)
}

func (t *Task) CanEdit() bool {
	return !t.IsOverdue() && !t.IsFinished()
}
