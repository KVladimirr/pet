package usecase

import (
	"context"

	"github.com/google/uuid"
)


type UpdateTaskDescriptionUsecase struct {
	repo TaskRepository
}

func NewUpdateTaskDescriptionUsecase(repo TaskRepository) *UpdateTaskDescriptionUsecase {
	return &UpdateTaskDescriptionUsecase{repo: repo}
}

type UpdateTaskDescriptionDTO struct {
	TaskID uuid.UUID
	NewDescription string
}

func (u *UpdateTaskDescriptionUsecase) Execute(ctx context.Context, cmd *UpdateTaskDescriptionDTO) error {
	if cmd == nil {
		return NilDTOError
	}
	
	task, err := u.repo.GetByID(ctx, cmd.TaskID)
	if err != nil {
		return err
	}

	if err = task.UpdateDescription(cmd.NewDescription); err != nil {
		return err
	}

	if err = u.repo.Save(ctx, task); err != nil {
		return err
	}

	return nil
}