package handlers

import (
	"log/slog"

	"github.com/gofiber/fiber/v2"
	"github.com/ilgiz-ayupov/libris/internal/entities"
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

	return genfiber.ExecReturn(c, func(tx *gorm.DB) (*entities.Book, error) {
		return h.bookUseCase.CreateBook(
			tx,
			p.Title,
			p.Description,
			p.AuthorID,
			p.Price,
			p.Year,
		)
	}, h.db, h.log)

}

func (h *BookHandler) findBooks(c *fiber.Ctx) error {
	q := c.Query("q")
	startYear := c.QueryInt("start_year")
	endYear := c.QueryInt("end_year")
	author := c.Query("author")

	return genfiber.LoadData(c, func(tx *gorm.DB) ([]entities.Book, error) {
		return h.bookUseCase.FindBooks(
			tx,
			q,
			startYear,
			endYear,
			author,
		)
	}, h.db, h.log)
}
