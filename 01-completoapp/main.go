package main

import (
	"log"
	"loja/database"
	"loja/routes"

	"github.com/gofiber/fiber/v3"
)

// importrante essa secret não está chumbada na app

func main() {
	database.Connect()
	// Initialize a new Fiber app
	app := fiber.New()
	routes.Setup(app)
	log.Fatal(app.Listen(":3000"))
}
