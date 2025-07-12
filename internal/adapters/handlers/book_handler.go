package handlers

import (
	"log/slog"

	"github.com/gofiber/fiber/v2"
	"github.com/ilgiz-ayupov/libris/internal/usecases"
	"github.com/ilgiz-ayupov/libris/pkg/genfiber"
	"gorm.io/gorm"
)

type BookHandler struct {
	engine      *fiber.App
	db          *gorm.DB
	log         *slog.Logger
	bookUseCase *usecases.BookUseCase
}

func NewBookHandler(
	engine *fiber.App,
	db *gorm.DB,
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
	g.Post("/", h.createBook)
	g.Get("/", h.findBooks)
}

func (h *BookHandler) createBook(c *fiber.Ctx) error {
	type Param struct {
		Title       string
		Description string
		AuthorID    int
		Price       float64
		Year        int
	}

	var p Param
	if err := c.BodyParser(&p); err != nil {
		h.log.Debug("получены некорректные параметры", "error", err)
		return genfiber.SendError(c, err)
	}

	tx := h.db.Begin()
	if err := tx.Error; err != nil {
		h.log.Error("не удалось открыть транзакцию", "error", err)
		return genfiber.SendError(c, err)
	}
	defer tx.Rollback()

	if err := h.bookUseCase.Create(
		tx,
		p.Title,
		p.Description,
		p.AuthorID,
		p.Price,
		p.Year,
	); err != nil {
		return genfiber.SendError(c, err)
	}

	if err := tx.Commit().Error; err != nil {
		h.log.Error("не удалось зафиксировать транзакцию", "error", err)
		return genfiber.SendError(c, err)
	}

	return c.SendStatus(fiber.StatusCreated)
}

func (h *BookHandler) findBooks(c *fiber.Ctx) error {
	q := c.Query("q")
	startYear := c.QueryInt("start_year")
	endYear := c.QueryInt("end_year")
	author := c.Query("author")

	tx := h.db.Begin()
	defer tx.Rollback()

	books, err := h.bookUseCase.FindBooks(
		tx,
		q,
		startYear,
		endYear,
		author,
	)
	if err != nil {
		return genfiber.SendError(c, err)
	}
	return genfiber.SendData(c, books)
}
