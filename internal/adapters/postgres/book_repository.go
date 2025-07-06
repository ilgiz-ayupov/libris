package postgres

import (
	"database/sql"

	"github.com/ilgiz-ayupov/libris/internal/entities"
)

type bookRepository struct{}

func NewBookRepository() *bookRepository {
	return &bookRepository{}
}

func (r *bookRepository) FindBookList(tx *sql.Tx) ([]entities.Book, error) {
	query := `
		SELECT b.book_id, b.title, b.description, b.author, b.price, b.year, b.created_date, b.deleted_date
		FROM books b
	`

	rows, err := tx.Query(query)
	if err != nil {
		return nil, err
	}

	var bookList []entities.Book
	for rows.Next() {
		var b entities.Book
		if err := rows.Scan(&b.ID, &b.Title, &b.Description, &b.Author, &b.Price, &b.Year, &b.CreatedDate, &b.DeletedDate); err != nil {
			return nil, err
		}

		bookList = append(bookList, b)
	}

	return bookList, nil
}
