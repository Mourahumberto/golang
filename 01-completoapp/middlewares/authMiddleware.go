package middlewares

import (
	"fmt"
	"loja/util"

	"github.com/gofiber/fiber/v3"
)

func IsAuthenticated(c fiber.Ctx) error {
	// cookie := c.Cookies("jwt")

	// if _, err := util.ParseJwt(cookie); err != nil {
	// 	c.Status(fiber.StatusUnauthorized)
	// 	return c.JSON(fiber.Map{
	// 		"message": "unauthenticated",
	// 	})
	// }

	// return c.Next()
	tokenString := c.Get("Authorization")
	if tokenString == "" {
		// w.WriteHeader(http.StatusUnauthorized)
		fmt.Printf("Missing authorization header")
		return c.JSON(fiber.Map{
			"message": "sem token",
		})
	}
	tokenString = tokenString[len("Bearer "):]

	err := util.VerifyToken(tokenString)
	if err != nil {
		return c.JSON(fiber.Map{
			"message": "sem autorização",
		})
	}
	return c.Next()

}
