package controllers

import (
	"context"
	"time"

	"github.com/gofiber/fiber/v2"
	"go.mongodb.org/mongo-driver/bson/primitive"

	"go.mongodb.org/mongo-driver/bson"

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

	// Context with  for the database operation
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	// Insert the task into the database 
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


// GetTasks retrieves all tasks from the database
func GetTasks(c *fiber.Ctx) error {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	var tasks []models.Task
	
	// bson.M{} means "match everything"
	cursor, err := database.TaskCollection.Find(ctx, bson.M{})
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Failed to fetch tasks"})
	}
	defer cursor.Close(ctx)

	// Decode all found documents into our tasks slice
	if err = cursor.All(ctx, &tasks); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Failed to parse tasks"})
	}

	return c.JSON(fiber.Map{
		"status": "success",
		"data":   tasks,
	})
}

// GetTask retrieves a single task by its ID
func GetTask(c *fiber.Ctx) error {
	id := c.Params("id") // Extract the ID from the URL

	// Convert the string ID into a MongoDB ObjectID
	objID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Invalid task ID"})
	}

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	var task models.Task
	err = database.TaskCollection.FindOne(ctx, bson.M{"_id": objID}).Decode(&task)
	if err != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{"error": "Task not found"})
	}

	return c.JSON(fiber.Map{
		"status": "success",
		"data":   task,
	})
}

// UpdateTask completely updates an existing task
func UpdateTask(c *fiber.Ctx) error {
	id := c.Params("id")
	objID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Invalid task ID"})
	}

	var updateData models.Task
	if err := c.BodyParser(&updateData); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Failed to parse JSON body"})
	}

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	// Define exactly what fields we want to update
	update := bson.M{
		"$set": bson.M{
			"title":       updateData.Title,
			"description": updateData.Description,
			"status":      updateData.Status,
			"updated_at":  time.Now(),
		},
	}

	result, err := database.TaskCollection.UpdateOne(ctx, bson.M{"_id": objID}, update)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Failed to update task"})
	}

	if result.MatchedCount == 0 {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{"error": "Task not found"})
	}

	return c.JSON(fiber.Map{
		"status":  "success",
		"message": "Task updated successfully",
	})
}

// DeleteTask removes a task from the database
func DeleteTask(c *fiber.Ctx) error {
	id := c.Params("id")
	objID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Invalid task ID"})
	}

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	result, err := database.TaskCollection.DeleteOne(ctx, bson.M{"_id": objID})
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Failed to delete task"})
	}

	if result.DeletedCount == 0 {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{"error": "Task not found"})
	}

	return c.JSON(fiber.Map{
		"status":  "success",
		"message": "Task deleted successfully",
	})
}