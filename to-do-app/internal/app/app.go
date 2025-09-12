package app

import (
	"context"
	"fmt"
	"to-do-app/internal/adapters/controller/service/handler"
	"to-do-app/internal/adapters/repo/postgre"
	"to-do-app/internal/usecase"
	postgres "to-do-app/pkg"
)

// App struct
type App struct {
	ctx         context.Context
	cfg         Config
	TaskHandler *handler.TaskHandler
}

func NewApp(cfg *Config) (*App, error) {
	db, err := postgres.NewPostgreConnection(cfg.Postgres)
	if err != nil {
		return nil, fmt.Errorf("failed to connect to PostgreSQL: %w", err)
	}

	taskRepo := repo.NewTaskRepo(db)
	taskUsecase := usecase.NewTaskUsecase(taskRepo)
	taskHandler := handler.NewTaskHandler(taskUsecase, nil)

	return &App{
		ctx:         nil,
		TaskHandler: taskHandler,
	}, err
}

func (a *App) Startup(ctx context.Context) {
	ctx = context.Background()
	a.ctx = ctx
	a.TaskHandler.SetContext(ctx)
}
