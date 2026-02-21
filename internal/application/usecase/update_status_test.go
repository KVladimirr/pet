package usecase

import (
	"context"
	"tasker/internal/application/usecase/mocks"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestUpdateStatusUC(t *testing.T) {
	t.Parallel()

	ctx := context.Background()
	
	testCases := GetUpdateStatusTestifyData()
	for _, tc := range testCases {
		tc := tc
		t.Run(tc.Name, func(t *testing.T) {
			t.Parallel()

			mockRepo := new(mocks.MockTestifyRepo)
			tc.SetupMock(mockRepo)

			usecase := NewUpdateTaskStatusUsecase(mockRepo)
			err := usecase.Execute(ctx, tc.InputData.Dto)

			if tc.WantErr != nil {
				require.Error(t, err)
				assert.ErrorIs(t, err, tc.WantErr)
				return
			}

			require.NoError(t, err)

			assert.Equal(t, tc.InputData.Dto.NewStatus, tc.InputData.Task.Status)

			mockRepo.AssertExpectations(t)
		})
	}
}