package test

import (
	"tasker/internal/domain"
	"time"
)

type testInput interface {
	newTaskData | updateStatusData
}

type testCase[T testInput] struct{
	Name string
	InputData T
	WantErr error
}

type newTaskData struct{
	Title 		string
	Description string
	Deadline 	time.Time
}

type updateStatusData struct {
	task 		*domain.Task
	newStatus	domain.TaskStatus
}


func GetNewTaskTestData() []testCase[newTaskData] {
	var deadline = time.Now().Add(time.Hour)
	
	return []testCase[newTaskData]{
		{
			Name: "valid",
			InputData: newTaskData{Title: "Valid", Description: "Valid desc", Deadline: deadline},
			WantErr: nil,
		},
		{
			Name: "empty description",
			InputData: newTaskData{Title: "Title", Description: "", Deadline: deadline},
			WantErr: nil,
		},
		{
			Name: "empty title",
			InputData: newTaskData{Title: "", Description: "Description", Deadline: deadline},
			WantErr: domain.ErrEmptyTitle,
		},
		{
			Name: "empty deadline",
			InputData: newTaskData{Title: "Title", Description: "Description", Deadline: time.Time{}},
			WantErr: domain.ErrEmptyDeadline,
		},
		{
			Name: "deadline in past",
			InputData: newTaskData{Title: "Title", Description: "Description", Deadline: time.Now().Add(-time.Hour)},
			WantErr: domain.ErrOldDeadline,
		},
	}
}

func GetUpdateStatusTestCases() []testCase[updateStatusData] {
	return []testCase[updateStatusData]{
		{
			Name: "Valid: TODO to INPROGRESS",
			InputData: updateStatusData{newStatus: domain.TaskStatusInProgress, task: NewTaskBuilder().task},
			WantErr: nil,
		},
	}
}
