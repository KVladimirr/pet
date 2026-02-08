package domain

import (
	"testing"

	"github.com/google/uuid"
)

func TestNewTask(t *testing.T) {
	testCases := GetNewTaskTestData()

	for _, tc := range testCases {
		tc := tc
		t.Run(tc.Name, func(t *testing.T) {
			t.Parallel()
			result, err := NewTask(tc.InputData.Title, tc.InputData.Description, tc.InputData.Deadline)
			
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

func TestUpdateTitle(t *testing.T) {
	testCases := GetUpdateTitleTestCases()
	for _, tc := range testCases {
		tc := tc
		t.Run(tc.Name, func(t *testing.T) {
			t.Parallel()

			originalTask := DeepCopyTask(tc.InputData.task)

			err := tc.InputData.task.UpdateTitle(tc.InputData.newTitle)

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

			if tc.InputData.task.Title != tc.InputData.newTitle {
				t.Errorf("expected title to be: %v, got %v", tc.InputData.newTitle, tc.InputData.task.Title)
			}

			if !tc.InputData.task.UpdatedAt.After(originalTask.UpdatedAt) {
				t.Errorf("updated_at should change, got  %v", tc.InputData.task.UpdatedAt)
			}

			if tc.InputData.task.ID != originalTask.ID {
				t.Errorf("ID should not change")
			}

			if tc.InputData.task.Status != originalTask.Status {
				t.Errorf("Status should not change")
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

func TestUpdateDescription(t *testing.T) {
	testCases := GetUpdateDescriptionTestCases()
	for _, tc := range testCases {
		tc := tc
		t.Run(tc.Name, func(t *testing.T) {
			t.Parallel()

			originalTask := DeepCopyTask(tc.InputData.task)

			err := tc.InputData.task.UpdateDescription(tc.InputData.newDescription)

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

			if tc.InputData.task.Description != tc.InputData.newDescription {
				t.Errorf("expected description to be: %v, got %v", tc.InputData.newDescription, tc.InputData.task.Description)
			}

			if !tc.InputData.task.UpdatedAt.After(originalTask.UpdatedAt) {
				t.Errorf("updated_at should change, got  %v", tc.InputData.task.UpdatedAt)
			}

			if tc.InputData.task.ID != originalTask.ID {
				t.Errorf("ID should not change")
			}

			if tc.InputData.task.Status != originalTask.Status {
				t.Errorf("Status should not change")
			}

			if tc.InputData.task.Title != originalTask.Title {
				t.Errorf("Title should not change")
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

func TestUpdateDeadline(t *testing.T) {
	testCases := GetUpdateDeadlineTestCases()
	for _, tc := range testCases {
		tc := tc
		t.Run(tc.Name, func(t *testing.T) {
			t.Parallel()

			originalTask := DeepCopyTask(tc.InputData.task)

			err := tc.InputData.task.UpdateDeadline(tc.InputData.newDeadline)

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

			if !tc.InputData.task.Deadline.Equal(tc.InputData.newDeadline) {
				t.Errorf("expected deadline to be: %v, got %v", tc.InputData.newDeadline, tc.InputData.task.Deadline)
			}

			if !tc.InputData.task.UpdatedAt.After(originalTask.UpdatedAt) {
				t.Errorf("updated_at should change, got  %v", tc.InputData.task.UpdatedAt)
			}

			if tc.InputData.task.ID != originalTask.ID {
				t.Errorf("ID should not change")
			}

			if tc.InputData.task.Status != originalTask.Status {
				t.Errorf("Status should not change")
			}

			if tc.InputData.task.Title != originalTask.Title {
				t.Errorf("Title should not change")
			}

			if tc.InputData.task.Description != originalTask.Description {
				t.Errorf("Description should not change")
			}

			if !tc.InputData.task.CreatedAt.Equal(originalTask.CreatedAt) {
				t.Errorf("CreatedAt should not change")
			}
		})
	}
}

func TestIsOverdue(t *testing.T) {
	testCases := GetIsOverdueTestCases()
	for _, tc := range testCases {
		tc := tc
		t.Run(tc.Name, func(t *testing.T) {
			t.Parallel()

			isOverdue := tc.InputData.task.IsOverdue()

			if isOverdue != tc.InputData.expectedResponse {
				t.Errorf("expected isOverdue to be: %v, got %v", tc.InputData.expectedResponse, isOverdue)
			}
		})
	}
}

func TestIsActive(t *testing.T) {
	testCases := GetIsActiveTestCases()
	for _, tc := range testCases {
		tc := tc
		t.Run(tc.Name, func(t *testing.T) {
			t.Parallel()

			isActive := tc.InputData.task.IsActive()

			if isActive != tc.InputData.expectedResponse {
				t.Errorf("expected isActive to be: %v, got %v", tc.InputData.expectedResponse, isActive)
			}
		})
	}
}

func TestIsFinished(t *testing.T) {
	testCases := GetIsFinishedTestCases()
	for _, tc := range testCases {
		tc := tc
		t.Run(tc.Name, func(t *testing.T) {
			t.Parallel()

			isFinished := tc.InputData.task.IsFinished()

			if isFinished != tc.InputData.expectedResponse {
				t.Errorf("expected isFinished to be: %v, got %v", tc.InputData.expectedResponse, isFinished)
			}
		})
	}
}

func TestCanEdit(t *testing.T) {
	testCases := GetCanEditTestCases()
	for _, tc := range testCases {
		tc := tc
		t.Run(tc.Name, func(t *testing.T) {
			t.Parallel()

			canEdit := tc.InputData.task.CanEdit()

			if canEdit != tc.InputData.expectedResponse {
				t.Errorf("expected canEdit to be: %v, got %v", tc.InputData.expectedResponse, canEdit)
			}
		})
	}
}

