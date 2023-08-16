package handlers

import (
	"go-ecommerce-demo/core/models/model_io"
	"go-ecommerce-demo/core/ports"
	"go-ecommerce-demo/utils"

	"github.com/gofiber/fiber/v2"
)

type UserHandler struct {
	userHandle ports.UserService
}

func NewUserHandler(userHandle ports.UserService) UserHandler {
	return UserHandler{userHandle}
}

func (h UserHandler) PostRegister(c *fiber.Ctx) error {

	user := model_io.User{}

	err := c.BodyParser(&user)
	if err != nil {
		return utils.CusErrorFiber(c, fiber.ErrUnprocessableEntity)
	}

	err = h.userHandle.Create(&user)
	if err != nil {
		return utils.CusErrorFiber(c, err)
	}

	return c.Status(fiber.StatusCreated).JSON(
		fiber.Map{
			"message": "Register Success",
		})
}

func (h UserHandler) PutUpdate(c *fiber.Ctx) error {

	userId := c.Locals("userInfo").(*model_io.User).ID

	user := model_io.User{}
	err := c.BodyParser(&user)
	if err != nil {
		return utils.CusErrorFiber(c, fiber.ErrUnprocessableEntity)
	}

	err = h.userHandle.Update(&user, &userId)
	if err != nil {
		return utils.CusErrorFiber(c, err)
	}

	return c.Status(fiber.StatusOK).JSON(
		fiber.Map{
			"message": "Update Success",
		})
}

func (h UserHandler) DeleteUser(c *fiber.Ctx) error {
	userId := c.Locals("userInfo").(*model_io.User).ID
	h.userHandle.Delete(&userId)
	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"message": "Delete Success",
	})
}

func (h UserHandler) GetMe(c *fiber.Ctx) error {
	user := c.Locals("userInfo").(*model_io.User)
	// user, err := h.userHandle.GetUserById(&userId)
	// if err != nil {
	// 	return utils.CusErrorFiber(c, err)
	// }

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"message": "Get user success",
		"data":    *user,
	})
}

func (h UserHandler) PostLogin(c *fiber.Ctx) error {

	data := model_io.User{}

	err := c.BodyParser(&data)
	if err != nil {
		return utils.CusErrorFiber(c, fiber.ErrUnprocessableEntity)
	}

	token, err := h.userHandle.Login(&data)
	if err != nil {
		return utils.CusErrorFiber(c, err)
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"message": "Login Success",
		"token":   token,
	})
}
