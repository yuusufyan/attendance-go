package routes

import (
	"attendance-go/src/handlers"
	"attendance-go/src/repositories"
	"attendance-go/src/usecase"

	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

func UserModule(app fiber.Router, db *gorm.DB) {
	userRepository := repositories.NewUserRepository(db)
	userUseCase := usecase.NewUserUseCase(userRepository)
	userHandler := handlers.NewUserHandler(userUseCase)

	// API
	user := app.Group("/user")
	user.Post("/", userHandler.CreateUser)
	user.Get("/", userHandler.GetAllUser)
}
