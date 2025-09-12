package handler

import (
	"context"
	"to-do-app/internal/model"
)

type TaskUsecase interface {
	CreateTask(ctx context.Context, task *model.Task) (*model.Task, error)
	GetAllTasks(ctx context.Context) ([]model.Task, error)
	ToggleTask(ctx context.Context, task *model.Task) error
	DeleteTask(ctx context.Context, task *model.Task) error
}
