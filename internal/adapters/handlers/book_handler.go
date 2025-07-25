package handlers

import (
	"github.com/gofiber/fiber/v2"
	"github.com/ilgiz-ayupov/libris/internal/entities"
	"github.com/ilgiz-ayupov/libris/internal/usecases"
	"github.com/ilgiz-ayupov/libris/pkg/genfiber"
	"github.com/ilgiz-ayupov/libris/pkg/logger"
	"github.com/jmoiron/sqlx"
)

type BookHandler struct {
	engine      *fiber.App
	db          *sqlx.DB
	log         logger.Logger
	bookUseCase *usecases.BookUseCase
}

func NewBookHandler(
	engine *fiber.App,
	db *sqlx.DB,
	log logger.Logger,
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
	g.Post("/", h.createBook)
}

func (h *BookHandler) createBook(c *fiber.Ctx) error {
	var p entities.BookCreateParam
	if err := c.BodyParser(&p); err != nil {
		h.log.Debug("получены некорректные параметры", "error", err)
		return genfiber.SendError(c, err)
	}

	return genfiber.ExecReturn(c, func(tx *sqlx.Tx) (entities.Book, error) {
		return h.bookUseCase.CreateBook(tx, p)
	}, h.db, h.log)
}
