package test

import (
	"tasker/internal/domain"
	"time"
)

type testInput interface {
	newTaskData | updateStatusData | updateTitleData | updateDescriptionData | updateDeadlineData
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

type updateTitleData struct {
	task 		*domain.Task
	newTitle	string
}

type updateDescriptionData struct {
	task 			*domain.Task
	newDescription	string
}

type updateDeadlineData struct {
	task 			*domain.Task
	newDeadline		time.Time
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
		{
			Name: "Valid: TODO to DONE",
			InputData: updateStatusData{newStatus: domain.TaskStatusDone, task: NewTaskBuilder().task},
			WantErr: nil,
		},
		{
			Name: "Valid: TODO to CANCELLED",
			InputData: updateStatusData{newStatus: domain.TaskStatusCanceled, task: NewTaskBuilder().task},
			WantErr: nil,
		},
		{
			Name: "Valid: TODO to HOLD",
			InputData: updateStatusData{newStatus: domain.TaskStatusHold, task: NewTaskBuilder().task},
			WantErr: nil,
		},
		{
			Name: "Valid: INPROGRESS to TODO",
			InputData: updateStatusData{newStatus: domain.TaskStatusTodo, task: NewTaskBuilder().Status(domain.TaskStatusInProgress).task},
			WantErr: nil,
		},
		{
			Name: "Valid: INPROGRESS to DONE",
			InputData: updateStatusData{newStatus: domain.TaskStatusDone, task: NewTaskBuilder().Status(domain.TaskStatusInProgress).task},
			WantErr: nil,
		},
		{
			Name: "Valid: INPROGRESS to CANCELLED",
			InputData: updateStatusData{newStatus: domain.TaskStatusCanceled, task: NewTaskBuilder().Status(domain.TaskStatusInProgress).task},
			WantErr: nil,
		},
		{
			Name: "Valid: INPROGRESS to HOLD",
			InputData: updateStatusData{newStatus: domain.TaskStatusHold, task: NewTaskBuilder().Status(domain.TaskStatusInProgress).task},
			WantErr: nil,
		},
		{
			Name: "Valid: HOLD to TODO",
			InputData: updateStatusData{newStatus: domain.TaskStatusTodo, task: NewTaskBuilder().Status(domain.TaskStatusHold).task},
			WantErr: nil,
		},
		{
			Name: "Valid: HOLD to INPROGRESS",
			InputData: updateStatusData{newStatus: domain.TaskStatusInProgress, task: NewTaskBuilder().Status(domain.TaskStatusHold).task},
			WantErr: nil,
		},
		{
			Name: "Valid: HOLD to DONE",
			InputData: updateStatusData{newStatus: domain.TaskStatusDone, task: NewTaskBuilder().Status(domain.TaskStatusHold).task},
			WantErr: nil,
		},
		{
			Name: "Valid: HOLD to CANCELLED",
			InputData: updateStatusData{newStatus: domain.TaskStatusCanceled, task: NewTaskBuilder().Status(domain.TaskStatusHold).task},
			WantErr: nil,
		},
		{
			Name: "Invalid: CANCELLED to TODO",
			InputData: updateStatusData{newStatus: domain.TaskStatusTodo, task: NewTaskBuilder().Status(domain.TaskStatusCanceled).task},
			WantErr: domain.ErrCannotEditTask,
		},
		{
			Name: "Invalid: CANCELLED to INPROGRESS",
			InputData: updateStatusData{newStatus: domain.TaskStatusInProgress, task: NewTaskBuilder().Status(domain.TaskStatusCanceled).task},
			WantErr: domain.ErrCannotEditTask,
		},
		{
			Name: "Invalid: CANCELLED to DONE",
			InputData: updateStatusData{newStatus: domain.TaskStatusDone, task: NewTaskBuilder().Status(domain.TaskStatusCanceled).task},
			WantErr: domain.ErrCannotEditTask,
		},
		{
			Name: "Invalid: CANCELLED to HOLD",
			InputData: updateStatusData{newStatus: domain.TaskStatusHold, task: NewTaskBuilder().Status(domain.TaskStatusCanceled).task},
			WantErr: domain.ErrCannotEditTask,
		},
		{
			Name: "Invalid: DONE to TODO",
			InputData: updateStatusData{newStatus: domain.TaskStatusTodo, task: NewTaskBuilder().Status(domain.TaskStatusDone).task},
			WantErr: domain.ErrCannotEditTask,
		},
		{
			Name: "Invalid: DONE to INPROGRESS",
			InputData: updateStatusData{newStatus: domain.TaskStatusInProgress, task: NewTaskBuilder().Status(domain.TaskStatusDone).task},
			WantErr: domain.ErrCannotEditTask,
		},
		{
			Name: "Invalid: DONE to CANCELLED",
			InputData: updateStatusData{newStatus: domain.TaskStatusCanceled, task: NewTaskBuilder().Status(domain.TaskStatusDone).task},
			WantErr: domain.ErrCannotEditTask,
		},
		{
			Name: "Invalid: DONE to HOLD",
			InputData: updateStatusData{newStatus: domain.TaskStatusHold, task: NewTaskBuilder().Status(domain.TaskStatusDone).task},
			WantErr: domain.ErrCannotEditTask,
		},
		{
			Name: "Invalid: Overdue task update",
			InputData: updateStatusData{newStatus: domain.TaskStatusInProgress, task: NewTaskBuilder().Deadline(time.Now().Add(-time.Hour)).task},
			WantErr: domain.ErrCannotEditTask,
		},
	}
}

func GetUpdateTitleTestCases() []testCase[updateTitleData] {
	return []testCase[updateTitleData]{
		{
			Name: "Valid",
			InputData: updateTitleData{newTitle: "Another title", task: NewTaskBuilder().task},
			WantErr: nil,
		},
		{
			Name: "Invalid: in DONE status",
			InputData: updateTitleData{newTitle: "Another title", task: NewTaskBuilder().Status(domain.TaskStatusDone).task},
			WantErr: domain.ErrCannotEditTask,
		},
		{
			Name: "Invalid: in CANCELLED status",
			InputData: updateTitleData{newTitle: "Another title", task: NewTaskBuilder().Status(domain.TaskStatusCanceled).task},
			WantErr: domain.ErrCannotEditTask,
		},
		{
			Name: "Invalid: Overdue task",
			InputData: updateTitleData{newTitle: "Another title", task: NewTaskBuilder().Deadline(time.Now().Add(-time.Hour)).task},
			WantErr: domain.ErrCannotEditTask,
		},
		{
			Name: "Invalid: empty title",
			InputData: updateTitleData{newTitle: "", task: NewTaskBuilder().task},
			WantErr: domain.ErrEmptyTitle,
		},
	}
}

func GetUpdateDescriptionTestCases() []testCase[updateDescriptionData] {
	return []testCase[updateDescriptionData]{
		{
			Name: "Valid",
			InputData: updateDescriptionData{newDescription: "Another description", task: NewTaskBuilder().task},
			WantErr: nil,
		},
		{
			Name: "Vaild: empty description",
			InputData: updateDescriptionData{newDescription: "", task: NewTaskBuilder().task},
			WantErr: nil,
		},
		{
			Name: "Invalid: in DONE status",
			InputData: updateDescriptionData{newDescription: "Another description", task: NewTaskBuilder().Status(domain.TaskStatusDone).task},
			WantErr: domain.ErrCannotEditTask,
		},
		{
			Name: "Invalid: in CANCELLED status",
			InputData: updateDescriptionData{newDescription: "Another description", task: NewTaskBuilder().Status(domain.TaskStatusCanceled).task},
			WantErr: domain.ErrCannotEditTask,
		},
		{
			Name: "Invalid: Overdue task",
			InputData: updateDescriptionData{newDescription: "Another description", task: NewTaskBuilder().Deadline(time.Now().Add(-time.Hour)).task},
			WantErr: domain.ErrCannotEditTask,
		},
	}
}

func GetUpdateDeadlineTestCases() []testCase[updateDeadlineData] {
	now := time.Now()

	return []testCase[updateDeadlineData]{
		{
			Name: "Valid: increase",
			InputData: updateDeadlineData{newDeadline: now.Add(2*time.Hour), task: NewTaskBuilder().task},
			WantErr: nil,
		},
		{
			Name: "Vaild: decrease",
			InputData: updateDeadlineData{newDeadline: now.Add(30*time.Minute), task: NewTaskBuilder().task},
			WantErr: nil,
		},
		{
			Name: "Invalid: in DONE status",
			InputData: updateDeadlineData{newDeadline: now.Add(2*time.Hour), task: NewTaskBuilder().Status(domain.TaskStatusDone).task},
			WantErr: domain.ErrCannotEditTask,
		},
		{
			Name: "Invalid: in CANCELLED status",
			InputData: updateDeadlineData{newDeadline: now.Add(2*time.Hour), task: NewTaskBuilder().Status(domain.TaskStatusCanceled).task},
			WantErr: domain.ErrCannotEditTask,
		},
		{
			Name: "Invalid: Overdue task",
			InputData: updateDeadlineData{newDeadline: now.Add(2*time.Hour), task: NewTaskBuilder().Deadline(time.Now().Add(-time.Hour)).task},
			WantErr: domain.ErrCannotEditTask,
		},
		{
			Name: "Invalid: Deadline in past",
			InputData: updateDeadlineData{newDeadline: now.Add(-2*time.Hour), task: NewTaskBuilder().task},
			WantErr: domain.ErrOldDeadline,
		},
		{
			Name: "Invalid: nil Deadline",
			InputData: updateDeadlineData{newDeadline: time.Time{}, task: NewTaskBuilder().task},
			WantErr: domain.ErrEmptyDeadline,
		},
	}
}
