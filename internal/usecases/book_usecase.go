package usecases

import (
	"database/sql"
	"log/slog"

	"github.com/ilgiz-ayupov/libris/internal/entities"
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

func (u *BookUseCase) FindBookList(tx *sql.Tx) ([]entities.Book, error) {
	bookList, err := u.bookRepo.FindBookList(tx)
	switch err {
	case nil:
		return bookList, nil
	case sql.ErrNoRows:
		return nil, entities.ErrNoData
	default:
		u.log.Error("не удалось получить список книг", "error", err)
		return nil, entities.ErrInternalError
	}
}
