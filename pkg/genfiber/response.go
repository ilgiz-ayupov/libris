package genfiber

import (
	"github.com/gofiber/fiber/v2"
	"github.com/ilgiz-ayupov/libris/internal/entities"
)

func SendData(c *fiber.Ctx, data any) error {
	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"error": "",
		"data":  data,
	})
}

func SendCreatedData(c *fiber.Ctx, data any) error {
	return c.Status(fiber.StatusCreated).JSON(fiber.Map{
		"error": "",
		"data":  data,
	})
}

func SendError(c *fiber.Ctx, err error) error {
	var status int
	switch err {
	case entities.ErrIncorrectParams:
		status = fiber.StatusBadRequest
	case entities.ErrNoData:
		status = fiber.StatusNotFound
	default:
		status = fiber.StatusInternalServerError
	}

	return c.Status(status).JSON(fiber.Map{
		"error": err.Error(),
		"data":  nil,
	})
}
