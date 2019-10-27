package model

import "time"

// Task is model of task.
type Task struct {
	ID        string
	Name      string
	Content   string
	CreatedAt time.Time
	UpdatedAt time.Time
}
