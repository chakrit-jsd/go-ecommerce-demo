package handlers

import (
	"go-ecommerce-demo/core/models/model_io"
	"go-ecommerce-demo/core/ports"
	"go-ecommerce-demo/utils"

	"github.com/gofiber/fiber/v2"
)

type CartHandler struct {
	cartSrv ports.CartService
}

func NewCartHandle(cartSrv ports.CartService) CartHandler {
	return CartHandler{cartSrv}
}

func (h CartHandler) GetCart(c *fiber.Ctx) error {
	userId := c.Locals("userInfo").(*model_io.User).ID
	cart, err := h.cartSrv.GetCart(&userId)
	if err != nil {
		return utils.CusErrorFiber(c, err)
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"message": "Get cart success",
		"data":    *cart,
	})
}

func (h CartHandler) PostAddProduct(c *fiber.Ctx) error {
	userId := c.Locals("userInfo").(*model_io.User).ID
	cart := model_io.CartDetail{}
	err := c.BodyParser(&cart)
	if err != nil {
		return utils.CusErrorFiber(c, fiber.ErrConflict)
	}

	cart.UserId = userId
	err = h.cartSrv.AddProduct(&cart)
	if err != nil {
		return utils.CusErrorFiber(c, err)
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"message": "Add product success",
	})
}

func (h CartHandler) PutUpdateProduct(c *fiber.Ctx) error {
	userId := c.Locals("userInfo").(*model_io.User).ID
	cart := model_io.CartDetail{}
	err := c.BodyParser(&cart)
	if err != nil {
		return utils.CusErrorFiber(c, fiber.ErrConflict)
	}

	cart.UserId = userId
	err = h.cartSrv.UpdateProduct(&cart)
	if err != nil {
		return utils.CusErrorFiber(c, err)
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"message": "Update product success",
	})
}

func (h CartHandler) DeleteProduct(c *fiber.Ctx) error {
	userId := c.Locals("userInfo").(*model_io.User).ID
	delProduct := model_io.DeleteProductsInCart{}
	err := c.BodyParser(&delProduct)
	if err != nil {
		return utils.CusErrorFiber(c, fiber.ErrConflict)
	}

	delProduct.UserId = userId
	err = h.cartSrv.DeleteProduct(&delProduct)
	if err != nil {
		return utils.CusErrorFiber(c, err)
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"message": "Delete product success",
	})
}
