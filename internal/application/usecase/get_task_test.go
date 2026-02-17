package usecase

import (
	"context"
	"tasker/internal/application/usecase/mocks"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestGetTaskUC(t *testing.T) {
	t.Parallel()

	ctx := context.Background()

	testCases := GetTaskTestifyData()
	for _, tc := range testCases {
		tc := tc
		t.Run(tc.Name, func(t *testing.T) {
			t.Parallel()

			mockRepo := new(mocks.MockTestifyRepo)
			tc.SetupMock(mockRepo)
			
			usecase := NewGetTaskByIDUsecase(mockRepo)
			task, err := usecase.Execute(ctx, tc.InputData.Dto)
			
			if tc.WantErr != nil {
				require.Error(t, err)
				assert.ErrorIs(t, err, tc.WantErr)
				assert.Nil(t, task)
				return
			}

			require.NoError(t, err)
			require.NotNil(t, task)

			mockRepo.AssertExpectations(t)
		})
	}
}