package usecases

import (
	"github.com/ilgiz-ayupov/libris/internal/entities"
	"github.com/jmoiron/sqlx"
)

type BookRepository interface {
	CreateBook(tx *sqlx.Tx, param entities.BookCreateParam) (bookID int, err error)
	BulkSaveBookAuthors(tx *sqlx.Tx, authorIDs []int) error
}

type BookAuthorRepository interface {
	FindBookAuthorsByID(tx *sqlx.Tx, authorIDs []int) ([]entities.BookAuthor, error)
}

type BookPublisherRepository interface {
	FindBookPublisherByID(tx *sqlx.Tx, id int) (entities.BookPublisher, error)
}
