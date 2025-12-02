package handlers

import (
	"attendance-go/src/dtos"
	"attendance-go/src/usecase"

	"github.com/gofiber/fiber/v2"
)

type UserHandler struct {
	userUseCase *usecase.UserUseCase
}

func NewUserHandler(userUseCase *usecase.UserUseCase) *UserHandler {
	return &UserHandler{userUseCase: userUseCase}
}

func (h *UserHandler) CreateUser(c *fiber.Ctx) error {
	var req dtos.UserCreateRequest
	if err := c.BodyParser(&req); err != nil {
		return dtos.ResponseError(c, fiber.StatusBadRequest, err.Error())
	}
	if err := h.userUseCase.Create(&req); err != nil {
		return dtos.ResponseError(c, fiber.StatusInternalServerError, err.Error())
	}
	return dtos.ResponseSuccess(c, nil, "user created")
}
