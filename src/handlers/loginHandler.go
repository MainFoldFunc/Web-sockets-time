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

	err := c.BodyParser(userLogin)
	if err != nil {
		return c.Status(400).JSON(map[string]string{"error": "Invalid reqest body"})
	}

	var user models.Users
	result := database.DB.Where("email = ?", userLogin.Email).First(&user)
	if result.Error != nil {
		return c.Status(404).JSON((map[string]string{"error": "User not found"}))
	}

	if user.Password != userLogin.Password {
		return c.Status(401).JSON(map[string]string{"error": "Incorrect passsword"})
	}
	fmt.Println("User logged in: ", user.Email)

	claims := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.StandardClaims{
		Issuer:    user.Email,
		ExpiresAt: time.Now().Add(time.Hour * 24).Unix(),
	})

	token, err := claims.SignedString([]byte(secretKey))
	if err != nil {
		return c.Status(500).JSON(map[string]string{"error": "couldn't generate a JWT token"})
	}

	cookie := fiber.Cookie{
		Name:     "jwt",
		Value:    token,
		Expires:  time.Now().Add(time.Hour * 24),
		HTTPOnly: true,
	}
	c.Cookie(&cookie)

	return c.Status(200).JSON(map[string]interface{}{
		"message": "Login succesful",
	})
}
