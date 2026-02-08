package routes

import (
	"attendance-go/src/handlers"
	"attendance-go/src/services"

	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

func HealthModule(app fiber.Router, db *gorm.DB) {
	healthService := services.NewHealthService(db)
	healthHandler := handlers.NewHealthHandler(healthService)

	app.Get("/", healthHandler.HealthCheck)
}
