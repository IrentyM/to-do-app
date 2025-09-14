package dto

import (
	"fmt"
	"time"
	"to-do-app/internal/model"
)

type Task struct {
	ID        int64   `json:"id"`
	Title     string  `json:"title"`
	Done      bool    `json:"done"`
	CreatedAt string  `json:"created_at"`
	Deadline  *string `json:"deadline"`
	Priority  string  `json:"priority"`
}

func FromRequest(task Task) *model.Task {
	var timeDeadline *time.Time
	if task.Deadline != nil && *task.Deadline != "" {
		t, err := time.Parse(time.RFC3339, *task.Deadline)
		if err == nil {
			timeDeadline = &t
		} else {
			fmt.Println(err)
		}
	}

	return &model.Task{
		ID:       task.ID,
		Title:    task.Title,
		Done:     task.Done,
		Deadline: timeDeadline,
		Priority: task.Priority,
	}
}

func ToResponse(task *model.Task) Task {
	timeCreatedAt := task.CreatedAt.Format(time.RFC3339)

	var deadline *string
	if task.Deadline != nil {
		d := task.Deadline.Format(time.RFC3339)
		deadline = &d
	}

	return Task{
		ID:        task.ID,
		Title:     task.Title,
		Done:      task.Done,
		CreatedAt: timeCreatedAt,
		Deadline:  deadline,
		Priority:  task.Priority,
	}
}
