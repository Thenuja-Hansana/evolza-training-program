package controllers

import (
	"context"
	"time"

	"github.com/gofiber/fiber/v2"
	"go.mongodb.org/mongo-driver/bson/primitive"

	"github.com/Thenuja-Hansana/evolza-training-program/task_manager_crud/internal/database"
	"github.com/Thenuja-Hansana/evolza-training-program/task_manager_crud/internal/models"
)

// POST requests to create a new task
func CreateTask(c *fiber.Ctx) error {
	// Create an empty instance of Task struct
	var task models.Task

	// Parse the incoming JSON body directly into our struct
	if err := c.BodyParser(&task); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Failed to parse JSON body",
		})
	}

	// Validation
	if task.Title == "" {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Task title is required",
		})
	}

	// Assign default values before saving
	task.ID = primitive.NewObjectID() // Generate a new MongoDB ObjectID
	task.CreatedAt = time.Now()
	task.UpdatedAt = time.Now()
	
	if task.Status == "" {
		task.Status = "pending" // Default status
	}

	// Create a context with  for the database operation
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	// Insert the task into the database using global TaskCollection
	_, err := database.TaskCollection.InsertOne(ctx, task)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Failed to save task to the database",
		})
	}

	// 7. Return a response along with the task data
	return c.Status(fiber.StatusCreated).JSON(fiber.Map{
		"status": "success",
		"data":   task,
	})
}
