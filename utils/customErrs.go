package utils

import (
	"errors"
	"strconv"

	"github.com/gofiber/fiber/v2"
)

func CusErrorDB(err error) error {
	msg := err.Error()
	code := msg[6:10]

	if _, err := strconv.Atoi(code); err == nil {
		msg = code
	}

	switch msg {
	case "1062":
		return fiber.NewError(fiber.StatusConflict, "Name already exists")
	case "1452":
		return fiber.NewError(fiber.StatusConflict, "References key not found")
	case "1690":
		return fiber.NewError(fiber.StatusBadRequest, "Insufficient stock")
	case "record not found":
		return fiber.NewError(fiber.StatusBadRequest, "Not Found")
	case "update fail id not found":
		return fiber.NewError(fiber.StatusConflict, "Update fail ID not found")
	case "delete fail id not found":
		return fiber.NewError(fiber.StatusConflict, "Delete fail ID not found")
	default:
		return fiber.NewError(fiber.StatusInternalServerError)
	}
}

func CusErrorFiber(c *fiber.Ctx, err error) error {
	code := fiber.StatusInternalServerError
	var e *fiber.Error
	if errors.As(err, &e) {
		code = e.Code
	}
	// c.Set(HeaderContentType, MIMETextPlainCharsetUTF8)
	return c.Status(code).JSON(fiber.Map{
		"message": err.Error(),
	})
}
