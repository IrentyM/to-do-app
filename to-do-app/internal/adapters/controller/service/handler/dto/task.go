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
	CreatedAt string  `json:"created_at" jsonpb:"string"`
	Deadline  *string `json:"deadline" jsonpb:"string"`
	Priority  int32   `json:"priority"`
}

func FromRequest(task Task) *model.Task {
	timeCreatedAt, err := time.Parse(time.RFC3339, task.CreatedAt)
	if err != nil {
		fmt.Println(err)
	}

	timeDeadline, err := time.Parse(time.RFC3339, *task.Deadline)
	if err != nil {
		fmt.Println(err)
	}

	return &model.Task{
		ID:        task.ID,
		Title:     task.Title,
		Done:      task.Done,
		CreatedAt: timeCreatedAt,
		Deadline:  &timeDeadline,
		Priority:  task.Priority,
	}
}

//func ToResponse(task *model.Task) *Task {
//
//	return &Task{
//		ID:        task.ID,
//		Title:     task.Title,
//		Done:      task.Done,
//		CreatedAt: task.CreatedAt,
//		Deadline:  task.Deadline,
//		Priority:  task.Priority,
//	}
//}
