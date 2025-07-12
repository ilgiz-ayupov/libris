package postgres

import (
	"github.com/ilgiz-ayupov/libris/internal/entities"
	"gorm.io/gorm"
)

type bookRepository struct{}

func NewBookRepository() *bookRepository {
	return &bookRepository{}
}

func (r *bookRepository) Create(tx *gorm.DB, book *entities.Book) error {
	return tx.Create(book).Error
}

func (r *bookRepository) FindBooks(
	tx *gorm.DB,
	q string,
	startYear, endYear int,
	author string,
) ([]entities.Book, error) {
	query := tx.Model(&entities.Book{})

	if q != "" {
		query = query.Where("title ILIKE ?", "%"+q+"%")
	}
	if startYear != 0 && endYear != 0 {
		query = query.Where("year BETWEEN ? AND ?", startYear, endYear)
	}
	if author != "" {
		query = query.Where("author ILIKE ?", "%"+author+"%")
	}

	var books []entities.Book
	if err := query.Find(&books).Error; err != nil {
		return nil, err
	}
	return books, nil
}
