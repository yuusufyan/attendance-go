package routes

import (
	"attendance-go/src/handlers"
	"attendance-go/src/repositories"
	"attendance-go/src/services"

	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

func UserModule(app fiber.Router, db *gorm.DB) {
	userRepository := repositories.NewUserRepository(db)
	userService := services.NewUserService(userRepository)
	userHandler := handlers.NewUserHandler(userService)

	// API
	user := app.Group("/user")
	user.Post("/", userHandler.CreateUser)
	user.Get("/", userHandler.GetAllUser)
	user.Get("/:id", userHandler.GetByID)
}
