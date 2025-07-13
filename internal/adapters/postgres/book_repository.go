package postgres

import (
	"github.com/ilgiz-ayupov/libris/internal/entities"
	"github.com/ilgiz-ayupov/libris/pkg/gensql"
	"gorm.io/gorm"
)

type bookRepository struct{}

func NewBookRepository() *bookRepository {
	return &bookRepository{}
}

func (r *bookRepository) CreateBook(tx *gorm.DB, book *entities.Book) error {
	return tx.Create(book).Error
}

func (r *bookRepository) FindBooks(
	tx *gorm.DB,
	q string,
	startYear, endYear int,
) ([]entities.Book, error) {
	query := tx.Model(&entities.Book{}).
		Preload("Publisher").
		Preload("Authors")

	if q != "" {
		query = query.Where("title ILIKE ?", "%"+q+"%")
	}
	if startYear != 0 && endYear != 0 {
		query = query.Where("year BETWEEN ? AND ?", startYear, endYear)
	}

	return gensql.Select[entities.Book](query)
}

func (r *bookRepository) FindBookByID(tx *gorm.DB, id int) (entities.Book, error) {
	query := tx.Model(&entities.Book{}).
		Preload("Publisher").
		Preload("Authors")

	return gensql.Get[entities.Book](query, id)
}
