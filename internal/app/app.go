package app

import (
	"context"
	"fmt"
	"to-do-app/internal/adapters/controller/service/handler"
	"to-do-app/internal/adapters/repo/sqlite"
	"to-do-app/internal/usecase"
	"to-do-app/pkg/sqlite"

	"go.uber.org/zap"
)

// App represents the main application structure.
type App struct {
	log         *zap.Logger
	TaskHandler *handler.TaskHandler
}

// NewApp initializes a new App instance.
func NewApp(ctx context.Context, dir string, mode string) (*App, error) {
	// Initialize logger
	logger, err := NewLogger(dir, mode)
	if err != nil {
		return nil, err
	}

	// Connect to SQLite
	db, err := sqlite.NewSQLiteConnection()
	if err != nil {
		return nil, fmt.Errorf("failed to connect to SQLite: %w", err)
	}

	// Wire up repository -> usecase -> handler
	taskRepo := repo.NewTaskRepo(db)
	taskUsecase := usecase.NewTaskUsecase(taskRepo)
	taskHandler := handler.NewTaskHandler(taskUsecase, ctx, logger)

	return &App{
		log:         logger,
		TaskHandler: taskHandler,
	}, err
}

// Startup is called when the application starts.
func (a *App) Startup(ctx context.Context) {
	a.log.Info("Startup app")
}

// Shutdown is called when the application is shutting down.
func (a *App) Shutdown(ctx context.Context) {
	a.log.Info("Shutdown app")
}
