package postgres

import (
	"context"
	"database/sql"
	"errors"
	"tasker/internal/domain"
	"time"

	"github.com/google/uuid"
	// _ "github.com/lib/pq" добавить в main.go
)

var (
    ErrTaskNotFound = errors.New("task not found")
)

type PostgresTaskRepository struct {
	db *sql.DB	
}

func NewPostgresTaskRepository(db *sql.DB) *PostgresTaskRepository {
	return &PostgresTaskRepository{db: db}
}

func (p *PostgresTaskRepository) Save(ctx context.Context, task *domain.Task) error {
	query := `
        INSERT INTO tasks (id, title, description, status, deadline, created_at, updated_at)
        VALUES ($1, $2, $3, $4, $5, $6, $7)
		ON CONFLICT (id) DO UPDATE SET
            title = EXCLUDED.title,
            description = EXCLUDED.description,
            status = EXCLUDED.status,
            deadline = EXCLUDED.deadline,
            updated_at = EXCLUDED.updated_at
    `

	_, err := p.db.ExecContext(ctx, query,
		task.ID,
		task.Title,
		task.Description,
		string(task.Status),
		task.Deadline,
		task.CreatedAt,
		task.UpdatedAt,
	)

	return err
}

func (p *PostgresTaskRepository) GetByID(ctx context.Context, id uuid.UUID) (*domain.Task, error) {
	ctx, cancel := context.WithTimeout(ctx, 5*time.Second)
    defer cancel()

	query := `SELECT id, title, description, status, deadline, created_at, updated_at
		FROM tasks WHERE ID = $1
	`

	var task domain.Task

	row := p.db.QueryRowContext(ctx, query, id)

	err := row.Scan(
		&task.ID,
		&task.Title,
		&task.Description,
		&task.Status,
		&task.Deadline,
		&task.CreatedAt,
		&task.UpdatedAt,
	)
	if err != nil {
		if err == sql.ErrNoRows {
            return nil, ErrTaskNotFound
        }
		return nil, err
	}

	return &task, nil
}

func (p *PostgresTaskRepository) GetAll(ctx context.Context, limit uint, offset uint) ([]*domain.Task, error) {
	ctx, cancel := context.WithTimeout(ctx, 5*time.Second)
    defer cancel()

	query := `SELECT id, title, description, status, deadline, created_at, updated_at FROM tasks
	    ORDER BY created_at DESC
        LIMIT $1 OFFSET $2
	`

	rows, err := p.db.QueryContext(ctx, query, limit, offset)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var tasks []*domain.Task

	for rows.Next() {
		var task domain.Task

		err := rows.Scan(
			&task.ID,
			&task.Title,
			&task.Description,
			&task.Status,
			&task.Deadline,
			&task.CreatedAt,
			&task.UpdatedAt,
		)
		if err != nil {
			return nil, err
		}

		tasks = append(tasks, &task)
	}

	if err := rows.Err(); err != nil {
        return nil, err
    }

	return tasks, nil
}

func (p *PostgresTaskRepository) Delete(ctx context.Context, id uuid.UUID) error {
	query := `DELETE FROM tasks WHERE id = $1`

	result, err := p.db.ExecContext(ctx, query, id)
	if err != nil {
		return err
	}

	rowsAffected, err := result.RowsAffected()
    if err != nil {
        return err
    }

	if rowsAffected == 0 {
        return ErrTaskNotFound
    }

	return nil

}