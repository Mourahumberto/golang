package main

import (
	"apploja/database"
	"apploja/routes"
	"log"

	"github.com/gofiber/fiber/v3"
)

func main() {
	database.Connect()
	// Initialize a new Fiber app
	app := fiber.New()

	//config for the frontend don't get de cokies
	// app.Use(cors.New(cors.Config{
	// 	AllowCredentials: true,
	// 	AllowOrigins:     []string{"http://localhost"},
	// 	AllowHeaders:     []string{"Origin", "Content-Type", "Accept"},
	// }))
	routes.Setup(app)
	// Define a route for the GET method on the root path '/'

	// Start the server on port 3000
	log.Fatal(app.Listen(":3000"))
}
