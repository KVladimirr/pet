package storage

import (
	"context"
	"errors"
	"fmt"
	pb "tasker/internal/task/pb"
	"time"

	"github.com/jackc/pgx/v5"
	"google.golang.org/protobuf/types/known/timestamppb"
)


type TaskStore struct {
	db *pgConnector
}

func NewTaskStore(dbConn *pgConnector) *TaskStore {
	return &TaskStore{db: dbConn}
}


// Методы для работы с БД
func (t *TaskStore) Create(ctx context.Context, task *pb.Task) error {
	query := `
		INSERT INTO tasks (
			id,
			title,
			description,
			status,
			deadline,
			created_at,
			updated_at
		) VALUES (
			$1, $2, $3, $4, $5, $6, $7
		)
	`

	_, err := t.db.Conn.Exec(
		ctx,
		query,
		task.GetId(),
		task.GetTitle(),
		task.GetDescription(),
		task.GetStatus(),
		task.GetDeadline().AsTime(),
		task.GetCreatedAt().AsTime(),
		task.GetUpdatedAt().AsTime(),
	)
	if err != nil {
		return fmt.Errorf("failed to create task: %w", err)
	}

	return nil
}

func (t *TaskStore) Get(ctx context.Context, id string) (*pb.Task, error) {
	query := `
		SELECT * FROM tasks
		WHERE id = $1
	`

	var task pb.Task
	var deadline, created_at, updated_at time.Time
	
	err := t.db.Conn.QueryRow(ctx, query, id).Scan(
		&task.Id,
		&task.Title,
		&task.Description,
		&task.Status,
		&deadline,
		&created_at,
		&updated_at,
	)
	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return nil, fmt.Errorf("task not found")
		}

		return nil, fmt.Errorf("failed to get task: %w", err)
	}

	task.Deadline = timestamppb.New(deadline)
	task.CreatedAt = timestamppb.New(created_at)
	task.UpdatedAt = timestamppb.New(updated_at)

	return &task, nil
}

func (t *TaskStore) List(ctx context.Context) ([]*pb.Task, error) {
	query := "SELECT * FROM tasks"

	var result []*pb.Task
	
	rows, err := t.db.Conn.Query(ctx, query)
	if err != nil {
		return nil, fmt.Errorf("failed to get all tasks: %w", err)
	}
	defer rows.Close()

	for rows.Next() {
		var task pb.Task
		var deadline, createdAt, updatedAt time.Time

		err := rows.Scan(
			&task.Id,
			&task.Title,
			&task.Description,
			&task.Status,
			&deadline,
			&createdAt,
			&updatedAt,
		)
		if err != nil {
			return nil, fmt.Errorf("failed to get task: %w", err)
		}

		task.Deadline = timestamppb.New(deadline)
		task.CreatedAt = timestamppb.New(createdAt)
		task.UpdatedAt = timestamppb.New(updatedAt)
		result = append(result, &task)
	}

	if err := rows.Err(); err != nil {
		return nil, fmt.Errorf("rows iteration error: %w", err)
	}

	return result, nil
}

func (t *TaskStore) UpdateStatus(ctx context.Context, id string, status string) (*pb.Task, error) {
	query := `
		UPDATE tasks
		SET status = $1, updated_at = $2
		WHERE id = $3
		RETURNING id, title, description, status, deadline, created_at, updated_at
	`

	var task pb.Task
	var deadline, created_at, updated_at time.Time

	err := t.db.Conn.QueryRow(ctx, query, status, timestamppb.Now().AsTime(), id).Scan(
		&task.Id,
		&task.Title,
		&task.Description,
		&task.Status,
		&deadline,
		&created_at,
		&updated_at,
	)
	if err != nil {
		return nil, fmt.Errorf("failed to update task: %w", err)
	}

	task.Deadline = timestamppb.New(deadline)
	task.CreatedAt = timestamppb.New(created_at)
	task.UpdatedAt = timestamppb.New(updated_at)

	return &task, nil
}

func (t *TaskStore) Delete(ctx context.Context, id string) error {
	query := `
		DELETE FROM tasks
		WHERE id = $1
	`

	_, err := t.db.Conn.Exec(ctx, query, id)
	if err != nil {
		return fmt.Errorf("failed to delete task: %w", err)
	}

	return nil
}
