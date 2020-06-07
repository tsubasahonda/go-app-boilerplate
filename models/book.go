package models

import (
	"github.com/tsubasahonda/go-app-boilerplate/db"
)

// Book model
type Book struct {
	Title       string
	Author      *string
	ISBN        *int64
	ID          string
	PublishedAt *string
	Publisher   *string
	CoverURL    *string
	Overview    *string
	keyInsights []string
	CategoryID  *string
	Category    *Category
	Starred     bool
}

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
		book.Overview,
		book.keyInsights,
		nil,
		nil,
		book.Starred,
	}

	if book.Category != nil {
		categoryID := book.CategoryID.String()
		b.CategoryID = &categoryID
		b.Category = NewCategory(book.Category)
	}

	retrun b
}
