package usecase

// import (
// 	"context"
// 	"errors"
// 	"tasker/internal/application/usecase/mocks"
// 	"testing"
// )

// func TestCreateTaskUC(t *testing.T) {
// 	ctx := context.Background()

// 	testCases := GetCreateTaskData()
	
// 	for _, tc := range testCases {
// 		tc := tc
// 		t.Run(tc.Name, func(t *testing.T) {
// 			t.Parallel()
			
// 			mockRepo := &mocks.MockRepositoryForCreate{}
// 			if tc.InputData.RepoErr != nil {
// 				mockRepo.SaveError = tc.InputData.RepoErr
// 			}

// 			usecase := NewCreateTaskUsecase(mockRepo)
// 			task, err := usecase.Execute(ctx, tc.InputData.Dto)
		
// 			if tc.WantErr != nil {
// 				if err == nil {
// 					t.Errorf("expected error: %v, got nil", tc.WantErr)
// 					return
// 				}
// 				if !errors.Is(err, tc.WantErr) {
// 					t.Errorf("expected error: %v, got: %v", tc.WantErr, err)
// 				}
// 				return
// 			}
		
// 			if err != nil {
// 				t.Fatalf("unexpected error: %v", err)
// 			}
		
// 			if task.Title != tc.InputData.Dto.Title {
// 				t.Errorf("expected title: %s, got: %s", tc.InputData.Dto.Title, task.Title)
// 			}
		
// 			if task.Description != tc.InputData.Dto.Description {
// 				t.Errorf("expected description: %s, got: %s", tc.InputData.Dto.Description, task.Description)
// 			}
		
// 			if !task.Deadline.Equal(tc.InputData.Dto.Deadline) {
// 				t.Errorf("expected deadline: %s, got: %s", tc.InputData.Dto.Deadline, task.Deadline)
// 			}
		
// 			if !mockRepo.SaveWasCalled {
// 				t.Errorf("Save method was not called: %v", err)
// 			}
		
// 			if mockRepo.SaveArgs == nil {
// 				t.Fatal("SaveArgs should not be nil")
// 			}
		
// 			if mockRepo.SaveArgs != task {
// 				t.Error("Expected the same task to be passed to repository")
// 			}
// 		})
// 	}

// }