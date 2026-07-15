package main

import (
	"log"

	"os"

	"context"

	"time"

	"github.com/gofiber/fiber/v2"

	"github.com/Thenuja-Hansana/evolza-training-program/task_manager_crud/internal/database"

	"github.com/Thenuja-Hansana/evolza-training-program/task_manager_crud/internal/controllers"

	"github.com/Thenuja-Hansana/evolza-training-program/task_manager_crud/internal/routes"

	"github.com/gofiber/fiber/v2/middleware/cors"

	"github.com/Thenuja-Hansana/evolza-training-program/task_manager_crud/internal/models"

	"go.mongodb.org/mongo-driver/bson"
)

func main() {
	// Intialise the database connection
	database.ConnectDB()

	database.ConnectDB()

	startPendingTaskChecker()

	// intialize the server
	app := fiber.New()

	app.Use(cors.New())

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

	routes.SetupTaskRoutes(app)

	// start the server 
	log.Fatal(app.Listen(":" + port))
	
}


func startPendingTaskChecker() {
	ticker := time.NewTicker(30 * time.Second)
	go func() {
		for range ticker.C {
			checkPendingTasks()
		}
	}()
}

func checkPendingTasks() {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	cursor, err := database.TaskCollection.Find(ctx, bson.M{"status": "pending"})
	if err != nil {
		log.Println("Error checking pending tasks:", err)
		return
	}
	defer cursor.Close(ctx)

	var pendingTasks []models.Task
	if err := cursor.All(ctx, &pendingTasks); err != nil {
		log.Println("Error decoding pending tasks:", err)
		return
	}

	if len(pendingTasks) == 0 {
		log.Println("No pending tasks found.")
		return
	}

	log.Println("Pending tasks found:")
	for _, t := range pendingTasks {
		log.Printf("- %s (id: %s)\n", t.Title, t.ID.Hex())
	}
}