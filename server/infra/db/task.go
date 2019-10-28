package db

import "cloud.google.com/go/firestore"

type taskRepository struct {
	client firestore.Client
}
