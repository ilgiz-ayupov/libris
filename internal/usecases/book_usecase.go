package usecases

import (
	"github.com/ilgiz-ayupov/libris/internal/entities"
	"github.com/ilgiz-ayupov/libris/pkg/gensql"
	"github.com/ilgiz-ayupov/libris/pkg/logger"
	"gorm.io/gorm"
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

func (u *BookUseCase) CreateBook(
	tx *gorm.DB,
	title string,
	description string,
	publisherID int,
	authorIDs []int,
	price float64,
	year int,
) (entities.Book, error) {
	authors, err := gensql.LoadRequiredData(func() ([]entities.BookAuthor, error) {
		return u.bookAuthorRepo.FindBookAuthorsByID(tx, authorIDs)
	}, u.log, entities.ErrBookAuthorsNotFound, "не удалось создать книгу")
	if err != nil {
		return entities.Book{}, err
	}

	publisher, err := gensql.LoadRequiredData(func() (entities.BookPublisher, error) {
		return u.bookPublisherRepo.FindBookPublisherByID(tx, publisherID)
	}, u.log, entities.ErrBookPublisherNotFound, "не удалось создать книгу")
	if err != nil {
		return entities.Book{}, err
	}

	book := entities.NewBook(
		title,
		description,
		publisher,
		authors,
		price,
		year,
	)

	if err := u.bookRepo.CreateBook(tx, &book); err != nil {
		u.log.Error("не удалось создать книгу", "error", err)
		return entities.Book{}, entities.ErrInternalError
	}

	return book, nil
}

func (u *BookUseCase) FindBooks(
	tx *gorm.DB,
	q string,
	startYear, endYear int,
) ([]entities.Book, error) {
	return gensql.LoadCanNoData(func() ([]entities.Book, error) {
		return u.bookRepo.FindBooks(tx, q, startYear, endYear)
	}, u.log, entities.ErrBooksNotFound, "не удалось найти книги")
}

func (u *BookUseCase) FindBook(tx *gorm.DB, id int) (entities.Book, error) {
	return gensql.LoadRequiredData(func() (entities.Book, error) {
		return u.bookRepo.FindBookByID(tx, id)
	}, u.log, entities.ErrBookNotFound, "не удалось найти книгу")
}
