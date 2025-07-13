package entities

import (
	"errors"

	"gorm.io/gorm"
)

var (
	ErrBookNotFound  = errors.New("книга не найдена")
	ErrBooksNotFound = errors.New("книги не найдены")
)

type Book struct {
	gorm.Model
	ID          int    `gorm:"primaryKey;autoIncrement"`
	Title       string `gorm:"unique; not null"`
	Description string
	Publisher   BookPublisher `gorm:"foreignKey:PublisherID; not null"`
	Authors     []BookAuthor  `gorm:"many2many:book_book_authors"`
	Price       float64       `gorm:"not null"`
	Year        int           `gorm:"not null"`

	PublisherID int `json:"-"`
}

func NewBook(
	title string,
	description string,
	publisher BookPublisher,
	authors []BookAuthor,
	price float64,
	year int,
) Book {
	return Book{
		Title:       title,
		Description: description,
		Publisher:   publisher,
		Authors:     authors,
		Price:       price,
		Year:        year,
	}
}
