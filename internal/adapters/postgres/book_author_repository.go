package postgres

import (
	"github.com/ilgiz-ayupov/libris/internal/entities"
	"github.com/ilgiz-ayupov/libris/pkg/gensql"
	"gorm.io/gorm"
)

type bookAuthorRepository struct{}

func NewBookAuthorRepository() *bookAuthorRepository {
	return &bookAuthorRepository{}
}

func (r *bookAuthorRepository) FindBookAuthorsByID(tx *gorm.DB, authorIDs []int) ([]entities.BookAuthor, error) {
	query := tx.Model(&entities.BookAuthor{})

	return gensql.Select[entities.BookAuthor](query, authorIDs)
}
