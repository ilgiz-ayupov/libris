package usecases

import (
	"github.com/ilgiz-ayupov/libris/internal/entities"
	"github.com/ilgiz-ayupov/libris/pkg/gensql"
	"github.com/ilgiz-ayupov/libris/pkg/logger"
	"github.com/jmoiron/sqlx"
)

type BookUseCase struct {
	log               logger.Logger
	bookRepo          BookRepository
	bookAuthorRepo    BookAuthorRepository
	bookPublisherRepo BookPublisherRepository
}

func NewBookUseCase(
	log logger.Logger,
	bookRepo BookRepository,
	bookAuthorRepo BookAuthorRepository,
	bookPublisherRepo BookPublisherRepository,
) *BookUseCase {
	return &BookUseCase{
		log:               log,
		bookRepo:          bookRepo,
		bookAuthorRepo:    bookAuthorRepo,
		bookPublisherRepo: bookPublisherRepo,
	}
}

func (u *BookUseCase) CreateBook(tx *sqlx.Tx, param entities.BookCreateParam) (entities.Book, error) {
	var zero entities.Book

	authors, err := gensql.LoadRequiredData(func() ([]entities.BookAuthor, error) {
		return u.bookAuthorRepo.FindBookAuthorsByID(tx, param.AuthorIDs)
	}, u.log, entities.ErrBookAuthorsNotFound, "не удалось создать книгу")
	if err != nil {
		return zero, err
	}

	publisher, err := gensql.LoadRequiredData(func() (entities.BookPublisher, error) {
		return u.bookPublisherRepo.FindBookPublisherByID(tx, param.PublisherID)
	}, u.log, entities.ErrBookPublisherNotFound, "не удалось создать книгу")
	if err != nil {
		return zero, err
	}

	bookID, err := u.bookRepo.CreateBook(tx, param)
	if err != nil {
		u.log.Error("не удалось создать книгу", "error", err)
		return zero, entities.ErrInternalError
	}

	if err := u.bookAuthorRepo.BulkSaveBookAuthors(tx, bookID, param.AuthorIDs); err != nil {
		u.log.Error("не удалось сохранить авторов книги", "error", err)
		return zero, entities.ErrInternalError
	}

	return entities.NewBook(
		bookID,
		param.Title,
		param.Description,
		param.Price,
		param.Year,
		publisher,
		authors,
	), nil
}
