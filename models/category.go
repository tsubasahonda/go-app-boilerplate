package models

import (
	"github.com/tsubasahonda/go-app-boilerplate/db"
)

// Category model
type Category struct {
	ID string
	Name string
}

// NewCategory make Category pointer
func NewCategory(category *db.Category) *Category {
	return &Category {
		cateogry.ID.String(),
		category.Name,
	}
}
