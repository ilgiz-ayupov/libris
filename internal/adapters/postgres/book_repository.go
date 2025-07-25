package postgres

import (
	"github.com/ilgiz-ayupov/libris/internal/entities"
	"github.com/jmoiron/sqlx"
)

type bookRepository struct{}

func NewBookRepository() *bookRepository {
	return &bookRepository{}
}

func (r *bookRepository) CreateBook(tx *sqlx.Tx, param entities.BookCreateParam) (bookID int, err error) {
	query := `
		INSERT INTO books (title, description, publisher_id, price, year)
		VALUES (?, ?, ?, ?, ?)
		RETURNING book_id
	`

	return 0, nil
}

func (r *bookRepository) BulkSaveBookAuthors(tx *sqlx.Tx, authorIDs []int) error {
	return nil
}
