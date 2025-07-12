package entities

import (
	"gorm.io/gorm"
)

type Book struct {
	gorm.Model
	ID          int    `gorm:"primaryKey;autoIncrement"`
	Title       string `gorm:"unique; not null"`
	Description string
	AuthorID    int     `gorm:"not null"`
	Price       float64 `gorm:"not null"`
	Year        int     `gorm:"not null"`

	Author BookAuthor `gorm:"foreignKey:AuthorID" json:"-"`
}

func NewBook(
	title string,
	description string,
	authorID int,
	price float64,
	year int,
) *Book {
	return &Book{
		Title:       title,
		Description: description,
		AuthorID:    authorID,
		Price:       price,
		Year:        year,
	}
}
