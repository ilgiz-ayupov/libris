package usecases

import (
	"database/sql"

	"github.com/ilgiz-ayupov/libris/internal/entities"
)

type BookRepository interface {
	FindBookList(tx *sql.Tx) ([]entities.Book, error)
}
