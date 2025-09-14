package usecase

import (
	"context"
	"time"
	"to-do-app/internal/model"
)

// taskUsecase provides business logic for working with tasks.
// It depends on a TaskRepo interface for persistence operations.
type taskUsecase struct {
	taskRepo TaskRepo
}

// NewTaskUsecase is a constructor for taskUsecase
func NewTaskUsecase(taskRepo TaskRepo) *taskUsecase {
	return &taskUsecase{
		taskRepo: taskRepo,
	}
}

// CreateTask adds a new task with the current timestamp
func (u *taskUsecase) CreateTask(ctx context.Context, task *model.Task) (*model.Task, error) {
	task.CreatedAt = time.Now()
	return u.taskRepo.CreateTask(ctx, task)
}

// GetAllTasks retrieves all tasks from the repository
func (u *taskUsecase) GetAllTasks(ctx context.Context) ([]model.Task, error) {
	return u.taskRepo.GetAllTask(ctx)
}

// ToggleTask switches the "done" status of a task
func (u *taskUsecase) ToggleTask(ctx context.Context, task *model.Task) error {
	return u.taskRepo.ToggleTask(ctx, task)
}

// DeleteTask removes a task from the repository
func (u *taskUsecase) DeleteTask(ctx context.Context, task *model.Task) error {
	return u.taskRepo.DeleteTask(ctx, task)
}

// UpdateTask updates the properties of an existing task
func (u *taskUsecase) UpdateTask(ctx context.Context, task *model.Task) error {
	return u.taskRepo.UpdateTask(ctx, task)
}
