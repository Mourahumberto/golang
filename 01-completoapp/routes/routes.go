package routes

import (
	"loja/controllers"
	"loja/middlewares"

	"github.com/gofiber/fiber/v3"
)

func Setup(app *fiber.App) {
	app.Post("/login", controllers.LoginHandler)
	app.Use(middlewares.IsAuthenticated)
	app.Get("/api/products", controllers.Products)
	app.Post("/api/product", controllers.CreateProduct)
	app.Get("/api/product/:id", controllers.GetProduct)
	app.Put("/api/product/:id", controllers.UpdateProduct)
	app.Delete("/api/product/:id", controllers.DeleteProduct)

}
