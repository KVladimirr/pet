package usecase

import (
	"context"
	"tasker/internal/application/usecase/mocks"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestGetAllTasksUC(t *testing.T) {
	t.Parallel()

	ctx := context.Background()

	testCases := GetAllTasksTestifyData()
	for _, tc := range testCases {
		tc := tc
		t.Run(tc.Name, func(t *testing.T) {
			t.Parallel()

			mockRepo := new(mocks.MockTestifyRepo)
			tc.SetupMock(mockRepo)
			
			usecase := NewGetAllTasksUsecase(mockRepo)
			taskList, err := usecase.Execute(ctx, tc.InputData.Dto)
			
			if tc.WantErr != nil {
				require.Error(t, err)
				assert.ErrorIs(t, err, tc.WantErr)
				assert.Nil(t, taskList)
				return
			}

			require.NoError(t, err)
			require.NotNil(t, taskList)

			lim := tc.InputData.Dto.Limit
			if lim > MaxLimit {
				lim = MaxLimit
			}
			assert.Len(t, taskList, int(lim))

			mockRepo.AssertExpectations(t)
		})
	}
}