package handlers

import (
	"attendance-go/src/dtos"
	"attendance-go/src/services"
	"strconv"

	"github.com/gofiber/fiber/v2"
)

type UserHandler struct {
	userService *services.UserService
}

func NewUserHandler(userService *services.UserService) *UserHandler {
	return &UserHandler{userService: userService}
}

func (h *UserHandler) CreateUser(c *fiber.Ctx) error {
	var req dtos.UserCreateRequest
	if err := c.BodyParser(&req); err != nil {
		return dtos.ResponseError(c, fiber.StatusBadRequest, err.Error(), nil)
	}
	if err := h.userService.Create(&req); err != nil {
		return dtos.ResponseError(c, fiber.StatusInternalServerError, err.Error(), nil)
	}
	return dtos.ResponseSuccess(c, "user created", nil)
}

func (h *UserHandler) GetAllUser(c *fiber.Ctx) error {
	var req dtos.PaginationRequest
	if err := c.QueryParser(&req); err != nil {
		return dtos.ResponseError(c, fiber.StatusBadRequest, err.Error(), nil)
	}
	response, err := h.userService.GetAll(req.Page, req.PerPage, req.Search, req.SortBy, req.OrderBy)
	if err != nil {
		return dtos.ResponseError(c, fiber.StatusInternalServerError, err.Error(), nil)
	}
	return dtos.ResponseSuccess(c, "user retrieved", response)
}

func (h *UserHandler) GetByID(c *fiber.Ctx) error {
	id, err := strconv.Atoi(c.Params("id"))
	if err != nil {
		return dtos.ResponseError(c, fiber.StatusBadRequest, err.Error(), nil)
	}
	response, err := h.userService.GetByID(uint(id))
	if err != nil {
		return dtos.ResponseError(c, fiber.StatusInternalServerError, err.Error(), nil)
	}
	return dtos.ResponseSuccess(c, "user retrieved", response)
}

func (h *UserHandler) UpdateByID(c *fiber.Ctx) error {
	id, err := strconv.Atoi(c.Params("id"))
	if err != nil {
		return dtos.ResponseError(c, fiber.StatusBadRequest, err.Error(), nil)
	}
	var req dtos.UserResponse
	if err := c.BodyParser(&req); err != nil {
		return dtos.ResponseError(c, fiber.StatusBadRequest, err.Error(), nil)
	}
	response, err := h.userService.UpdateByID(uint(id), &req)
	if err != nil {
		return dtos.ResponseError(c, fiber.StatusInternalServerError, err.Error(), nil)
	}
	return dtos.ResponseSuccess(c, "user updated", response)
}
