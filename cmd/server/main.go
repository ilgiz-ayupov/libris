package main

import (
	"database/sql"
	"log"
	"log/slog"

	"github.com/gofiber/fiber/v2"
	"github.com/ilgiz-ayupov/libris/config"
	"github.com/ilgiz-ayupov/libris/internal/adapters/handlers"
	"github.com/ilgiz-ayupov/libris/internal/adapters/postgres"
	"github.com/ilgiz-ayupov/libris/internal/usecases"
	"github.com/ilgiz-ayupov/libris/pkg/logger"
	"github.com/ilgiz-ayupov/libris/pkg/pgdb"
)

type App struct {
	db   *sql.DB
	log  *slog.Logger
	conf *config.Config

	bookUseCase *usecases.BookUseCase
}

func NewApp() *App {
	db, err := pgdb.Connect(config.PostgresConnectURL())
	if err != nil {
		log.Fatalln(err)
	}

	log := logger.Init()

	bookRepo := postgres.NewBookRepository()

	bookUseCase := usecases.NewBookUseCase(log, bookRepo)

	return &App{
		db:          db,
		log:         log,
		conf:        config.Load(),
		bookUseCase: bookUseCase,
	}
}

func (a *App) StartServer() error {
	app := fiber.New()

	bookHandler := handlers.NewBookHandler(app, a.db, a.log, a.bookUseCase)
	bookHandler.RegisterRoutes()

	log.Println("Сервер запущен на", a.conf.Server.Address)
	return app.Listen(a.conf.Server.Address)
}

func main() {
	app := NewApp()

	if err := app.StartServer(); err != nil {
		log.Fatalln(err)
	}
}
