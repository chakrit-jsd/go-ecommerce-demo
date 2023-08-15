package handlers

import (
	"go-ecommerce-demo/core/models/model_io"
	"go-ecommerce-demo/core/ports"
	"go-ecommerce-demo/utils"

	"github.com/gofiber/fiber/v2"
)

type CategoryHandler struct {
	catSrv ports.CategoryService
}

func NewCategoryHandle(catSrv ports.CategoryService) CategoryHandler {
	return CategoryHandler{catSrv}
}

func (h CategoryHandler) GetCategories(c *fiber.Ctx) error {
	categories, err := h.catSrv.GetCategories()
	if err != nil {
		return utils.CusErrorFiber(c, err)
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"message": "Get categories",
		"data":    *categories,
	})
}

func (h CategoryHandler) PostCreateCategory(c *fiber.Ctx) error {
	category := model_io.Category{}
	c.BodyParser(&category)
	err := h.catSrv.Create(&category)
	if err != nil {
		return utils.CusErrorFiber(c, err)
	}

	return c.Status(fiber.StatusCreated).JSON(fiber.Map{
		"message": "Create category success",
	})
}

func (h CategoryHandler) PutUpdateCategory(c *fiber.Ctx) error {
	category := model_io.Category{}
	c.BodyParser(&category)
	err := h.catSrv.Update(&category, &category.ID)
	if err != nil {
		return utils.CusErrorFiber(c, err)
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"message": "Update category success",
	})
}

func (h CategoryHandler) DeleteCategory(c *fiber.Ctx) error {
	category := model_io.Category{}
	c.BodyParser(&category)
	err := h.catSrv.Delete(&category.ID)
	if err != nil {
		return utils.CusErrorFiber(c, err)
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"message": "Delete category success",
	})
}
