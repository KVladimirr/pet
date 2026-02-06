package domain

import (
	"errors"
	"time"

	"github.com/google/uuid"
)

type Task struct {
	ID uuid.UUID
	Title string
	Description string
	Status TaskStatus
	Deadline time.Time
	CreatedAt time.Time
	UpdatedAt time.Time
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
	ErrEmptyDeadline = errors.New("deadline cannot be empty")
	ErrOldDeadline = errors.New("deadline cannot be in the past")
	ErrInvalidStatus = errors.New("invalid status")
)

func NewTask(title string, description string, deadline time.Time) (*Task, error) {
    if title == "" {
        return nil, errors.New("title cannot be empty")
    }

	now := time.Now()

	if deadline.IsZero() {
        return nil, ErrEmptyDeadline
    }

	if deadline.Before(now) {
		return nil, ErrOldDeadline
	}

	return &Task{
		ID: uuid.New(),
		Title: title,
		Description: description,
		Status: TaskStatusTodo,
		Deadline: deadline,
		CreatedAt: now,
		UpdatedAt: now,
	}, nil
}

func (t *Task) UpdateStatus(newStatus TaskStatus) error {
	if !t.CanEdit() {
		return ErrCannotEditTask
	}

	switch newStatus{
	case TaskStatusTodo, TaskStatusInProgress, TaskStatusHold, TaskStatusCanceled, TaskStatusDone:
		t.Status = newStatus
		t.UpdatedAt = time.Now()
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

	t.Title = newTitle
	t.UpdatedAt = time.Now()
	return nil
}

func (t *Task) UpdateDescription(newDescription string) error {
	if !t.CanEdit() {
		return ErrCannotEditTask
	}

	t.Description = newDescription
	t.UpdatedAt = time.Now()
	return nil
}

func (t *Task) UpdateDeadline(newDeadline time.Time) error {
	now := time.Now()

	if !t.CanEdit() {
		return ErrCannotEditTask
	}

	if newDeadline.IsZero() {
        return ErrEmptyDeadline
    }

	if newDeadline.Before(now) {
		return ErrOldDeadline
	}

	t.Deadline = newDeadline
	t.UpdatedAt = time.Now()
	return nil
}

func (t *Task) IsOverdue() bool {
	return time.Now().After(t.Deadline)
}

func (t *Task) IsActive() bool {
	return (t.Status == TaskStatusTodo || t.Status == TaskStatusInProgress) && !t.IsOverdue()
}

func (t *Task) IsFinished() bool {
	return (t.Status == TaskStatusCanceled || t.Status == TaskStatusDone)
}

func (t *Task) CanEdit() bool {
	return !t.IsOverdue() && !t.IsFinished()
}
