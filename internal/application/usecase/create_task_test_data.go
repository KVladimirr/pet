package usecase

import (
	"errors"
	"tasker/internal/domain"
	"time"
)


type testCase struct{
	Name string
	InputData createTaskData
	WantErr error
}

type createTaskData struct{
	Dto 	*CreateTaskDTO
	RepoErr	error
}


func GetCreateTaskData() []testCase {
	repoError := errors.New("Repo error")

	return []testCase{
		{
			Name: "Valid",
			InputData: createTaskData{
				Dto: &CreateTaskDTO{
					Title: "Create test title",
					Description: "Create test description",
					Deadline: time.Now().Add(time.Hour),
				},
				RepoErr: nil,
			},
			WantErr: nil,
		},
		{
			Name: "Invalid: domain error",
			InputData: createTaskData{
				Dto: &CreateTaskDTO{
					Title: "",
					Description: "Create test description",
					Deadline: time.Now().Add(time.Hour),
				},
				RepoErr: nil,
			},
			WantErr: domain.ErrEmptyTitle,
		},
		{
			Name: "Invalid: repo error",
			InputData: createTaskData{
				Dto: &CreateTaskDTO{
					Title: "Title",
					Description: "Create test description",
					Deadline: time.Now().Add(time.Hour),
				},
				RepoErr: repoError,
			},
			WantErr: repoError,
		},
		{
			Name: "Invalid: nil dto",
			InputData: createTaskData{
				Dto: nil,
				RepoErr: nil,
			},
			WantErr: NilDTOError,
		},
	}
}