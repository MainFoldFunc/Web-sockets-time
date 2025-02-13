package handlers

import (
	"github.com/gofiber/fiber/v2"
	"time"
)

func LogoutHandler(c *fiber.Ctx) error {
	cookie := fiber.Cookie{
		Name:     "jwt",
		Value:    "",
		Expires:  time.Now().Add(-time.Hour),
		HTTPOnly: true,
	}
	c.Cookie(&cookie)
	return c.JSON(map[string]string{"message": "succes"})
}
