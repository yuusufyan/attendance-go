package handlers

import (
	"attendance-go/src/dtos"
	"attendance-go/src/services"
	"context"
	"time"

	"github.com/gofiber/fiber/v2"
)

type HealthHandler struct {
	healthSvc *services.HealthService
}

func NewHealthHandler(svc *services.HealthService) *HealthHandler {
	return &HealthHandler{
		healthSvc: svc,
	}
}

func (h *HealthHandler) HealthCheck(c *fiber.Ctx) error {
	ctx, cancel := context.WithTimeout(c.UserContext(), 10*time.Second)
	defer cancel()

	details, isHealthy := h.healthSvc.CheckReadiness(ctx)
	if !isHealthy {
		return dtos.ResponseError(c, 503, "Application is not ready", details)
	}

	return dtos.ResponseSuccess(c, "Application is ready", details)
}
