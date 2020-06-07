package db

import uuid "github.com/satori/go.uuid"

// Category is struct of db model
type Category struct {
	ID   uuid.UUID
	Name string
}
