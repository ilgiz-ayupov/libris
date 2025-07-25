package entities

import (
	"errors"
)

var (
	ErrBookNotFound  = errors.New("книга не найдена")
	ErrBooksNotFound = errors.New("книги не найдены")
)

type Book struct {
	ID          int    `gorm:"primaryKey;autoIncrement"`
	Title       string `gorm:"unique; not null"`
	Description string
	Publisher   BookPublisher `gorm:"foreignKey:PublisherID; not null"`
	Authors     []BookAuthor  `gorm:"many2many:book_book_authors"`
	Price       float64       `gorm:"not null"`
	Year        int           `gorm:"not null"`
}

func NewBook(
	bookID int,
	title string,
	description string,
	publisher BookPublisher,
	authors []BookAuthor,
	price float64,
	year int,
) Book {
	return Book{
		ID:          bookID,
		Title:       title,
		Description: description,
		Publisher:   publisher,
		Authors:     authors,
		Price:       price,
		Year:        year,
	}
}

type BookCreateParam struct {
	Title       string  `json:"title"`
	Description string  `json:"description"`
	PublisherID int     `json:"publisher_id"`
	AuthorIDs   []int   `json:"author_ids"`
	Price       float64 `json:"price"`
	Year        int     `json:"year"`
}
