package repository

import (
	"context"

	"github.com/sekky0905/gae-go-firestore-nuxt-todo/domain/model"
)

// TaskRepository is repository of task.
type TaskRepository interface {
	// TODO paging
	ListTasks(ctx context.Context) ([]model.Task, error)
	GetTask(ctx context.Context, id string) (*model.Task, error)
	CreateTask(ctx context.Context, task *model.Task) (string, error)
	UpdateTask(ctx context.Context, id string) error
	DeleteTask(ctx context.Context, id string) error
}
