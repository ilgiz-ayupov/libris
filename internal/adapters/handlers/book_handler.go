package handlers

import (
	"database/sql"
	"log/slog"

	"github.com/gofiber/fiber/v2"
	"github.com/ilgiz-ayupov/libris/internal/entities"
	"github.com/ilgiz-ayupov/libris/internal/usecases"
)

type BookHandler struct {
	engine      *fiber.App
	db          *sql.DB
	log         *slog.Logger
	bookUseCase *usecases.BookUseCase
}

func NewBookHandler(
	engine *fiber.App,
	db *sql.DB,
	log *slog.Logger,
	bookUseCase *usecases.BookUseCase,
) *BookHandler {
	return &BookHandler{
		engine:      engine,
		db:          db,
		log:         log,
		bookUseCase: bookUseCase,
	}
}

func (h *BookHandler) RegisterRoutes() {
	g := h.engine.Group("/books")
	g.Get("/", h.findBookList)
}

func (h *BookHandler) findBookList(c *fiber.Ctx) error {
	tx, err := h.db.Begin()
	if err != nil {
		h.log.Info("не удалось открыть транзакцию", "error", err)
		return entities.ErrInternalError
	}
	defer tx.Rollback()

	bookList, err := h.bookUseCase.FindBookList(tx)
	switch err {
	case nil:
		return c.JSON(bookList)
	case entities.ErrNoData:
		return c.SendStatus(fiber.StatusNotFound)
	default:
		return c.SendStatus(fiber.StatusInternalServerError)
	}
}
