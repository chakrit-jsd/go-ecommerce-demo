package handlers

import (
	"go-ecommerce-demo/core/models/model_io"
	"go-ecommerce-demo/core/ports"
	"go-ecommerce-demo/utils"

	"github.com/gofiber/fiber/v2"
)

type ProductHandler struct {
	prodSrv ports.ProductService
}

func NewProductHandle(prodSrv ports.ProductService) ProductHandler {
	return ProductHandler{prodSrv}
}

func (h ProductHandler) GetProducts(c *fiber.Ctx) error {
	queryProducts := model_io.QueryProducts{}
	err := c.ParamsParser(&queryProducts)
	if err != nil {
		return utils.CusErrorFiber(c, fiber.NewError(fiber.StatusBadRequest, "Invalid search products"))
	}

	products, err := h.prodSrv.GetProductsByQuery(&queryProducts, queryProducts.Count)
	if err != nil {
		return utils.CusErrorFiber(c, err)
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"message": "Get producst succsess",
		"data":    *products,
	})
}

func (h ProductHandler) PostCreateProduct(c *fiber.Ctx) error {
	product := model_io.Product{}
	err := c.BodyParser(&product)
	if err != nil {
		return utils.CusErrorFiber(c, fiber.NewError(fiber.StatusConflict, "Fail create product"))
	}

	err = h.prodSrv.Create(&product)
	if err != nil {
		return utils.CusErrorFiber(c, err)
	}

	return c.Status(fiber.StatusCreated).JSON(fiber.Map{
		"message": "Create product success",
	})
}

func (h ProductHandler) PutUpdateProduct(c *fiber.Ctx) error {

	product := model_io.Product{}
	err := c.BodyParser(&product)
	if err != nil {
		return utils.CusErrorFiber(c, fiber.NewError(fiber.StatusConflict, "Fail update product"))
	}

	err = h.prodSrv.Update(&product, &product.ID)
	if err != nil {
		return utils.CusErrorFiber(c, err)
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"message": "Update product success",
	})
}

func (h ProductHandler) DeleteProduct(c *fiber.Ctx) error {
	pId := model_io.Product{}
	err := c.BodyParser(&pId)
	if err != nil {
		return utils.CusErrorFiber(c, fiber.NewError(fiber.StatusConflict, "Fail delete product"))
	}

	err = h.prodSrv.Delete(&pId.ID)
	if err != nil {
		return utils.CusErrorFiber(c, err)
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"message": "Delete product success",
	})
}
