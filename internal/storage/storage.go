package storage

import (
	"context"
	"errors"
	"fmt"
	pb "tasker/internal/task/pb"
	"time"

	"github.com/jackc/pgx/v5"
	"google.golang.org/protobuf/types/known/timestamppb"
	// "google.golang.org/protobuf/types/known/timestamppb"
)


type TaskStore struct {
	// tasks map[string]*pb.Task
	db *pgConnector
}

func NewTaskStore(dbConn *pgConnector) *TaskStore {
	return &TaskStore{db: dbConn}
}


// Методы для работы с мапой(БД)
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

// func (s *TaskStore) UpdateStatus(id string, status string) bool {

// 	task, ok := s.tasks[id]
// 	if !ok {
// 		return false
// 	}

// 	task.Status = status
// 	task.UpdatedAt = timestamppb.Now()
// 	return true
// }

// func (s *TaskStore) Delete(id string) bool {

// 	_, ok := s.tasks[id]
// 	if !ok {
// 		return false
// 	}

// 	delete(s.tasks, id)
// 	return true
// }
