package entities

import (
	"errors"
)

var (
	ErrBookNotFound  = errors.New("книга не найдена")
	ErrBooksNotFound = errors.New("книги не найдены")
)

type Book struct {
	ID          int     `db:"book_id"`
	Title       string  `db:"title"`
	Description string  `db:"description"`
	Price       float64 `db:"price"`
	Year        int     `db:"year"`
	Publisher   BookPublisher
	Authors     []BookAuthor
}

func NewBook(
	bookID int,
	title string,
	description string,
	price float64,
	year int,
	publisher BookPublisher,
	authors []BookAuthor,
) Book {
	return Book{
		ID:          bookID,
		Title:       title,
		Description: description,
		Price:       price,
		Year:        year,
		Publisher:   publisher,
		Authors:     authors,
	}
}

type BookCreateParam struct {
	Title       string  `json:"title"`
	Description string  `json:"description"`
	Price       float64 `json:"price"`
	Year        int     `json:"year"`
	PublisherID int     `json:"publisher_id"`
	AuthorIDs   []int   `json:"author_ids"`
}
