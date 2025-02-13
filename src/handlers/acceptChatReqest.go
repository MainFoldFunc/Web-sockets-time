package handlers

import (
	"fmt"

	"github.com/MainFoldFunc/Web-sockets-time/src/database"
	"github.com/MainFoldFunc/Web-sockets-time/src/models"
	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt"
)

// AcceptChatRequestHandler allows a user to accept a chat request and create a new conversation.
func AcceptChatRequestHandler(c *fiber.Ctx) error {
	// 1. Verify the JWT token from the "jwt" cookie.
	tokenString := c.Cookies("jwt")
	if tokenString == "" {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"error": "unauthenticated",
		})
	}

	token, err := jwt.ParseWithClaims(tokenString, &jwt.StandardClaims{}, func(token *jwt.Token) (interface{}, error) {
		return []byte(secretKey), nil
	})
	if err != nil || !token.Valid {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"error": "unauthenticated",
		})
	}

	claims, ok := token.Claims.(*jwt.StandardClaims)
	if !ok {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"error": "unauthenticated",
		})
	}

	// The current authenticated user's email is stored in the Issuer field.
	currentUserEmail := claims.Issuer

	// 2. Parse the JSON body.
	var body struct {
		UserS string `json:"userS"`
		User1 string `json:"user1"`
		User2 string `json:"user2"`
	}
	if err := c.BodyParser(&body); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "invalid request body",
		})
	}

	// 3. Find the pending chat request.
	var chatRequest models.ChatReqest
	result := database.DB.Where("user_r = ? AND user_s = ? AND status = ?", currentUserEmail, body.UserS, "1").First(&chatRequest)
	if result.Error != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"error": "chat request not found",
		})
	}

	// 4. Update the chat request's status to "accepted".
	chatRequest.Status = "accepted"
	if err := database.DB.Save(&chatRequest).Error; err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "failed to accept chat request",
		})
	}

	// 5. Check if a conversation already exists.
	var existsConv models.Conv
	results := database.DBCONVS.Where("user1 = ? AND user2 = ?", body.User1, body.User2).First(&existsConv)
	if results.RowsAffected > 0 {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "conversation already exists",
		})
	}

	// 6. Create a new conversation.
	conv := models.Conv{
		User1: body.User1,
		User2: body.User2,
	}
	if err := database.DBCONVS.Create(&conv).Error; err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "failed to create conversation",
		})
	}
	fmt.Println("Conversation created")

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"message": "chat request accepted and conversation created",
	})
}
