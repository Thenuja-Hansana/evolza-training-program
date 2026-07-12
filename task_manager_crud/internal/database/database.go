package database

import (
	"context"
	"log"
	"os"
	"time"

	"github.com/joho/godotenv"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

// Variables for controllers to access the databas
var Client *mongo.Client
var TaskCollection *mongo.Collection

func ConnectDB() {
	// Load the .env file. 
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file: ", err)
	}

	// Read the variables from the OS environment.
	mongoURI := os.Getenv("MONGO_URI")
	dbName := os.Getenv("MONGO_DB_NAME")

	// Set up the MongoDB client options.
	clientOptions := options.Client().ApplyURI(mongoURI)

	// Create context timeout(10s) 
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel() // close connection

	// Connect to the database.
	client, err := mongo.Connect(ctx, clientOptions)
	if err != nil {
		log.Fatal("Failed to connect to MongoDB: ", err)
	}

	// Verify the connection
	err = client.Ping(ctx, nil)
	if err != nil {
		log.Fatal("Failed to ping MongoDB: ", err)
	}

	log.Println("Successfully connected to MongoDB!")

	// 7. Store the client and specific collection in exported variables.
	Client = client
	TaskCollection = client.Database(dbName).Collection("tasks")
}