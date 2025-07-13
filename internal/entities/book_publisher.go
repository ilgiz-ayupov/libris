package entities

import (
	"errors"

	"gorm.io/gorm"
)

var (
	ErrBookPublisherNotFound = errors.New("редактор не найден")
)

type BookPublisher struct {
	gorm.Model
	ID          int    `gorm:"primaryKey;autoIncrement"`
	Name        string `gorm:"unique; not null"`
	Description string
	Books       []Book       `gorm:"foreignKey:PublisherID"`
	Authors     []BookAuthor `gorm:"many2many:book_publisher_book_authors"`
}
