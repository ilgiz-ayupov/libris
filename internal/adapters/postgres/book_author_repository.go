package postgres

import (
	"github.com/ilgiz-ayupov/libris/internal/entities"
	"github.com/jmoiron/sqlx"
)

type bookAuthorRepository struct{}

func NewBookAuthorRepository() *bookAuthorRepository {
	return &bookAuthorRepository{}
}

func (r *bookAuthorRepository) FindBookAuthorsByID(tx *sqlx.Tx, authorIDs []int) ([]entities.BookAuthor, error) {
	query := `
		SELECT author_id, fio, biogpraphy, rating
		FROM book_authors
		WHERE author_id IN (:author_ids)
	`

	return nil, nil
}
