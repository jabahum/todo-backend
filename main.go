package main

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/logger"
	db "github.com/jabahum/todo-backend/db"
	"github.com/jabahum/todo-backend/router"
)

func main() {
	app := fiber.New()

	app.Use(logger.New())
	app.Use(cors.New())

	// Connect to the Database
	db.ConnectDB()

	// handle unavailable route
	app.Use(func(c *fiber.Ctx) error {
		return c.SendStatus(404) // => 404 "Not Found"
	})

	// Setup the router
	router.SetUpRoutes(app)

	// Listen on PORT 300
	app.Listen(":3000")
}
