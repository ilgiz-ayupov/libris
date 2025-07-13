package postgres

import (
	"github.com/ilgiz-ayupov/libris/internal/entities"
	"github.com/ilgiz-ayupov/libris/pkg/gensql"
	"gorm.io/gorm"
)

type bookPublisherRepository struct{}

func NewBookPublisherRepository() *bookPublisherRepository {
	return &bookPublisherRepository{}
}

func (r *bookPublisherRepository) FindBookPublisherByID(tx *gorm.DB, id int) (entities.BookPublisher, error) {
	query := tx.Model(&entities.BookPublisher{})

	return gensql.Get[entities.BookPublisher](query, id)
}
