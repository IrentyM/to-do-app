package main

import (
	"context"
	"time"
)

// App struct
type Task struct {
	ID        int       `json:"id"`
	Title     string    `json:"title"`
	Done      bool      `json:"done"`
	CreatedAt time.Time `json:"created_at"`
}

type App struct {
	ctx    context.Context
	tasks  []Task
	nextID int
}

func NewApp() *App {
	return &App{
		tasks:  []Task{},
		nextID: 1,
	}
}

func (a *App) startup(ctx context.Context) {
	a.ctx = ctx
}

// Добавить задачу
func (a *App) AddTask(title string) Task {
	task := Task{
		ID:        a.nextID,
		Title:     title,
		Done:      false,
		CreatedAt: time.Now(),
	}
	a.tasks = append(a.tasks, task)
	a.nextID++
	return task
}

// Получить список задач
func (a *App) GetTasks() []Task {
	return a.tasks
}

// Удалить задачу
func (a *App) DeleteTask(id int) bool {
	for i, t := range a.tasks {
		if t.ID == id {
			a.tasks = append(a.tasks[:i], a.tasks[i+1:]...)
			return true
		}
	}
	return false
}

// Переключить статус задачи
func (a *App) ToggleTask(id int) bool {
	for i, t := range a.tasks {
		if t.ID == id {
			a.tasks[i].Done = !a.tasks[i].Done
			return true
		}
	}
	return false
}
