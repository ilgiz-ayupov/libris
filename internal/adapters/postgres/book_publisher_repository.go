package postgres

import (
	"github.com/ilgiz-ayupov/libris/internal/entities"
	"github.com/ilgiz-ayupov/libris/pkg/gensql"
	"github.com/jmoiron/sqlx"
)

type bookPublisherRepository struct{}

func NewBookPublisherRepository() *bookPublisherRepository {
	return &bookPublisherRepository{}
}

func (r *bookPublisherRepository) FindBookPublisherByID(tx *sqlx.Tx, publisherID int) (entities.BookPublisher, error) {
	query := `
		SELECT publisher_id, name, description
		FROM book_publishers
		WHERE publisher_id = :publisher_id
	`

	return gensql.Get[entities.BookPublisher](tx, query, map[string]any{
		"publisher_id": publisherID,
	})
}
