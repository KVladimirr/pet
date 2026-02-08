package domain

import (
	"time"
)

type testInput interface {
	newTaskData | updateStatusData | updateTitleData | updateDescriptionData | updateDeadlineData | 
	isOverdueData | isActiveData | isFinishedData | canEditData
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
	task 		*Task
	newStatus	TaskStatus
}

type updateTitleData struct {
	task 		*Task
	newTitle	string
}

type updateDescriptionData struct {
	task 			*Task
	newDescription	string
}

type updateDeadlineData struct {
	task 			*Task
	newDeadline		time.Time
}

type isOverdueData struct {
	task 			*Task
	expectedResponse bool
}

type isActiveData struct {
	task 			*Task
	expectedResponse bool
}

type isFinishedData struct {
	task 			*Task
	expectedResponse bool
}

type canEditData struct {
	task 			*Task
	expectedResponse bool
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
			WantErr: ErrEmptyTitle,
		},
		{
			Name: "empty deadline",
			InputData: newTaskData{Title: "Title", Description: "Description", Deadline: time.Time{}},
			WantErr: ErrEmptyDeadline,
		},
		{
			Name: "deadline in past",
			InputData: newTaskData{Title: "Title", Description: "Description", Deadline: time.Now().Add(-time.Hour)},
			WantErr: ErrOldDeadline,
		},
	}
}

func GetUpdateStatusTestCases() []testCase[updateStatusData] {
	return []testCase[updateStatusData]{
		{
			Name: "Valid: TODO to INPROGRESS",
			InputData: updateStatusData{newStatus: TaskStatusInProgress, task: NewTaskBuilder().task},
			WantErr: nil,
		},
		{
			Name: "Valid: TODO to DONE",
			InputData: updateStatusData{newStatus: TaskStatusDone, task: NewTaskBuilder().task},
			WantErr: nil,
		},
		{
			Name: "Valid: TODO to CANCELLED",
			InputData: updateStatusData{newStatus: TaskStatusCanceled, task: NewTaskBuilder().task},
			WantErr: nil,
		},
		{
			Name: "Valid: TODO to HOLD",
			InputData: updateStatusData{newStatus: TaskStatusHold, task: NewTaskBuilder().task},
			WantErr: nil,
		},
		{
			Name: "Valid: INPROGRESS to TODO",
			InputData: updateStatusData{newStatus: TaskStatusTodo, task: NewTaskBuilder().Status(TaskStatusInProgress).task},
			WantErr: nil,
		},
		{
			Name: "Valid: INPROGRESS to DONE",
			InputData: updateStatusData{newStatus: TaskStatusDone, task: NewTaskBuilder().Status(TaskStatusInProgress).task},
			WantErr: nil,
		},
		{
			Name: "Valid: INPROGRESS to CANCELLED",
			InputData: updateStatusData{newStatus: TaskStatusCanceled, task: NewTaskBuilder().Status(TaskStatusInProgress).task},
			WantErr: nil,
		},
		{
			Name: "Valid: INPROGRESS to HOLD",
			InputData: updateStatusData{newStatus: TaskStatusHold, task: NewTaskBuilder().Status(TaskStatusInProgress).task},
			WantErr: nil,
		},
		{
			Name: "Valid: HOLD to TODO",
			InputData: updateStatusData{newStatus: TaskStatusTodo, task: NewTaskBuilder().Status(TaskStatusHold).task},
			WantErr: nil,
		},
		{
			Name: "Valid: HOLD to INPROGRESS",
			InputData: updateStatusData{newStatus: TaskStatusInProgress, task: NewTaskBuilder().Status(TaskStatusHold).task},
			WantErr: nil,
		},
		{
			Name: "Valid: HOLD to DONE",
			InputData: updateStatusData{newStatus: TaskStatusDone, task: NewTaskBuilder().Status(TaskStatusHold).task},
			WantErr: nil,
		},
		{
			Name: "Valid: HOLD to CANCELLED",
			InputData: updateStatusData{newStatus: TaskStatusCanceled, task: NewTaskBuilder().Status(TaskStatusHold).task},
			WantErr: nil,
		},
		{
			Name: "Invalid: CANCELLED to TODO",
			InputData: updateStatusData{newStatus: TaskStatusTodo, task: NewTaskBuilder().Status(TaskStatusCanceled).task},
			WantErr: ErrCannotEditTask,
		},
		{
			Name: "Invalid: CANCELLED to INPROGRESS",
			InputData: updateStatusData{newStatus: TaskStatusInProgress, task: NewTaskBuilder().Status(TaskStatusCanceled).task},
			WantErr: ErrCannotEditTask,
		},
		{
			Name: "Invalid: CANCELLED to DONE",
			InputData: updateStatusData{newStatus: TaskStatusDone, task: NewTaskBuilder().Status(TaskStatusCanceled).task},
			WantErr: ErrCannotEditTask,
		},
		{
			Name: "Invalid: CANCELLED to HOLD",
			InputData: updateStatusData{newStatus: TaskStatusHold, task: NewTaskBuilder().Status(TaskStatusCanceled).task},
			WantErr: ErrCannotEditTask,
		},
		{
			Name: "Invalid: DONE to TODO",
			InputData: updateStatusData{newStatus: TaskStatusTodo, task: NewTaskBuilder().Status(TaskStatusDone).task},
			WantErr:ErrCannotEditTask,
		},
		{
			Name: "Invalid: DONE to INPROGRESS",
			InputData: updateStatusData{newStatus: TaskStatusInProgress, task: NewTaskBuilder().Status(TaskStatusDone).task},
			WantErr: ErrCannotEditTask,
		},
		{
			Name: "Invalid: DONE to CANCELLED",
			InputData: updateStatusData{newStatus: TaskStatusCanceled, task: NewTaskBuilder().Status(TaskStatusDone).task},
			WantErr: ErrCannotEditTask,
		},
		{
			Name: "Invalid: DONE to HOLD",
			InputData: updateStatusData{newStatus: TaskStatusHold, task: NewTaskBuilder().Status(TaskStatusDone).task},
			WantErr: ErrCannotEditTask,
		},
		{
			Name: "Invalid: Overdue task update",
			InputData: updateStatusData{newStatus: TaskStatusInProgress, task: NewTaskBuilder().Deadline(time.Now().Add(-time.Hour)).task},
			WantErr: ErrCannotEditTask,
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
			InputData: updateTitleData{newTitle: "Another title", task: NewTaskBuilder().Status(TaskStatusDone).task},
			WantErr: ErrCannotEditTask,
		},
		{
			Name: "Invalid: in CANCELLED status",
			InputData: updateTitleData{newTitle: "Another title", task: NewTaskBuilder().Status(TaskStatusCanceled).task},
			WantErr: ErrCannotEditTask,
		},
		{
			Name: "Invalid: Overdue task",
			InputData: updateTitleData{newTitle: "Another title", task: NewTaskBuilder().Deadline(time.Now().Add(-time.Hour)).task},
			WantErr: ErrCannotEditTask,
		},
		{
			Name: "Invalid: empty title",
			InputData: updateTitleData{newTitle: "", task: NewTaskBuilder().task},
			WantErr: ErrEmptyTitle,
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
			InputData: updateDescriptionData{newDescription: "Another description", task: NewTaskBuilder().Status(TaskStatusDone).task},
			WantErr: ErrCannotEditTask,
		},
		{
			Name: "Invalid: in CANCELLED status",
			InputData: updateDescriptionData{newDescription: "Another description", task: NewTaskBuilder().Status(TaskStatusCanceled).task},
			WantErr: ErrCannotEditTask,
		},
		{
			Name: "Invalid: Overdue task",
			InputData: updateDescriptionData{newDescription: "Another description", task: NewTaskBuilder().Deadline(time.Now().Add(-time.Hour)).task},
			WantErr: ErrCannotEditTask,
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
			InputData: updateDeadlineData{newDeadline: now.Add(2*time.Hour), task: NewTaskBuilder().Status(TaskStatusDone).task},
			WantErr: ErrCannotEditTask,
		},
		{
			Name: "Invalid: in CANCELLED status",
			InputData: updateDeadlineData{newDeadline: now.Add(2*time.Hour), task: NewTaskBuilder().Status(TaskStatusCanceled).task},
			WantErr: ErrCannotEditTask,
		},
		{
			Name: "Invalid: Overdue task",
			InputData: updateDeadlineData{newDeadline: now.Add(2*time.Hour), task: NewTaskBuilder().Deadline(time.Now().Add(-time.Hour)).task},
			WantErr: ErrCannotEditTask,
		},
		{
			Name: "Invalid: Deadline in past",
			InputData: updateDeadlineData{newDeadline: now.Add(-2*time.Hour), task: NewTaskBuilder().task},
			WantErr: ErrOldDeadline,
		},
		{
			Name: "Invalid: nil Deadline",
			InputData: updateDeadlineData{newDeadline: time.Time{}, task: NewTaskBuilder().task},
			WantErr: ErrEmptyDeadline,
		},
	}
}

func GetIsOverdueTestCases() []testCase[isOverdueData] {
	return []testCase[isOverdueData]{
		{
			Name: "Not Overdue",
			InputData: isOverdueData{task: NewTaskBuilder().task, expectedResponse: false },
			WantErr: nil,
		},
		{
			Name: "Overdue",
			InputData: isOverdueData{task: NewTaskBuilder().Deadline(time.Now().Add(-2*time.Hour)).task, expectedResponse: true },
			WantErr: nil,
		},
	}
}

func GetIsActiveTestCases() []testCase[isActiveData] {
	return []testCase[isActiveData]{
		{
			Name: "Active: TODO",
			InputData: isActiveData{task: NewTaskBuilder().task, expectedResponse: true },
			WantErr: nil,
		},
		{
			Name: "Active: INPROGRESS",
			InputData: isActiveData{task: NewTaskBuilder().Status(TaskStatusInProgress).task, expectedResponse: true },
			WantErr: nil,
		},
		{
			Name: "Not active: overdue TODO",
			InputData: isActiveData{task: NewTaskBuilder().Deadline(time.Now().Add(-2*time.Hour)).task, expectedResponse: false },
			WantErr: nil,
		},
		{
			Name: "Not active: overdue INPROGRESS",
			InputData: isActiveData{
				task: NewTaskBuilder().Deadline(time.Now().Add(-2*time.Hour)).Status(TaskStatusInProgress).task,
				expectedResponse: false,
			},
			WantErr: nil,
		},
		{
			Name: "Not active: DONE",
			InputData: isActiveData{task: NewTaskBuilder().Status(TaskStatusDone).task, expectedResponse: false},
			WantErr: nil,
		},
		{
			Name: "Not active: CANCELLED",
			InputData: isActiveData{task: NewTaskBuilder().Status(TaskStatusCanceled).task, expectedResponse: false},
			WantErr: nil,
		},
		{
			Name: "Not active: HOLD",
			InputData: isActiveData{task: NewTaskBuilder().Status(TaskStatusHold).task, expectedResponse: false},
			WantErr: nil,
		},
		{
			Name: "Not active: DONE and overdue",
			InputData: isActiveData{
				task: NewTaskBuilder().Status(TaskStatusDone).Deadline(time.Now().Add(-2*time.Hour)).task,
				expectedResponse: false,
			},
			WantErr: nil,
		},
		{
			Name: "Not active: CANCELLED and overdue",
			InputData: isActiveData{
				task: NewTaskBuilder().Status(TaskStatusCanceled).Deadline(time.Now().Add(-2*time.Hour)).task,
				expectedResponse: false,
			},
			WantErr: nil,
		},
		{
			Name: "Not active: HOLD and overdue",
			InputData: isActiveData{
				task: NewTaskBuilder().Status(TaskStatusHold).Deadline(time.Now().Add(-2*time.Hour)).task,
				expectedResponse: false,
			},
			WantErr: nil,
		},
	}
}

func GetIsFinishedTestCases() []testCase[isFinishedData] {
	return []testCase[isFinishedData]{
		{
			Name: "Active: TODO",
			InputData: isFinishedData{task: NewTaskBuilder().task, expectedResponse: false },
			WantErr: nil,
		},
		{
			Name: "Active: INPROGRESS",
			InputData: isFinishedData{task: NewTaskBuilder().Status(TaskStatusInProgress).task, expectedResponse: false },
			WantErr: nil,
		},
		{
			Name: "Active: HOLD",
			InputData: isFinishedData{task: NewTaskBuilder().Status(TaskStatusHold).task, expectedResponse: false },
			WantErr: nil,
		},
		{
			Name: "Finished: DONE",
			InputData: isFinishedData{task: NewTaskBuilder().Status(TaskStatusDone).task, expectedResponse: true },
			WantErr: nil,
		},
		{
			Name: "Finished: CANCELLED",
			InputData: isFinishedData{task: NewTaskBuilder().Status(TaskStatusCanceled).task, expectedResponse: true },
			WantErr: nil,
		},
	}
}

func GetCanEditTestCases() []testCase[canEditData] {
	return []testCase[canEditData]{
		{
			Name: "Can edit",
			InputData: canEditData{task: NewTaskBuilder().task, expectedResponse: true },
			WantErr: nil,
		},
		{
			Name: "Cannot edit: overdue",
			InputData: canEditData{task: NewTaskBuilder().Deadline(time.Now().Add(-2*time.Hour)).task, expectedResponse: false },
			WantErr: nil,
		},
		{
			Name: "Cannot edit: finished",
			InputData: canEditData{task: NewTaskBuilder().Status(TaskStatusDone).task, expectedResponse: false },
			WantErr: nil,
		},
		{
			Name: "Cannot edit: overdue and finished",
			InputData: canEditData{
				task: NewTaskBuilder().Status(TaskStatusDone).Deadline(time.Now().Add(-2*time.Hour)).task,
				expectedResponse: false,
			},
			WantErr: nil,
		},
	}
}

