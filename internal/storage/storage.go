package storage

import (
	"context"
	"fmt"
	pb "tasker/internal/task/pb"

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

// func (s *TaskStore) Get(id string) (*pb.Task, bool) {
// 	task, ok := s.tasks[id]
// 	return task, ok
// }

// func (s *TaskStore) List() []*pb.Task {
// 	task_list := make([]*pb.Task, 0, len(s.tasks))
// 	for _, t := range s.tasks {
// 		task_list = append(task_list, t)
// 	}
// 	return task_list
// }

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
