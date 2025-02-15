package handlers

import (
	"fmt"
	"github.com/MainFoldFunc/Web-sockets-time/src/database"
	"github.com/MainFoldFunc/Web-sockets-time/src/models"
	"github.com/gofiber/fiber/v2"
)

func RegisterHandler(c *fiber.Ctx) error {
	user := new(models.Users)

	err := c.BodyParser(user)
	if err != nil {
		return c.Status(400).JSON(map[string]string{"error": "Invalid request body"})
	}

	var existUser models.Users
	result := database.DB.Where("email = ?", user.Email).First(&existUser)

	// ✅ Correct way to check if the email already exists
	if result.RowsAffected > 0 {
		return c.Status(400).JSON(map[string]string{"error": "Email already in use"})
	}

	resolt := database.DB.Create(&user)
	if resolt.Error != nil {
		return c.Status(500).JSON(map[string]string{"error": "Failed to create user"})
	}
	fmt.Println("User signed up: ", user.Email)

	return c.Status(201).JSON(map[string]interface{}{
		"message": "User registered successfully",
		"user":    user,
	})
}
