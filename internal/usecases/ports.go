package usecases

import (
	"github.com/ilgiz-ayupov/libris/internal/entities"
	"gorm.io/gorm"
)

type BookRepository interface {
	CreateBook(tx *gorm.DB, book *entities.Book) error
	FindBooks(tx *gorm.DB, q string, startYear, endYear int) ([]entities.Book, error)
	FindBookByID(tx *gorm.DB, id int) (entities.Book, error)
}

type BookAuthorRepository interface {
	FindBookAuthorsByID(tx *gorm.DB, authorIDs []int) ([]entities.BookAuthor, error)
}

type BookPublisherRepository interface {
	FindBookPublisherByID(tx *gorm.DB, id int) (entities.BookPublisher, error)
}
