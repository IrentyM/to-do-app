package app

import (
	"context"
	"fmt"
	"to-do-app/internal/adapters/controller/service/handler"
	"to-do-app/internal/adapters/repo/postgre"
	"to-do-app/internal/usecase"
	"to-do-app/pkg/sqlite"

	"go.uber.org/zap"
)

// App struct
type App struct {
	log         *zap.Logger
	TaskHandler *handler.TaskHandler
}

func NewApp(ctx context.Context, dir string, mode string) (*App, error) {
	logger, err := NewLogger(dir, mode)
	if err != nil {
		return nil, err
	}

	db, err := sqlite.NewSQLiteConnection()
	if err != nil {
		return nil, fmt.Errorf("failed to connect to SQLite: %w", err)
	}

	taskRepo := repo.NewTaskRepo(db)
	taskUsecase := usecase.NewTaskUsecase(taskRepo)
	taskHandler := handler.NewTaskHandler(taskUsecase, ctx, logger)

	return &App{
		log:         logger,
		TaskHandler: taskHandler,
	}, err
}

func (a *App) Startup(ctx context.Context) {
	a.log.Info("Startup app")
}

func (a *App) Shutdown(ctx context.Context) {
	a.log.Info("Shutdown app")
}
