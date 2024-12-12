package controllers

// LoginService to provide user login with JWT token support
import (
	"fmt"
	"loja/util"

	"github.com/gofiber/fiber/v3"
)

type User struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

func LoginHandler(c fiber.Ctx) error {
	var u User
	if err := c.Bind().Body(&u); err != nil {
		return err
	}

	fmt.Printf("The user request value %v \n", u)

	if u.Username == "Chek" && u.Password == "123456" {
		tokenString, err := util.CreateToken(u.Username)
		if err != nil {
			c.Status(500)
			fmt.Errorf("No username found", err)
		}
		c.Status(200)
		return c.JSON(tokenString)
	}
	c.Status(401)
	return c.JSON(fiber.Map{
		"message": "Invalid credentials",
	})
}
