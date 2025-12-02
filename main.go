package main

import (
	"attendance-go/src/configs"
	"attendance-go/src/routes"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
)

func main() {
	app := fiber.New()

	app.Use(logger.New(logger.Config{
		Format:     "[${time}] ${ip} - ${status} ${method} ${path}\n",
		TimeFormat: "02-Jan-2006 15:04:05",
		TimeZone:   "Asia/Jakarta",
	}))

	configs.ConnectDB()
	db := configs.DB

	app.Get("/", func(ctx *fiber.Ctx) error {
		return ctx.Send([]byte("Welcome to IPROC API"))
	})
	routes.SetupRoutes(app, db)

	app.Listen(":3000")
}
