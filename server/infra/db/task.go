package db

import (
	"context"
	"os"

	"golang.org/x/xerrors"

	"cloud.google.com/go/firestore"
	"github.com/sekky0905/gae-go-firestore-nuxt-todo/domain/repository"
)

type taskRepository struct {
	client *firestore.Client
}

func NewTaskRepository(ctx context.Context) (repository.TaskRepository, error) {
	projectID := os.Getenv("PROJECT_ID")
	if projectID == "" {
		return nil, xerrors.New("error in main method")
	}
	client, err := firestore.NewClient(ctx, projectID)
	if err != nil {
		return nil, xerrors.Errorf("new firestore client: %w", err)
	}

	return &taskRepository{client: client}, nil
}
