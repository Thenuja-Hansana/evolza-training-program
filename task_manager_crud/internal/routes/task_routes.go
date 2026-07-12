package routes

import (
	"github.com/gofiber/fiber/v2"

	// Ensure this matches your exact go.mod module name
	"github.com/Thenuja-Hansana/evolza-training-program/task_manager_crud/internal/controllers"
)

// SetupTaskRoutes organizes all task-related endpoints
func SetupTaskRoutes(app *fiber.App) {
	// Create a route group to keep things clean.
	// Every route inside this group automatically starts with "/tasks"
	taskGroup := app.Group("/tasks")

	// Create
	taskGroup.Post("/", controllers.CreateTask)       // POST /tasks

	// Read
	taskGroup.Get("/", controllers.GetTasks)          // GET /tasks
	taskGroup.Get("/:id", controllers.GetTask)        // GET /tasks/:id

	// Update
	taskGroup.Put("/:id", controllers.UpdateTask)     // PUT /tasks/:id

	// Delete
	taskGroup.Delete("/:id", controllers.DeleteTask)  // DELETE /tasks/:id
}