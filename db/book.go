package db

import (
	"time"

	uuid "github.com/satori/go.uuid"
)

// Book is struct of db model
type Book struct {
	ID          uuid.UUID
	Title       string
	Author      *string
	ISBN        *int64
	PublishedAt *time.Time
	Publisher   *string
	CoverURL    *string
	Overview    *string
	KeyInsights []string
	CategoryID  *uuid.UUID
	Category    *Category
	Starred     bool
}
