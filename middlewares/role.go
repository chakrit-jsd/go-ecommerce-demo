package middlewares

import (
	"go-ecommerce-demo/core/models/model_io"
	"go-ecommerce-demo/utils"

	"github.com/gofiber/fiber/v2"
)

func Role(c *fiber.Ctx) error {
	user := c.Locals("userInfo").(*model_io.User)
	if user.Role != "admin" {
		return utils.CusErrorFiber(c, fiber.ErrUnauthorized)
	}

	return c.Next()
}
