package usecases

import (
	"github.com/ilgiz-ayupov/libris/internal/entities"
	"gorm.io/gorm"
)

type BookRepository interface {
	Create(tx *gorm.DB, book *entities.Book) error
	FindBooks(tx *gorm.DB, q string, startYear, endYear int, author string) ([]entities.Book, error)
}
