package domain

import (
	"testing"
	"time"

	"github.com/google/uuid"
)

type createNewTaskData struct{
	title 		string
	description string
	deadline 	time.Time
}

var deadline = time.Now().Add(time.Hour)

func TestNewTask(t *testing.T) {
	testCases := []struct{
		name string
		inputData createNewTaskData
		wantErr error
	}{
		{
			name: "valid",
			inputData: createNewTaskData{title: "Valid", description: "Valid desc", deadline: deadline},
			wantErr: nil,
		},
		{
			name: "empty description",
			inputData: createNewTaskData{title: "Title", description: "", deadline: deadline},
			wantErr: nil,
		},
		{
			name: "empty title",
			inputData: createNewTaskData{title: "", description: "Description", deadline: deadline},
			wantErr: ErrEmptyTitle,
		},
		{
			name: "empty deadline",
			inputData: createNewTaskData{title: "Title", description: "Description", deadline: time.Time{}},
			wantErr: ErrEmptyDeadline,
		},
		{
			name: "deadline in past",
			inputData: createNewTaskData{title: "Title", description: "Description", deadline: time.Now().Add(-time.Hour)},
			wantErr: ErrOldDeadline,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			result, err := NewTask(tc.inputData.title, tc.inputData.description, tc.inputData.deadline)
			
			if tc.wantErr != nil { 
				if err == nil {
					t.Errorf("expected error: %v, got nil", tc.wantErr)
				}
				if tc.wantErr.Error() != err.Error() {
					t.Errorf("expected error: %v, got %v", tc.wantErr, err)
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

			if result.Title != tc.inputData.title {
				t.Errorf("expected title: %s, got: %s", tc.inputData.title, result.Title)
			}

			if result.Description != tc.inputData.description {
				t.Errorf("expected description: %s, got: %s", tc.inputData.description, result.Description)
			}

			if !result.Deadline.Equal(tc.inputData.deadline) {
				t.Errorf("expected deadline: %v, got: %v", tc.inputData.deadline, result.Deadline)
			}

			if result.ID == uuid.Nil {
				t.Errorf("ID should not be nil")
			}

			if result.Status != TaskStatusTodo {
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