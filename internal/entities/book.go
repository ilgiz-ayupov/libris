package entities

import (
	"gorm.io/gorm"
)

type Book struct {
	gorm.Model
	ID          int    `gorm:"primaryKey;autoIncrement"`
	Title       string `gorm:"unique; not null"`
	Description string
	Author      string  `gorm:"not null"`
	Price       float64 `gorm:"not null"`
	Year        int     `gorm:"not null"`
}

func NewBook(
	title string,
	description string,
	author string,
	price float64,
	year int,
) *Book {
	return &Book{
		Title:       title,
		Description: description,
		Author:      author,
		Price:       price,
		Year:        year,
	}
}
