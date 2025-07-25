package postgres

import (
	"github.com/ilgiz-ayupov/libris/internal/entities"
	"gorm.io/gorm"
)

type bookPublisherRepository struct{}

func NewBookPublisherRepository() *bookPublisherRepository {
	return &bookPublisherRepository{}
}

func (r *bookPublisherRepository) FindBookPublisherByID(tx *gorm.DB, id int) (entities.BookPublisher, error) {
	query := `
		SELECT publisher_id, name, description
		FROM book_publishers
		WHERE publisher_id = ?
	`

	return entities.BookPublisher{}, nil
}
