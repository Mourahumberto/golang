package controllers

import (
	"apploja/database"
	"apploja/models"
	"strconv"

	"github.com/gofiber/fiber/v3"
)

func AllRoles(c fiber.Ctx) error {
	var roles []models.Role

	database.DB.Find(&roles)

	return c.JSON(roles)
}

type RoleCreateDTO struct {
	Name        string
	permissions []string
}

func CreateRole(c fiber.Ctx) error {
	var roleDto RoleCreateDTO

	if err := c.Bind().Body(&roleDto); err != nil {
		return err
	}
	permissions := make([]models.Permision, len(roleDto.permissions))

	for i, perpermissionId := range roleDto.permissions {
		id, _ := strconv.Atoi(perpermissionId)
		permissions[i] = models.Permision{
			Id: uint(id),
		}
	}
	role := models.Role{
		Name: roleDto.Name,
	}

	database.DB.Create(&role)

	return c.JSON(role)
}

func GetRole(c fiber.Ctx) error {
	id, _ := strconv.Atoi(c.Params("id"))
	role := models.Role{
		Id: uint(id),
	}
	database.DB.Find(&role)
	return c.JSON(role)
}

func UpdateRole(c fiber.Ctx) error {
	id, _ := strconv.Atoi(c.Params("id"))
	role := models.Role{
		Id: uint(id),
	}

	if err := c.Bind().Body(&role); err != nil {
		return err
	}

	database.DB.Model(&role).Updates((role))
	return c.JSON(role)
}

func DeleteRole(c fiber.Ctx) error {
	id, _ := strconv.Atoi(c.Params("id"))
	role := models.Role{
		Id: uint(id),
	}

	database.DB.Delete(&role)
	return nil
}
