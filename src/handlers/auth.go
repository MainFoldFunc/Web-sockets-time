package handlers

import (
	"github.com/MainFoldFunc/Web-sockets-time/src/models"
	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt"
	"strings"
)

func Authenticate(c *fiber.Ctx) error {
	// Get token from the cookie or Authorization header
	tokenString := c.Cookies("jwt")
	if tokenString == "" {
		// If the cookie is empty, check the Authorization header
		tokenString = c.Get("Authorization")
	}

	// If no token, return an error
	if tokenString == "" {
		return c.Status(401).JSON(map[string]string{"error": "Authorization token missing"})
	}

	// If the token is in the Authorization header, remove the "Bearer " prefix
	if strings.HasPrefix(tokenString, "Bearer ") {
		tokenString = tokenString[len("Bearer "):]
	}

	// Parse and validate the token
	claims := jwt.StandardClaims{}
	token, err := jwt.ParseWithClaims(tokenString, &claims, func(token *jwt.Token) (interface{}, error) {
		return []byte(secretKey), nil
	})

	// If there's an error parsing the token, return an error
	if err != nil || !token.Valid {
		return c.Status(401).JSON(map[string]string{"error": "Invalid token"})
	}

	// Set the authenticated user into the context
	user := models.Users{Email: claims.Issuer} // Assuming the email is stored in the Issuer claim
	c.Locals("user", user)

	return nil
}
