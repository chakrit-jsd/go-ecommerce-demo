package handlers

import (
	"go-ecommerce-demo/core/models/model_gorm"
	"go-ecommerce-demo/core/models/model_io"
	"go-ecommerce-demo/core/ports"
	"go-ecommerce-demo/utils"

	"github.com/gofiber/fiber/v2"
)

type OrderHandler struct {
	orderSrv   ports.OrderService
	cartSrv    ports.CartService
	productSrv ports.ProductService
}

func NewOrderHandle(orderSrv ports.OrderService, cartSrv ports.CartService, productSrv ports.ProductService) OrderHandler {
	return OrderHandler{orderSrv, cartSrv, productSrv}
}

func (h OrderHandler) PostCreateOrder(c *fiber.Ctx) error {
	orders := model_gorm.Order{}
	err := c.BodyParser(&orders)
	if err != nil {
		return utils.CusErrorFiber(c, fiber.ErrUnprocessableEntity)
	}

	userId := c.Locals("userInfo").(*model_io.User).ID
	orders.UserId = userId

	err = h.orderSrv.Create(&orders)
	if err != nil {
		return utils.CusErrorFiber(c, err)
	}

	delProductOnCart := model_io.DeleteProductsInCart{
		UserId: userId,
	}

	for _, order := range orders.OrderDetails {
		delProductOnCart.ProductsId = append(delProductOnCart.ProductsId, order.ProductId)
	}

	err = h.cartSrv.DeleteProduct(&delProductOnCart)
	if err != nil {
		return utils.CusErrorFiber(c, err)
	}

	return c.Status(fiber.StatusCreated).JSON(fiber.Map{
		"message": "Create order success",
	})
}

func (h OrderHandler) GetPendingOrders(c *fiber.Ctx) error {
	userId := c.Locals("userInfo").(*model_io.User).ID
	orders, err := h.orderSrv.GetPendingOrders(&userId)
	if err != nil {
		return utils.CusErrorFiber(c, err)
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"message": "Get pending orders success",
		"data":    *orders,
	})
}

func (h OrderHandler) PutUpdateStatus(c *fiber.Ctx) error {
	order := model_io.Order{}
	err := c.BodyParser(&order)
	if err != nil {
		return utils.CusErrorFiber(c, fiber.ErrUnprocessableEntity)
	}

	err = h.orderSrv.UpdateStatus(&order.ID, &order.Status)
	if err != nil {
		return utils.CusErrorFiber(c, err)
	}

	if order.Status == "confirmed" {
		orderRes, err := h.orderSrv.GetOrderById(&order.ID)
		if err != nil {
			return utils.CusErrorFiber(c, err)
		}
		for _, o := range orderRes.OrderDetails {
			o.Quantity = -o.Quantity
			err = h.productSrv.UpdateStock(&o.Quantity, &o.ProductId)
			if err != nil {
				return utils.CusErrorFiber(c, err)
			}
		}
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"message": "Update order success",
	})
}

func (h OrderHandler) GetHistory(c *fiber.Ctx) error {
	userId := c.Locals("userInfo").(*model_io.User).ID
	orders, err := h.orderSrv.GetOrders(&userId)
	if err != nil {
		return utils.CusErrorFiber(c, err)
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"message": "Get orders history success",
		"data":    *orders,
	})
}
