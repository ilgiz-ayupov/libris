package genfiber

import (
	"log/slog"

	"github.com/gofiber/fiber/v2"
	"github.com/ilgiz-ayupov/libris/internal/entities"
	"gorm.io/gorm"
)

func Exec(c *fiber.Ctx, fn func(tx *gorm.DB) error, db *gorm.DB, log *slog.Logger) error {
	tx := db.Begin()
	if err := tx.Error; err != nil {
		log.Error("не удалось открыть транзакцию", "error", err)
		return SendError(c, entities.ErrInternalError)
	}
	defer tx.Rollback()

	err := fn(tx)
	if err != nil {
		return SendError(c, err)
	}

	if err := tx.Commit().Error; err != nil {
		log.Error("не удалось зафиксировать транзакцию", "error", err)
		return SendError(c, entities.ErrInternalError)
	}

	return c.SendStatus(fiber.StatusOK)
}

func ExecReturn[T any](c *fiber.Ctx, fn func(tx *gorm.DB) (T, error), db *gorm.DB, log *slog.Logger) error {
	tx := db.Begin()
	if err := tx.Error; err != nil {
		log.Error("не удалось открыть транзакцию", "error", err)
		return SendError(c, entities.ErrInternalError)
	}
	defer tx.Rollback()

	data, err := fn(tx)
	if err != nil {
		return SendError(c, err)
	}

	if err := tx.Commit().Error; err != nil {
		log.Error("не удалось зафиксировать транзакцию", "error", err)
		return SendError(c, entities.ErrInternalError)
	}

	return SendCreatedData(c, data)
}

func LoadData[T any](c *fiber.Ctx, fn func(tx *gorm.DB) (T, error), db *gorm.DB, log *slog.Logger) error {
	tx := db.Begin()
	if err := tx.Error; err != nil {
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
