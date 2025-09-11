package usecase

import (
	"context"
	"to-do-app/internal/model"
)

type taskUsecase struct {
	taskRepo TaskRepo
}

func NewTaskUsecase(taskRepo TaskRepo) *taskUsecase {
	return &taskUsecase{
		taskRepo: taskRepo,
	}
}

func (u *taskUsecase) CreateTask(ctx context.Context, task *model.Task) (*model.Task, error) {
	return u.taskRepo.CreateTask(ctx, task)
}

func (u *taskUsecase) GetAllTasks(ctx context.Context) ([]*model.Task, error) {
	return u.taskRepo.GetAllTask(ctx)
}

func (u *taskUsecase) DeleteTask(ctx context.Context, task *model.Task) error {
	return u.taskRepo.DeleteTask(ctx, task)
}
