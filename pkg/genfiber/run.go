package genfiber

import (
	"github.com/gofiber/fiber/v2"
	"github.com/ilgiz-ayupov/libris/internal/entities"
	"github.com/ilgiz-ayupov/libris/pkg/logger"
	"github.com/jmoiron/sqlx"
)

func Exec(c *fiber.Ctx, fn func(tx *sqlx.Tx) error, db *sqlx.DB, log logger.Logger) error {
	tx, err := db.Beginx()
	if err != nil {
		log.Error("не удалось открыть транзакцию", "error", err)
		return SendError(c, entities.ErrInternalError)
	}
	defer tx.Rollback()

	if err := fn(tx); err != nil {
		return SendError(c, err)
	}

	if err := tx.Commit(); err != nil {
		log.Error("не удалось зафиксировать транзакцию", "error", err)
		return SendError(c, entities.ErrInternalError)
	}

	return c.SendStatus(fiber.StatusOK)
}

func ExecReturn[T any](c *fiber.Ctx, fn func(tx *sqlx.Tx) (T, error), db *sqlx.DB, log logger.Logger) error {
	tx, err := db.Beginx()
	if err != nil {
		log.Error("не удалось открыть транзакцию", "error", err)
		return SendError(c, entities.ErrInternalError)
	}
	defer tx.Rollback()

	data, err := fn(tx)
	if err != nil {
		return SendError(c, err)
	}

	if err := tx.Commit(); err != nil {
		log.Error("не удалось зафиксировать транзакцию", "error", err)
		return SendError(c, entities.ErrInternalError)
	}

	return SendCreatedData(c, data)
}

func LoadData[T any](c *fiber.Ctx, fn func(tx *sqlx.Tx) (T, error), db *sqlx.DB, log logger.Logger) error {
	tx, err := db.Beginx()
	if err != nil {
		log.Error("не удалось открыть транзакцию", "error", err)
		return SendError(c, entities.ErrInternalError)
	}
	defer tx.Rollback()

	data, err := fn(tx)
	if err != nil {
		return SendError(c, err)
	}
	return SendData(c, data)
}
