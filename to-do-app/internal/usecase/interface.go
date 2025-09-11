package usecase

import (
	"context"
	"to-do-app/internal/model"
)

type TaskRepo interface {
	CreateTask(ctx context.Context, task *model.Task) (*model.Task, error)
	GetAllTask(ctx context.Context) ([]*model.Task, error)
	DeleteTask(ctx context.Context, task *model.Task) error
}
