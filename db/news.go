package db

import (
	"database/sql"
	"fmt"
	"time"

	uuid "github.com/satori/go.uuid"
	"github.com/tsubasahonda/go-app-boilerplate/ptr"
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
	initialSelect := "SELECT a.id, a.title, a.author, a.published_at, a.publisher, a.cover_url, a.overview, a.category_id, c.name as category_name"
	initialFrom := "FROM article a LEFT JOIN category c ON (a.category_id = c.id)"

	query := fmt.Sprintf(`
%s
%s
ORDER BY a.published_at DESC`, initialSelect, initialFrom)

	rows, err := data.DB.Query(query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var news []*Article
	for rows.Next() {
		a, err := articleFromRows(rows)
		if err != nil {
			return nil, err
		}

		a.PublishedAt = ptr.NewTime(a.PublishedAt.UTC())
		news = append(news, a)
	}

	if err = rows.Err(); err != nil {
		return nil, err
	}

	return news, nil
}

func articleFromRows(rows *sql.Rows) (*Article, error) {
	a := &Article{}
	var categoryName *string
	err := rows.Scan(
		&a.ID,
		&a.Title,
		&a.Author,
		&a.PublishedAt,
		&a.Publisher,
		&a.CoverURL,
		&a.Overview,
		&a.CategoryID,
		&categoryName,
	)
	if err != nil {
		return nil, err
	}

	if a.CategoryID != nil {
		c := &Category{
			*a.CategoryID,
			*categoryName,
		}
		a.Category = c
	}

	return a, nil
}
