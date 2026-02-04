package test

import (
	"testing"
	"tasker/internal/domain"

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
                return
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

// func TestUpdateStatus(t *testing.T) {
// 	testCases := 
// }