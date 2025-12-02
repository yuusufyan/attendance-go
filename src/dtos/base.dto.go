package dtos

import (
	"github.com/gofiber/fiber/v2"
)

type PaginationRequest struct {
	Page    int    `query:"page" json:"page"`
	PerPage int    `query:"perPage" json:"per_page"`
	SortBy  string `query:"sortBy" json:"sortBy"`
	OrderBy string `query:"orderBy" json:"orderBy"`
	Limit   int    `query:"limit" json:"limit"`
	Search  string `query:"search" json:"search,omitempty"`
}

type Response struct {
	Status  string      `json:"status"`
	Message string      `json:"message"`
	Data    interface{} `json:"data,omitempty"`
}

type PaginationResponse[T any] struct {
	Total     int64 `json:"total"`
	TotalPage int   `json:"total_page"`
	Page      int   `json:"page"`
	PerPage   int   `json:"per_page"`
	Data      []T   `json:"data"`
}

func ResponseSuccess(c *fiber.Ctx, data interface{}, message string) error {
	resp := Response{
		Status:  "success",
		Message: message,
		Data:    data,
	}
	return c.Status(fiber.StatusOK).JSON(resp)
}

func ResponseError(c *fiber.Ctx, httpCode int, message string) error {
	return c.Status(httpCode).JSON(Response{
		Status:  "error",
		Message: message,
		Data:    nil,
	})
}
