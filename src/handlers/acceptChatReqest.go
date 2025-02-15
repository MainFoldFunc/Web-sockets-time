package handlers

import (
	"fmt"
	"github.com/MainFoldFunc/Web-sockets-time/src/database"
	"github.com/MainFoldFunc/Web-sockets-time/src/models"
	"github.com/gofiber/fiber/v2"
)
import "strings"

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
	request := new(models.AcceptChatReqest)
	if err := c.BodyParser(request); err != nil {
		return c.Status(400).JSON(map[string]string{"error": "Invalid request body"})
	}

	// Ensure the request is made by the intended recipient
	if request.UserEmail != user.Email {
		return c.Status(403).JSON(map[string]string{"error": "You cannot accept requests for someone else"})
	}

	// Find the specific chat request
	var chatReq models.ChatReqest
	result := database.DBC.Where("user_s = ? AND user_r = ?", request.UserS, request.UserEmail).First(&chatReq)

	if result.RowsAffected == 0 {
		return c.Status(404).JSON(map[string]string{"error": "Chat request not found"})
	}

	// Update the request status to "accepted"
	chatReq.Status = "accepted"
	if err := database.DBC.Save(&chatReq).Error; err != nil {
		return c.Status(500).JSON(map[string]string{"error": "Error while updating request"})
	}

	// Sanitize email addresses to remove dots and @
	sanitize := func(email string) string {
		return strings.ReplaceAll(strings.ReplaceAll(email, ".", "_"), "@", "_")
	}

	tableName := fmt.Sprintf("conv_%s_%s", sanitize(request.UserS), sanitize(request.UserEmail))

	// Define the new conversation table structure
	type DynamicConv struct {
		ID    uint   `gorm:"primaryKey"`
		User1 string `gorm:"size:255"`
		User2 string `gorm:"size:255"`
		Body  string `gorm:"size:255"`
	}

	// Create the new table dynamically
	err := database.DBC.Table(tableName).AutoMigrate(&DynamicConv{})
	if err != nil {
		return c.Status(500).JSON(map[string]string{"error": "Error while creating the conversation table"})
	}

	fmt.Println("New conversation table created:", tableName)

	// Success response
	return c.Status(200).JSON(map[string]interface{}{
		"message":    "Request accepted successfully",
		"table_name": tableName,
	})
}
