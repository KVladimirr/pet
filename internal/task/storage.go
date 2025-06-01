package task

import (
	"sync"
	pb "tasker/internal/task/pb"
	"time"

	"google.golang.org/protobuf/types/known/timestamppb"
)


type TaskStore struct {
	mu sync.RWMutex
	tasks map[string]*pb.Task
}

func NewTaskStore() *TaskStore {

	return &TaskStore{
		tasks: make(map[string]*pb.Task),
	}
}


// Методы для работы с мапой(БД)
func (s *TaskStore) Create(task *pb.Task) {
	s.mu.Lock()
	defer s.mu.Unlock()
	s.tasks[task.Id] = task
}

func (s *TaskStore) Get(id string) (*pb.Task, bool) {
	s.mu.RLock()
	defer s.mu.RUnlock()
	task, ok := s.tasks[id]
	return task, ok
}

func (s *TaskStore) List() []*pb.Task {
	s.mu.RLock()
	defer s.mu.RUnlock()
	task_list := make([]*pb.Task, 0, len(s.tasks))
	for _, t := range s.tasks {
		task_list = append(task_list, t)
	}
	return task_list
}

func (s *TaskStore) UpdateStatus(id string, status string) bool {
	s.mu.Lock()
	defer s.mu.Unlock()

	task, ok := s.tasks[id]
	if !ok {
		return false
	}

	task.Status = status
	task.UpdatedAt = timestamppb.New(time.Now())
	return true
}

func (s *TaskStore) Delete(id string) bool {
	s.mu.Lock()
	defer s.mu.Unlock()

	_, ok := s.tasks[id]
	if !ok {
		return false
	}

	delete(s.tasks, id)
	return true
}
