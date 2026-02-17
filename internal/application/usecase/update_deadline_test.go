package usecase

import (
	"context"
	"tasker/internal/application/usecase/mocks"
	"tasker/internal/domain"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestUpdateDeadlineUC(t *testing.T) {
	t.Parallel()

	ctx := context.Background()

	testCases := GetUpdateDeadlineTestifyData()
	for _, tc := range testCases {
		tc := tc
		t.Run(tc.Name, func(t *testing.T) {
			t.Parallel()

			beforeUpdateTask := domain.DeepCopyTask(tc.InputData.Task)

			mockRepo := new(mocks.MockTestifyRepo)
			tc.SetupMock(mockRepo)
			
			usecase := NewUpdateTaskDeadlineUsecase(mockRepo)
			err := usecase.Execute(ctx, tc.InputData.Dto)
			
			if tc.WantErr != nil {
				require.Error(t, err)
				assert.ErrorIs(t, err, tc.WantErr)
				return
			}

			require.NoError(t, err)

			assert.True(t, tc.InputData.Task.Deadline.After(beforeUpdateTask.Deadline))
			assert.Equal(t, tc.InputData.Dto.NewDeadline, tc.InputData.Task.Deadline)

			mockRepo.AssertExpectations(t)
		})
	}
}