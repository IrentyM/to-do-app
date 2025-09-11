package dto

import (
	"time"
	"to-do-app/internal/model"
)

type Task struct {
	ID        int64      `json:"id"`
	Title     string     `json:"title"`
	Done      bool       `json:"done"`
	CreatedAt time.Time  `json:"created_at"`
	Deadline  *time.Time `json:"deadline"`
	Priority  int32      `json:"priority"`
}

func FromRequest(task Task) *model.Task {
	return &model.Task{
		ID:        task.ID,
		Title:     task.Title,
		Done:      task.Done,
		CreatedAt: task.CreatedAt,
		Deadline:  task.Deadline,
		Priority:  task.Priority,
	}
}

func ToResponse(task *model.Task) *Task {
	return &Task{
		ID:        task.ID,
		Title:     task.Title,
		Done:      task.Done,
		CreatedAt: task.CreatedAt,
		Deadline:  task.Deadline,
		Priority:  task.Priority,
	}
}
