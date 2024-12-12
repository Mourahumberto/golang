package routes

import (
	"apploja/controllers"
	"apploja/middlewares"

	"github.com/gofiber/fiber/v3"
)

func Setup(app *fiber.App) {
	app.Post("/api/register", controllers.Register)
	app.Post("/api/login", controllers.Login)

	app.Use(middlewares.IsAuthenticated)
	app.Get("/api/user", controllers.User)
	app.Get("/api/logout", controllers.Logout)

	app.Get("/api/users", controllers.AllUsers)
	app.Post("/api/users", controllers.CreateUser)
	app.Get("/api/user/:id", controllers.GetUser)
	app.Put("/api/user/:id", controllers.UpdateUser)
	app.Delete("/api/user/:id", controllers.DeleteUser)

	app.Get("/api/roles", controllers.AllRoles)
	app.Post("/api/roles", controllers.CreateRole)
	app.Get("/api/role/:id", controllers.GetRole)
	app.Put("/api/role/:id", controllers.UpdateRole)
	app.Delete("/api/role/:id", controllers.DeleteRole)

	app.Get("/api/permissions", controllers.AllPermissions)
}
