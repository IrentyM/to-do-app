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
	ctx context.Context

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
		TaskHandler: taskHandler,
	}, err
}

func (a *App) Startup(ctx context.Context) {
	a.ctx = ctx
	a.TaskHandler.SetContext(ctx)
}

//// Добавить задачу
//func (a *App) AddTask(title string) Task {
//	task := Task{
//		ID:        a.nextID,
//		Title:     title,
//		Done:      false,
//		CreatedAt: time.Now(),
//	}
//	a.tasks = append(a.tasks, task)
//	a.nextID++
//	return task
//}
//
//// Получить список задач
//func (a *App) GetTasks() []Task {
//	return a.tasks
//}
//
//// Удалить задачу
//func (a *App) DeleteTask(id int) bool {
//	for i, t := range a.tasks {
//		if t.ID == id {
//			a.tasks = append(a.tasks[:i], a.tasks[i+1:]...)
//			return true
//		}
//	}
//	return false
//}
//
//// Переключить статус задачи
//func (a *App) ToggleTask(id int) bool {
//	for i, t := range a.tasks {
//		if t.ID == id {
//			a.tasks[i].Done = !a.tasks[i].Done
//			return true
//		}
//	}
//	return false
//}
