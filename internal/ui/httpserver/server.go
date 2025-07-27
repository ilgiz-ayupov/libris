package httpserver

import (
	"github.com/gofiber/fiber/v2"
	"github.com/ilgiz-ayupov/libris/internal/usecases"
	"github.com/ilgiz-ayupov/libris/pkg/logger"
	"github.com/jmoiron/sqlx"
)

type HTTPServer struct {
	db  *sqlx.DB
	log logger.Logger

	bookUseCase *usecases.BookUseCase
}

func NewHTTPServer(
	db *sqlx.DB,
	log logger.Logger,
	bookUseCase *usecases.BookUseCase,
) *HTTPServer {
	return &HTTPServer{
		db:          db,
		log:         log,
		bookUseCase: bookUseCase,
	}
}

func (s *HTTPServer) Run(addr string) error {
	app := fiber.New(fiber.Config{})

	bookHandler := NewBookHandler(app, s.db, s.log, s.bookUseCase)
	bookHandler.RegisterRoutes()

	s.log.Info("Сервер запущен на порту: ", "address", addr)
	if err := app.Listen(addr); err != nil {
		s.log.Error("не удалось запустить сервер", "error", err)
	}
	return nil
}
