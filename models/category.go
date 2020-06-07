package models

import (
	"github.com/tsubasahonda/go-app-boilerplate/db"
)

// Category model
type Category struct {
	ID   string `json:"id"`
	Name string `json:"name"`
}

// NewCategory make Category pointer
func NewCategory(category *db.Category) *Category {
	return &Category{
		category.ID.String(),
		category.Name,
	}
}
