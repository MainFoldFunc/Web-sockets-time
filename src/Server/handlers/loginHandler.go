package handlers

import (
	"fmt"
	"time"

	"github.com/MainFoldFunc/Web-sockets-time/src/database"
	"github.com/MainFoldFunc/Web-sockets-time/src/models"
	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt"
)

const secretKey = "secretWoohoo"

func LoginHandler(c *fiber.Ctx) error {
	userLogin := new(models.UserLogin)

	// Parse request body
	err := c.BodyParser(userLogin)
	if err != nil {
		return c.Status(400).JSON(map[string]string{"error": "Invalid request body"})
	}

	// Find user in the database
	var user models.Users
	result := database.DB.Where("email = ?", userLogin.Email).First(&user)
	if result.Error != nil {
		return c.Status(404).JSON(map[string]string{"error": "User not found"})
	}

	// Check password
	if user.Password != userLogin.Password {
		return c.Status(401).JSON(map[string]string{"error": "Incorrect password"})
	}
	fmt.Println("User logged in: ", user.Email)

	// Generate JWT token
	claims := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.StandardClaims{
		Issuer:    user.Email,
		ExpiresAt: time.Now().Add(time.Hour * 24).Unix(),
	})

	token, err := claims.SignedString([]byte(secretKey))
	if err != nil {
		return c.Status(500).JSON(map[string]string{"error": "Couldn't generate a JWT token"})
	}

	// ✅ Instead of setting a cookie, return the JWT token in JSON response
	return c.Status(200).JSON(map[string]interface{}{
		"message": "Login successful",
		"token":   token,
	})
}
