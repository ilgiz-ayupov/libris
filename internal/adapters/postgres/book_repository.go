package postgres

import (
	"github.com/ilgiz-ayupov/libris/internal/entities"
	"github.com/ilgiz-ayupov/libris/pkg/gensql"
	"github.com/jmoiron/sqlx"
)

type bookRepository struct{}

func NewBookRepository() *bookRepository {
	return &bookRepository{}
}

func (r *bookRepository) CreateBook(tx *sqlx.Tx, param entities.BookCreateParam) (bookID int, err error) {
	query := `
		INSERT INTO books (title, description, price, year, publisher_id)
		VALUES (:title, :description, :price, :year, :publisher_id)
		RETURNING book_id
	`

	return gensql.ExecReturnID[int](tx, query, map[string]any{
		"title":        param.Title,
		"description":  param.Description,
		"price":        param.Price,
		"year":         param.Year,
		"publisher_id": param.PublisherID,
	})
}
