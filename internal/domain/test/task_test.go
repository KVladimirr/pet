package test

import (
	"tasker/internal/domain"
	"testing"

	"github.com/google/uuid"
)

func TestNewTask(t *testing.T) {
	testCases := GetNewTaskTestData()

	for _, tc := range testCases {
		tc := tc
		t.Run(tc.Name, func(t *testing.T) {
			t.Parallel()
			result, err := domain.NewTask(tc.InputData.Title, tc.InputData.Description, tc.InputData.Deadline)
			
			if tc.WantErr != nil { 
				if err == nil {
					t.Errorf("expected error: %v, got nil", tc.WantErr)
				}
				if tc.WantErr.Error() != err.Error() {
					t.Errorf("expected error: %v, got %v", tc.WantErr, err)
				}
				return
			}

			if err != nil {
                t.Fatalf("unexpected error: %v", err)
            }

			if result == nil {
				t.Fatal("result should not be nil when no error")
			}

			if result.Title != tc.InputData.Title {
				t.Errorf("expected title: %s, got: %s", tc.InputData.Title, result.Title)
			}

			if result.Description != tc.InputData.Description {
				t.Errorf("expected description: %s, got: %s", tc.InputData.Description, result.Description)
			}

			if !result.Deadline.Equal(tc.InputData.Deadline) {
				t.Errorf("expected deadline: %v, got: %v", tc.InputData.Deadline, result.Deadline)
			}

			if result.ID == uuid.Nil {
				t.Errorf("ID should not be nil")
			}

			if result.Status != domain.TaskStatusTodo {
				t.Errorf("expected status TODO, got: %v", result.Status)
			}

			if result.CreatedAt.IsZero() {
				t.Errorf("created_at should not be zero")
			}

			if result.UpdatedAt.IsZero() {
				t.Errorf("updated_at should not be zero")
			}
		})
	}
}

func TestUpdateStatus(t *testing.T) {
	testCases := GetUpdateStatusTestCases()
	for _, tc := range testCases {
		tc := tc
		t.Run(tc.Name, func(t *testing.T) {
			t.Parallel()

			originalTask := DeepCopyTask(tc.InputData.task)

			err := tc.InputData.task.UpdateStatus(tc.InputData.newStatus)
			
			if tc.WantErr != nil {
				if err == nil {
					t.Errorf("expected error: %v, got nil", tc.WantErr)
				}
				if err.Error() != tc.WantErr.Error() {
					t.Errorf("expected error: %v, got %v", tc.WantErr, err)
				}
				return
			}

			if err != nil {
				t.Fatalf("unexpected error: %v", err)
			}

			if tc.InputData.task.Status != tc.InputData.newStatus {
				t.Errorf("expected status: %v, got %v", tc.InputData.newStatus, tc.InputData.task.Status)
			}

			if !tc.InputData.task.UpdatedAt.After(originalTask.UpdatedAt) {
				t.Errorf("updated_at should change, got  %v", tc.InputData.task.UpdatedAt)
			}

			if tc.InputData.task.ID != originalTask.ID {
				t.Errorf("ID should not change")
			}

			if tc.InputData.task.Title != originalTask.Title {
				t.Errorf("Title should not change")
			}

			if tc.InputData.task.Description != originalTask.Description {
				t.Errorf("Description should not change")
			}

			if !tc.InputData.task.Deadline.Equal(originalTask.Deadline) {
				t.Errorf("Deadline should not change")
			}

			if !tc.InputData.task.CreatedAt.Equal(originalTask.CreatedAt) {
				t.Errorf("CreatedAt should not change")
			}
		})
	}
}

