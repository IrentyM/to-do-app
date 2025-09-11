package model

import "time"

type Task struct {
	ID        int64
	Title     string
	Done      bool
	CreatedAt time.Time
	Deadline  *time.Time
	Priority  int32
}
