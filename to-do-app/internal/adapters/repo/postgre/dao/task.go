package dao

import (
	"time"
	"to-do-app/internal/model"
)

type Task struct {
	ID        int64      `db:"id"`
	Title     string     `db:"title"`
	Done      bool       `db:"done"`
	CreatedAt time.Time  `db:"created_at"`
	Deadline  *time.Time `db:"deadline"`
	Priority  string     `db:"priority"`
}

func FromDomain(task *model.Task) Task {
	return Task{
		ID:        task.ID,
		Title:     task.Title,
		Done:      task.Done,
		CreatedAt: task.CreatedAt,
		Deadline:  task.Deadline,
		Priority:  task.Priority,
	}
}

func ToDomain(task *Task) model.Task {
	return model.Task{
		ID:        task.ID,
		Title:     task.Title,
		Done:      task.Done,
		CreatedAt: task.CreatedAt,
		Deadline:  task.Deadline,
		Priority:  task.Priority,
	}
}
