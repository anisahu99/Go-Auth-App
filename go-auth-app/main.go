package main

import (
	"log"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/recover"
	"github.com/joho/godotenv"

	"go-auth-app/handlers"
	"go-auth-app/database"
)

func main() {
	// Load environment variables
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	// Initialize the database connection
	database.Connect()

	// Create a new Fiber app
	app := fiber.New()

	// Middleware
	app.Use(cors.New())
	app.Use(logger.New())  // Logs requests
	app.Use(recover.New()) // Recover from panics

	// Routes
	app.Post("/register", handlers.Register)
	app.Post("/login", handlers.Login)
	app.Get("/", handlers.Home)

	// Start the server
	log.Fatal(app.Listen(":3000"))
}
