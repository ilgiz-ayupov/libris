package postgres

import (
	"fmt"
	"strings"

	"github.com/ilgiz-ayupov/libris/internal/entities"
	"github.com/ilgiz-ayupov/libris/pkg/gensql"
	"github.com/jmoiron/sqlx"
)

type bookAuthorRepository struct{}

func NewBookAuthorRepository() *bookAuthorRepository {
	return &bookAuthorRepository{}
}

func (r *bookAuthorRepository) FindBookAuthorsByID(tx *sqlx.Tx, authorIDs []int) ([]entities.BookAuthor, error) {
	query := `
		SELECT author_id, fio, biography, rating
		FROM book_authors
		WHERE author_id IN (:author_ids)
	`

	return gensql.SelectRebind[entities.BookAuthor](tx, query, map[string]any{
		"author_ids": authorIDs,
	})
}

func (r *bookAuthorRepository) BulkSaveBookAuthors(tx *sqlx.Tx, bookID int, authorIDs []int) error {
	query := `
		INSERT INTO book_book_authors (book_id, author_id)
		VALUES %s
	`

	values := make([]string, 0, len(authorIDs))
	for _, authorID := range authorIDs {
		v := fmt.Sprintf("(%d, %d)", bookID, authorID)
		values = append(values, v)
	}

	query = fmt.Sprintf(query, strings.Join(values, ","))

	_, err := tx.Exec(query)
	return err
}
