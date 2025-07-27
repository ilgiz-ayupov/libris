package main

import (
	"github.com/ilgiz-ayupov/libris/config"
	"github.com/ilgiz-ayupov/libris/internal/adapters/postgres"
	"github.com/ilgiz-ayupov/libris/internal/ui/httpserver"
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

func main() {
	app := NewApp()

	httpServer := httpserver.NewHTTPServer(app.db, app.log, app.bookUseCase)
	httpServer.Run(app.conf.Server.Address)
}
