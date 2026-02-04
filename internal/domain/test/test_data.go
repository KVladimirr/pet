package test

import(
	"time"
	"tasker/internal/domain"
)

type newTaskData struct{
	Title 		string
	Description string
	Deadline 	time.Time
}

type updateStatusData struct {
	newStatus	domain.TaskStatus
}

type testCase struct{
	Name string
	InputData newTaskData
	WantErr error
}

func GetNewTaskTestData() []testCase {
	var deadline = time.Now().Add(time.Hour)
	
	return []testCase{
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

// func GetUpdateStatusTestCases() []testCase {
// 	return []testCase{
// 		{
// 			Name: "valid",
// 			InputData: updateStatusData{newStatus: },
// 			WantErr: nil,
// 		},
// 	}
// }
