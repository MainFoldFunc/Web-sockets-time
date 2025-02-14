package handlers

import (
	"fmt"
	"github.com/MainFoldFunc/Web-sockets-time/src/database"
	"github.com/MainFoldFunc/Web-sockets-time/src/models"
	"github.com/gofiber/fiber/v2"
)

func AcceptChatReqest(c *fiber.Ctx) error {
	fmt.Println("Accept chat request handler called")

	// Authenticate user
	if err := Authenticate(c); err != nil {
		fmt.Println("Error while auth")
		return err
	}

	// Get authenticated user
	user := c.Locals("user").(models.Users)

	// Parse request body
	reqest := new(models.AcceptChatReqest)
	if err := c.BodyParser(reqest); err != nil {
		return c.Status(400).JSON(map[string]string{"error": "Invalid request body"})
	}

	// Ensure the request is made by the intended recipient
	if reqest.UserEmail != user.Email {
		return c.Status(403).JSON(map[string]string{"error": "You cannot accept requests for someone else"})
	}

	// Find the specific chat request
	var chatReq models.ChatReqest
	result := database.DBC.Where("user_s = ? AND user_r = ?", reqest.UserS, reqest.UserEmail).First(&chatReq)

	if result.RowsAffected == 0 {
		return c.Status(404).JSON(map[string]string{"error": "Chat request not found"})
	}

	// Update the request status to "accepted"
	chatReq.Status = "accepted"
	if err := database.DBC.Save(&chatReq).Error; err != nil {
		return c.Status(500).JSON(map[string]string{"error": "Error while updating request"})
	}

	// Success response
	return c.Status(200).JSON(map[string]interface{}{
		"message": "Request accepted successfully",
	})
}
