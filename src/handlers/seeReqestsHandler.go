package handlers

import (
	"fmt"
	"github.com/MainFoldFunc/Web-sockets-time/src/database"
	"github.com/MainFoldFunc/Web-sockets-time/src/models"
	"github.com/gofiber/fiber/v2"
)

func SeeChatReqestsHandler(c *fiber.Ctx) error {
	fmt.Println("See chat requests handler called")
	if err := Authenticate(c); err != nil {
		fmt.Println("Error while auth")
		return err
	}

	// Get the authenticated user
	user := c.Locals("user").(models.Users)

	// Parse the request body into the SeeChatReqests model
	reqest := new(models.SeeChatReqests)
	err := c.BodyParser(reqest)
	if err != nil {
		return c.Status(400).JSON(map[string]string{"error": "Invalid JSON body"})
	}

	// Ensure that the email in the request matches the authenticated user
	if reqest.UserEmail != user.Email {
		return c.Status(403).JSON(map[string]string{"error": "You can only see requests for yourself"})
	}

	// Query for all requests where the user is either the sender or the receiver
	var allRequests []models.ChatReqest
	result := database.DBC.Where("user_r = ?", reqest.UserEmail).Find(&allRequests)

	// Check if there were any errors during the query
	if result.Error != nil {
		return c.Status(500).JSON(map[string]string{"error": "Error fetching chat requests"})
	}

	// Return the chat requests
	return c.Status(200).JSON(map[string]interface{}{
		"message":      "Successfully fetched chat requests",
		"chatRequests": allRequests,
	})
}
