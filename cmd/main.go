package main

import (
	"github.com/gofiber/fiber/v2"
	"github.com/ilgiz-ayupov/libris/config"
	"github.com/ilgiz-ayupov/libris/internal/adapters/handlers"
	"github.com/ilgiz-ayupov/libris/internal/adapters/postgres"
	"github.com/ilgiz-ayupov/libris/internal/usecases"
	"github.com/ilgiz-ayupov/libris/pkg/logger"
	"github.com/ilgiz-ayupov/libris/pkg/pgdb"
	"github.com/jmoiron/sqlx"
)

type App struct {
	db   *sqlx.DB
	log  logger.Logger
	conf *config.Config

	bookUseCase *usecases.BookUseCase
}

func NewApp() *App {
	log := logger.NewSlogLogger()

	db, err := pgdb.Connect(config.PostgresConnectURL())
	if err != nil {
		log.Error("не удалось подключиться к БД", "error", err)
		return nil
	}

	bookRepo := postgres.NewBookRepository()
	bookAuthorRepo := postgres.NewBookAuthorRepository()
	bookPublisherRepo := postgres.NewBookPublisherRepository()

	bookUseCase := usecases.NewBookUseCase(log, bookRepo, bookAuthorRepo, bookPublisherRepo)

	return &App{
		db:          db,
		log:         log,
		conf:        config.Load(),
		bookUseCase: bookUseCase,
	}
}

func (a *App) StartServer() error {
	app := fiber.New(fiber.Config{})

	bookHandler := handlers.NewBookHandler(app, a.db, a.log, a.bookUseCase)
	bookHandler.RegisterRoutes()

	a.log.Info("Сервер запущен", "address", a.conf.Server.Address)
	return app.Listen(a.conf.Server.Address)
}

func main() {
	app := NewApp()

	if err := app.StartServer(); err != nil {
		app.log.Error("не удалось запустить сервер", "error", err)
	}
}
