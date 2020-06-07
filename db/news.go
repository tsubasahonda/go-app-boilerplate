package db

import (
	"time"

	uuid "github.com/satori/go.uuid"
)

// Article is struct of db model
type Article struct {
	ID          uuid.UUID
	Title       string
	Author      *string
	PublishedAt *time.Time
	Publisher   *string
	CoverURL    *string
	Overview    *string
	CategoryID  *uuid.UUID
	Category    *Category
}

// GetNews from SQLDB
func (data SQLDataStorage) GetNews() ([]*Article, error) {
	var news []*Article
	return news, nil
}
