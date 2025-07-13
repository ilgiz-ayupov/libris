package main

import (
	"log"
	"log/slog"

	"github.com/gofiber/fiber/v2"
	"github.com/ilgiz-ayupov/libris/config"
	"github.com/ilgiz-ayupov/libris/internal/adapters/handlers"
	"github.com/ilgiz-ayupov/libris/internal/adapters/postgres"
	"github.com/ilgiz-ayupov/libris/internal/entities"
	"github.com/ilgiz-ayupov/libris/internal/usecases"
	"github.com/ilgiz-ayupov/libris/pkg/logger"
	"github.com/ilgiz-ayupov/libris/pkg/pgdb"
	"gorm.io/gorm"
)

type App struct {
	db   *gorm.DB
	log  *slog.Logger
	conf *config.Config

	bookUseCase *usecases.BookUseCase
}

func NewApp() *App {
	log := logger.Init()

	db, err := pgdb.Connect(config.PostgresConnectURL())
	if err != nil {
		log.Error("не удалось подключиться к БД", "error", err)
		return nil
	}

	db.AutoMigrate(
		&entities.BookAuthor{},
		&entities.Book{},
		&entities.BookPublisher{},
	)

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
