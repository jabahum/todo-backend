package router

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
	route "github.com/jabahum/todo-backend/controller"
)

func SetUpRoutes(app *fiber.App) {
	api := app.Group("/api", logger.New())

	todo := api.Group("/todo")

	todo.Post("/", route.CreateTodos)

	todo.Get("/", route.GetTodos)

	todo.Get("/:todoId", route.GetTodo)

	todo.Put("/:todoId", route.UpdateTodo)

	todo.Delete("/:todoId", route.DeleteTodo)
}
