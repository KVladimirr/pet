package usecase

import (
	"errors"
	"tasker/internal/application/usecase/mocks"
	"tasker/internal/domain"
	"time"

	"github.com/google/uuid"
	"github.com/stretchr/testify/mock"
)

type inputData interface {
	testifyCreateTaskData | testifyDeleteTaskData | testifyGetAllTasksData
}

type testifyTestCase[T inputData] struct{
	Name string
	InputData T
	WantErr error
	SetupMock func(*mocks.MockTestifyRepo)
}

type testifyCreateTaskData struct{
	Dto 	*CreateTaskDTO
	RepoErr	error
}

type testifyDeleteTaskData struct{
	Dto 	*DeleteTaskDTO
	RepoErr	error
}

type testifyGetAllTasksData struct{
	Dto 	*GetAllTasksDTO
	RepoErr	error
}


func GetCreateTaskTestifyData() []testifyTestCase[testifyCreateTaskData] {
	repoError := errors.New("Repo error")
	
	return []testifyTestCase[testifyCreateTaskData]{
		{
			Name: "Valid",
			InputData: testifyCreateTaskData{
				Dto: &CreateTaskDTO{
					Title: "Create testify title",
					Description: "Create testify description",
					Deadline: time.Now().Add(time.Hour),
				},
				RepoErr: nil,
			},
			WantErr: nil,
			SetupMock: func(m *mocks.MockTestifyRepo) {
				m.On("Save", mock.Anything, mock.AnythingOfType("*domain.Task")).Return(nil).Once()
			},
		},
		{
			Name: "Invalid: domain error",
			InputData: testifyCreateTaskData{
				Dto: &CreateTaskDTO{
					Title: "",
					Description: "Create test description",
					Deadline: time.Now().Add(time.Hour),
				},
				RepoErr: nil,
			},
			WantErr: domain.ErrEmptyTitle,
			SetupMock: func(m *mocks.MockTestifyRepo) {},
		},
		{
			Name: "Invalid: repo error",
			InputData: testifyCreateTaskData{
				Dto: &CreateTaskDTO{
					Title: "Title",
					Description: "Create test description",
					Deadline: time.Now().Add(time.Hour),
				},
				RepoErr: repoError,
			},
			WantErr: repoError,
			SetupMock: func(m *mocks.MockTestifyRepo) {
				m.On("Save", mock.Anything, mock.AnythingOfType("*domain.Task")).Return(repoError).Once()
			},
		},
		{
			Name: "Invalid: nil dto",
			InputData: testifyCreateTaskData{
				Dto: nil,
				RepoErr: nil,
			},
			WantErr: NilDTOError,
			SetupMock: func(m *mocks.MockTestifyRepo) {},
		},
	}
}

func GetDeleteTaskTestifyData() []testifyTestCase[testifyDeleteTaskData] {
	repoError := errors.New("Repo error")
	
	return []testifyTestCase[testifyDeleteTaskData]{
		{
			Name: "Valid",
			InputData: testifyDeleteTaskData{
				Dto: &DeleteTaskDTO{
					TaskID: uuid.New(),
				},
				RepoErr: nil,
			},
			WantErr: nil,
			SetupMock: func(m *mocks.MockTestifyRepo) {
				m.On("Delete", mock.Anything, mock.AnythingOfType("uuid.UUID")).Return(nil).Once()
			},
		},
		{
			Name: "Invalid: repo error",
			InputData: testifyDeleteTaskData{
				Dto: &DeleteTaskDTO{
					TaskID: uuid.New(),
				},
				RepoErr: repoError,
			},
			WantErr: repoError,
			SetupMock: func(m *mocks.MockTestifyRepo) {
				m.On("Delete", mock.Anything, mock.AnythingOfType("uuid.UUID")).Return(repoError).Once()
			},
		},
		{
			Name: "Invalid: nil DTO",
			InputData: testifyDeleteTaskData{
				Dto: nil,
				RepoErr: nil,
			},
			WantErr: NilDTOError,
			SetupMock: func(m *mocks.MockTestifyRepo) {},
		},
	}
}

func GetAllTasksTestifyData() []testifyTestCase[testifyGetAllTasksData] {
	repoError := errors.New("Repo error")
	
	return []testifyTestCase[testifyGetAllTasksData]{
		{
			Name: "Valid",
			InputData: testifyGetAllTasksData{
				Dto: &GetAllTasksDTO{
					Limit: 3,
					Offset: 1,
				},
				RepoErr: nil,
			},
			WantErr: nil,
			SetupMock: func(m *mocks.MockTestifyRepo) {
				expectedTasks := GenerateTasks(3)
				m.On("GetAll", mock.Anything, uint(3), uint(1)).Return(expectedTasks, nil).Once()
			},
		},
		{
			Name: "Valid: zero limit",
			InputData: testifyGetAllTasksData{
				Dto: &GetAllTasksDTO{
					Limit: 0,
					Offset: 0,
				},
				RepoErr: nil,
			},
			WantErr: nil,
			SetupMock: func(m *mocks.MockTestifyRepo) {},
		},
		{
			Name: "Valid: big limit",
			InputData: testifyGetAllTasksData{
				Dto: &GetAllTasksDTO{
					Limit: MaxLimit + 1,
					Offset: 0,
				},
				RepoErr: nil,
			},
			WantErr: nil,
			SetupMock: func(m *mocks.MockTestifyRepo) {
				expectedTasks := GenerateTasks(MaxLimit)
				m.On("GetAll", mock.Anything, uint(MaxLimit), uint(0)).Return(expectedTasks, nil).Once()
			},
		},
		{
			Name: "Invalid: nil DTO",
			InputData: testifyGetAllTasksData{
				Dto: nil,
				RepoErr: nil,
			},
			WantErr: NilDTOError,
			SetupMock: func(m *mocks.MockTestifyRepo) {},
		},
		{
			Name: "Invalid: repo error",
			InputData: testifyGetAllTasksData{
				Dto: &GetAllTasksDTO{
					Limit: 1,
					Offset: 0,
				},
				RepoErr: repoError,
			},
			WantErr: repoError,
			SetupMock: func(m *mocks.MockTestifyRepo) {
				m.On("GetAll", mock.Anything, uint(1), uint(0)).Return(nil, repoError).Once()
			},
		},
	}
}
