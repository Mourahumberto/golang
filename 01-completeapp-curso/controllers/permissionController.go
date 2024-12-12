package controllers

import (
	"apploja/database"
	"apploja/models"

	"github.com/gofiber/fiber/v3"
)

func AllPermissions(c fiber.Ctx) error {
	var permissions []models.Permision

	database.DB.Find(&permissions)

	return c.JSON(permissions)
}
