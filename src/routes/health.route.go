package routes

import (
	"os"

	"attendance/src/handlers"

	"github.com/gofiber/fiber/v2"
)

func CheckHealthModule(app fiber.Router) {
	// Init repository
	// healthRepo := repositories.NewCheckHealthRepository()
	handler := handlers.NewHealthHandler(os.Getenv("SERVICE_NAME"), os.Getenv("VERSION"))
	controller := controllers.NewHealthController(handler)

	// Init routes
	app.Get("/", controller.Check)
}
