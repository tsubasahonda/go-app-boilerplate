package models

import (
	"github.com/tsubasahonda/go-app-boilerplate/db"
)

// Article model
type Article struct {
	Title       string    `json:"title"`
	Author      *string   `json:"author"`
	ID          string    `json:"id"`
	PublishedAt *string   `json:"publishedAt"`
	Publisher   *string   `json:"publisher"`
	CoverURL    *string   `json:"cover_url"`
	Overview    *string   `json:"overview"`
	CategoryID  *string   `json:"category_id"`
	Category    *Category `json:"cateogry,omitempty"`
}

// Newarticle is initialize pointer of article
func Newarticle(article *db.Article) *Article {
	var publishedAt string

	if article.PublishedAt != nil {
		publishedAt = article.PublishedAt.Format(DateTimeFormat)
	}

	b := &Article{
		article.Title,
		article.Author,
		article.ID.String(),
		&publishedAt,
		article.Publisher,
		article.CoverURL,
		article.Overview,
		nil,
		nil,
	}

	if article.Category != nil {
		categoryID := article.CategoryID.String()
		b.CategoryID = &categoryID
		b.Category = NewCategory(article.Category)
	}

	return b
}
