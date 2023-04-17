package controllers

import (
	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
	database "github.com/jabahum/todo-backend/db"
	"github.com/jabahum/todo-backend/model"
)

func GetTodos(c *fiber.Ctx) error {
	db := database.DB
	var todo []model.Todo

	// find all todo in the database
	db.Find(&todo)

	// If no todo is present return an error
	if len(todo) == 0 {
		return c.Status(404).JSON(fiber.Map{"status": "error", "message": "No todo present", "data": nil})
	}

	// Else return todo
	return c.JSON(fiber.Map{"status": "success", "message": "todo Found", "data": todo})
}

func CreateTodos(c *fiber.Ctx) error {
	db := database.DB
	todo := new(model.Todo)

	// Store the body in the todo and return error if encountered
	err := c.BodyParser(todo)
	if err != nil {
		return c.Status(500).JSON(fiber.Map{"status": "error", "message": "Review your input", "data": err})
	}
	// Add a uuid to the todo
	todo.ID = uuid.New()
	// Create the todo and return error if encountered
	err = db.Create(&todo).Error
	if err != nil {
		return c.Status(500).JSON(fiber.Map{"status": "error", "message": "Could not create todo", "data": err})
	}

	// Return the created todo
	return c.JSON(fiber.Map{"status": "success", "message": "Created todo", "data": todo})
}

func GetTodo(c *fiber.Ctx) error {
	db := database.DB
	var todo model.Todo

	// Read the param todoId
	id := c.Params("todoId")

	// Find the todo with the given Id
	db.Find(&todo, "id = ?", id)

	// If no such todo present return an error
	if todo.ID == uuid.Nil {
		return c.Status(404).JSON(fiber.Map{"status": "error", "message": "No todo present", "data": nil})
	}

	// Return the todo with the Id
	return c.JSON(fiber.Map{"status": "success", "message": "todo Found", "data": todo})
}

func UpdateTodo(c *fiber.Ctx) error {
	type updatetodo struct {
		Title    string `json:"title"`
		SubTitle string `json:"sub_title"`
		Text     string `json:"Text"`
	}
	db := database.DB
	var todo model.Todo

	// Read the param todoId
	id := c.Params("todoId")

	// Find the todo with the given Id
	db.Find(&todo, "id = ?", id)

	// If no such todo present return an error
	if todo.ID == uuid.Nil {
		return c.Status(404).JSON(fiber.Map{"status": "error", "message": "No todo present", "data": nil})
	}

	// Store the body containing the updated data and return error if encountered
	var updatetodoData updatetodo
	err := c.BodyParser(&updatetodoData)
	if err != nil {
		return c.Status(500).JSON(fiber.Map{"status": "error", "message": "Review your input", "data": err})
	}

	// Edit the todo
	todo.Title = updatetodoData.Title
	todo.SubTitle = updatetodoData.SubTitle
	todo.Text = updatetodoData.Text

	// Save the Changes
	db.Save(&todo)

	// Return the updated todo
	return c.JSON(fiber.Map{"status": "success", "message": "todo Found", "data": todo})
}

func DeleteTodo(c *fiber.Ctx) error {
	db := database.DB
	var todo model.Todo

	// Read the param todoId
	id := c.Params("todoId")

	// Find the todo with the given Id
	db.Find(&todo, "id = ?", id)

	// If no such todo present return an error
	if todo.ID == uuid.Nil {
		return c.Status(404).JSON(fiber.Map{"status": "error", "message": "No todo present", "data": nil})
	}

	// Delete the todo and return error if encountered
	err := db.Delete(&todo, "id = ?", id).Error

	if err != nil {
		return c.Status(404).JSON(fiber.Map{"status": "error", "message": "Failed to delete todo", "data": nil})
	}

	// Return success message
	return c.JSON(fiber.Map{"status": "success", "message": "Deleted todo"})
}
