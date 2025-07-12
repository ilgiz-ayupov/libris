package usecases

import (
	"database/sql"
	"log/slog"

	"github.com/ilgiz-ayupov/libris/internal/entities"
	"gorm.io/gorm"
)

type BookUseCase struct {
	log      *slog.Logger
	bookRepo BookRepository
}

func NewBookUseCase(
	log *slog.Logger,
	bookRepo BookRepository,
) *BookUseCase {
	return &BookUseCase{
		log:      log,
		bookRepo: bookRepo,
	}
}

func (u *BookUseCase) CreateBook(
	tx *gorm.DB,
	title string,
	description string,
	authorID int,
	price float64,
	year int,
) (*entities.Book, error) {
	book := entities.NewBook(
		title,
		description,
		authorID,
		price,
		year,
	)

	if err := u.bookRepo.Create(tx, book); err != nil {
		u.log.Error("не удалось создать книгу", "error", err)
		return nil, entities.ErrInternalError
	}
	return book, nil
}

func (u *BookUseCase) FindBooks(
	tx *gorm.DB,
	q string,
	startYear, endYear int,
	author string,
) ([]entities.Book, error) {
	books, err := u.bookRepo.FindBooks(tx, q, startYear, endYear, author)
	switch err {
	case nil:
		return books, nil
	case sql.ErrNoRows:
		return nil, entities.ErrNoData
	default:
		u.log.Error("не удалось получить список книг", "error", err)
		return nil, entities.ErrInternalError
	}
}
