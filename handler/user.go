package handler

import (
	"github.com/gofiber/fiber/v2"
	"ozanpay/model"
	"ozanpay/service"
	"ozanpay/viewmodel"
)

type UserHandler struct {
	userService service.IUserService
}

func NewUserHandler(userService service.IUserService) *UserHandler {
	return &UserHandler{
		userService: userService,
	}
}

func (h *UserHandler) Create(c *fiber.Ctx) error {
	var userCreateVM viewmodel.UserCreateVM
	if err := c.BodyParser(&userCreateVM); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": err.Error(),
		})
	}
	m := userCreateVM.ToModel(model.User{})

	err := h.userService.Create(c.Context(), &m)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	return c.Status(fiber.StatusCreated).JSON(m)
}
