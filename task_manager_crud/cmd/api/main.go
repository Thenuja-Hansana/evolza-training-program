package main

import (
	"log"

	"os"

	"github.com/gofiber/fiber/v2"

	"github.com/Thenuja-Hansana/evolza-training-program/task_manager_crud/internal/database"

	"github.com/Thenuja-Hansana/evolza-training-program/task_manager_crud/internal/controllers"

	"github.com/Thenuja-Hansana/evolza-training-program/task_manager_crud/internal/routes"
)

func main() {
	// Intialise the database connection
	database.ConnectDB()


	// intialize the server
	app := fiber.New()

	// create a health check request 
	// "/health" - endpoint 
	app.Get("/health", func (c *fiber.Ctx) error {
		// set "Content-Type: application/json"
		return c.JSON(fiber.Map{
			"status" : "success",
			"message" : "Task manager API running smoothly",
		})
	})

	// Task Routes
	app.Post("/tasks", controllers.CreateTask)

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080" // fallback 
	}

	// 4. THIS IS THE CRITICAL LINE: It loads all your CRUD routes
	routes.SetupTaskRoutes(app)

	// start the server 
	log.Fatal(app.Listen(":" + port))
	
}