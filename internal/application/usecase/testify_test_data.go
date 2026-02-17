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
	testifyCreateTaskData | testifyDeleteTaskData | testifyGetAllTasksData | testifyGetTaskData |
	testifyUpdateDeadlineTaskData | testifyUpdateDescriptionTaskData
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

type testifyGetTaskData struct{
	Dto 	*GetTaskByIDDTO
	RepoErr	error
}

type testifyUpdateDeadlineTaskData struct{
	Dto		*UpdateTaskDeadlineDTO
	Task	*domain.Task
	RepoErr	error
}

type testifyUpdateDescriptionTaskData struct{
	Dto		*UpdateTaskDescriptionDTO
	Task	*domain.Task
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

func GetTaskTestifyData() []testifyTestCase[testifyGetTaskData] {
	repoError := errors.New("Repo error")
	
	return []testifyTestCase[testifyGetTaskData]{
		{
			Name: "Valid",
			InputData: testifyGetTaskData{
				Dto: &GetTaskByIDDTO{
					TaskID: uuid.New(),
				},
				RepoErr: nil,
			},
			WantErr: nil,
			SetupMock: func(m *mocks.MockTestifyRepo) {
				expectedTask := GenerateTasks(1)[0]
				m.On("GetByID", mock.Anything, mock.AnythingOfType("uuid.UUID")).Return(expectedTask, nil).Once()
			},
		},
		{
			Name: "Invalid: nil DTO",
			InputData: testifyGetTaskData{
				Dto: nil,
				RepoErr: nil,
			},
			WantErr: NilDTOError,
			SetupMock: func(m *mocks.MockTestifyRepo) {},
		},
		{
			Name: "Invalid: repo error",
			InputData: testifyGetTaskData{
				Dto: &GetTaskByIDDTO{
					TaskID: uuid.New(),
				},
				RepoErr: repoError,
			},
			WantErr: repoError,
			SetupMock: func(m *mocks.MockTestifyRepo) {
				m.On("GetByID", mock.Anything, mock.AnythingOfType("uuid.UUID")).Return(nil, repoError).Once()
			},
		},
		{
			Name: "Invalid: nil taskID",
			InputData: testifyGetTaskData{
				Dto: &GetTaskByIDDTO{
					TaskID: uuid.UUID{},
				},
				RepoErr: nil,
			},
			WantErr: NilIDError,
			SetupMock: func(m *mocks.MockTestifyRepo) {},
		},
	}
}

func GetUpdateDeadlineTestifyData() []testifyTestCase[testifyUpdateDeadlineTaskData] {
	repoError := errors.New("Repo error")
	taskForValidCase := domain.NewTaskBuilder().ID(uuid.New()).Task
	taskForGetTaskErrorCase := domain.NewTaskBuilder().Task
	taskForDomainErrorCase := domain.NewTaskBuilder().Deadline(time.Time{}).Task
	taskForSaveErrorCase := domain.NewTaskBuilder().Task
	taskForNilDTOErrorCase := domain.NewTaskBuilder().Task
	
	return []testifyTestCase[testifyUpdateDeadlineTaskData]{
		{
			Name: "Valid",
			InputData: testifyUpdateDeadlineTaskData{
				Dto: &UpdateTaskDeadlineDTO{
					TaskID: taskForValidCase.ID,
					NewDeadline: time.Now().Add(2*time.Hour),
				},
				Task: taskForValidCase,
				RepoErr: nil,
			},
			WantErr: nil,
			SetupMock: func(m *mocks.MockTestifyRepo) {
				m.On("GetByID", mock.Anything, taskForValidCase.ID).Return(taskForValidCase, nil).Once()
				m.On("Save", mock.Anything, taskForValidCase).Return(nil).Once()
			},
		},
		{
			Name: "Invalid: GetByID error",
			InputData: testifyUpdateDeadlineTaskData{
				Dto: &UpdateTaskDeadlineDTO{
					TaskID: taskForGetTaskErrorCase.ID,
					NewDeadline: time.Now().Add(2*time.Hour),
				},
				Task: taskForGetTaskErrorCase,
				RepoErr: repoError,
			},
			WantErr: repoError,
			SetupMock: func(m *mocks.MockTestifyRepo) {
				m.On("GetByID", mock.Anything, taskForGetTaskErrorCase.ID).Return(nil, repoError).Once()
			},
		},
		{
			Name: "Invalid: domain error",
			InputData: testifyUpdateDeadlineTaskData{
				Dto: &UpdateTaskDeadlineDTO{
					TaskID: taskForDomainErrorCase.ID,
					NewDeadline: time.Now().Add(2*time.Hour),
				},
				Task: taskForDomainErrorCase,
				RepoErr: domain.ErrCannotEditTask,
			},
			WantErr: domain.ErrCannotEditTask,
			SetupMock: func(m *mocks.MockTestifyRepo) {
				m.On("GetByID", mock.Anything, taskForDomainErrorCase.ID).Return(taskForDomainErrorCase, nil).Once()
			},
		},
		{
			Name: "Invalid: Save error",
			InputData: testifyUpdateDeadlineTaskData{
				Dto: &UpdateTaskDeadlineDTO{
					TaskID: taskForSaveErrorCase.ID,
					NewDeadline: time.Now().Add(2*time.Hour),
				},
				Task: taskForSaveErrorCase,
				RepoErr: repoError,
			},
			WantErr: repoError,
			SetupMock: func(m *mocks.MockTestifyRepo) {
				m.On("GetByID", mock.Anything, taskForSaveErrorCase.ID).Return(taskForSaveErrorCase, nil).Once()
				m.On("Save", mock.Anything, taskForSaveErrorCase).Return(repoError).Once()
			},
		},
		{
			Name: "Invalid: nil DTO",
			InputData: testifyUpdateDeadlineTaskData{
				Dto: nil,
				Task: taskForNilDTOErrorCase,
				RepoErr: nil,
			},
			WantErr: NilDTOError,
			SetupMock: func(m *mocks.MockTestifyRepo) {},
		},
	}
}

func GetUpdateDescriptionTestifyData() []testifyTestCase[testifyUpdateDescriptionTaskData] {
	repoError := errors.New("Repo error")
	taskForValidCase := domain.NewTaskBuilder().Task
	taskForNilDTOErrorCase := domain.NewTaskBuilder().Task
	taskForGetTaskErrorCase := domain.NewTaskBuilder().Task
	taskForDomainErrorCase := domain.NewTaskBuilder().Status(domain.TaskStatusCanceled).Task
	taskForSaveErrorCase := domain.NewTaskBuilder().Task
	
	return []testifyTestCase[testifyUpdateDescriptionTaskData]{
		{
			Name: "Valid",
			InputData: testifyUpdateDescriptionTaskData{
				Dto: &UpdateTaskDescriptionDTO{
					TaskID: taskForValidCase.ID,
					NewDescription: "Updated description",
				},
				Task: taskForValidCase,
				RepoErr: nil,
			},
			WantErr: nil,
			SetupMock: func(m *mocks.MockTestifyRepo) {
				m.On("GetByID", mock.Anything, taskForValidCase.ID).Return(taskForValidCase, nil).Once()
				m.On("Save", mock.Anything, taskForValidCase).Return(nil).Once()
			},
		},
		{
			Name: "Invalid: nil DTO",
			InputData: testifyUpdateDescriptionTaskData{
				Dto: nil,
				Task: taskForNilDTOErrorCase,
				RepoErr: nil,
			},
			WantErr: NilDTOError,
			SetupMock: func(m *mocks.MockTestifyRepo) {},
		},
		{
			Name: "Invalid: GetByID error",
			InputData: testifyUpdateDescriptionTaskData{
				Dto: &UpdateTaskDescriptionDTO{
					TaskID: taskForGetTaskErrorCase.ID,
					NewDescription: "Updated description",
				},
				Task: taskForGetTaskErrorCase,
				RepoErr: repoError,
			},
			WantErr: repoError,
			SetupMock: func(m *mocks.MockTestifyRepo) {
				m.On("GetByID", mock.Anything, taskForGetTaskErrorCase.ID).Return(nil, repoError).Once()
			},
		},
		{
			Name: "Invalid: domain error",
			InputData: testifyUpdateDescriptionTaskData{
				Dto: &UpdateTaskDescriptionDTO{
					TaskID: taskForDomainErrorCase.ID,
					NewDescription: "Updated",
				},
				Task: taskForDomainErrorCase,
				RepoErr: domain.ErrCannotEditTask,
			},
			WantErr: domain.ErrCannotEditTask,
			SetupMock: func(m *mocks.MockTestifyRepo) {
				m.On("GetByID", mock.Anything, taskForDomainErrorCase.ID).Return(taskForDomainErrorCase, nil).Once()
			},
		},
		{
			Name: "Invalid: Save error",
			InputData: testifyUpdateDescriptionTaskData{
				Dto: &UpdateTaskDescriptionDTO{
					TaskID: taskForSaveErrorCase.ID,
					NewDescription: "Updated", 
				},
				Task: taskForSaveErrorCase,
				RepoErr: repoError,
			},
			WantErr: repoError,
			SetupMock: func(m *mocks.MockTestifyRepo) {
				m.On("GetByID", mock.Anything, taskForSaveErrorCase.ID).Return(taskForSaveErrorCase, nil).Once()
				m.On("Save", mock.Anything, taskForSaveErrorCase).Return(repoError).Once()
			},
		},
	}
}
