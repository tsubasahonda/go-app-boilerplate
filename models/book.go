package models

import (
	"github.com/tsubasahonda/go-app-boilerplate/db"
)

// Book model
type Book struct {
	Title       string    `json:"title"`
	Author      *string   `json:"author"`
	ISBN        *int64    `json:"isbn"`
	ID          string    `json:"id"`
	PublishedAt *string   `json:"publishedAt"`
	Publisher   *string   `json:"publisher"`
	CoverURL    *string   `json:"cover_url"`
	Overview    *string   `json:"overview"`
	KeyInsights []string  `json:"key_insights"`
	CategoryID  *string   `json:"category_id"`
	Category    *Category `json:"cateogry,omitempty"`
	Starred     bool      `json:"starred"`
}

// NewBook is initialize pointer of Book
func NewBook(book *db.Book) *Book {
	var publishedAt string

	if book.PublishedAt != nil {
		publishedAt = book.PublishedAt.Format(DateTimeFormat)
	}

	b := &Book{
		book.Title,
		book.Author,
		book.ISBN,
		book.ID.String(),
		&publishedAt,
		book.Publisher,
		book.CoverURL,
		book.Overview,
		book.KeyInsights,
		nil,
		nil,
		book.Starred,
	}

	if book.Category != nil {
		categoryID := book.CategoryID.String()
		b.CategoryID = &categoryID
		b.Category = NewCategory(book.Category)
	}

	return b
}
