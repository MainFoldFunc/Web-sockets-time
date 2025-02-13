package handlers

import (
	"github.com/MainFoldFunc/Web-sockets-time/src/database"
	"github.com/MainFoldFunc/Web-sockets-time/src/models"
	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt"
)

func DeclineChatRequestHandler(c *fiber.Ctx) error {
	// 1. Retrieve the JWT token from cookies
	tokenString := c.Cookies("jwt")
	if tokenString == "" {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{"error": "unauthorized"})
	}

	// 2. Parse the JWT token
	token, err := jwt.ParseWithClaims(tokenString, &jwt.StandardClaims{}, func(token *jwt.Token) (interface{}, error) {
		return []byte(secretKey), nil
	})
	if err != nil || !token.Valid {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{"error": "unauthorized"})
	}

	// 3. Extract user info from the token claims
	claims, ok := token.Claims.(*jwt.StandardClaims)
	if !ok {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{"error": "unauthorized"})
	}
	currentUserEmail := claims.Issuer

	// 4. Parse the request body for the sender's email
	var body struct {
		UserS string `json:"userS"`
	}
	if err := c.BodyParser(&body); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Invalid request body"})
	}

	// 5. Find the chat request with the specified criteria
	var chatRequest models.ChatReqest
	result := database.DB.Where("user_r = ? AND user_s = ? AND status = ?", currentUserEmail, body.UserS, "1").First(&chatRequest)
	if result.Error != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{"error": "chat request not found"})
	}

	// Option A: Decline by deleting the chat request
	if err := database.DB.Delete(&chatRequest).Error; err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "failed to decline chat request"})
	}

	/* Option B: Decline by updating the status (do not delete the record)
	   chatRequest.Status = "declined"
	   if err := database.DB.Save(&chatRequest).Error; err != nil {
	       return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "failed to decline chat request"})
	   }
	*/

	return c.Status(fiber.StatusOK).JSON(fiber.Map{"message": "chat request declined"})
}
