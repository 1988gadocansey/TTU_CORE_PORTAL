package main

import (
	"TTU_CORE_ERP_PORTAL/database"
	"TTU_CORE_ERP_PORTAL/routes"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
)

func main() {

	database.Connect()

	app := fiber.New()

	app.Use(cors.New(cors.Config{
		AllowCredentials: true,
	}))

	routes.Create(app)
	routes.Setup(app)

	app.Listen(":8000")
}
