package handlers

import (
	"fmt"
	"github.com/MainFoldFunc/Web-sockets-time/src/database"
	"github.com/MainFoldFunc/Web-sockets-time/src/models"
	"github.com/gofiber/fiber/v2"
)

func ChatReqest(c *fiber.Ctx) error {
	// Use Authenticate middleware to ensure the user is authenticated
	fmt.Println("Handler called")
	if err := Authenticate(c); err != nil {
		fmt.Println("Error while auth")
		return err
	}

	// Get the authenticated user from the context
	user := c.Locals("user").(models.Users)

	// Parse the request body into the ChatReqest model
	reqest := new(models.ChatReqest)
	err := c.BodyParser(reqest)
	if err != nil {
		return c.Status(400).JSON(map[string]string{"error": "invalid json body"})
	}

	// Check if the user making the request is the same as either userS or userR (you can adjust this logic)
	if reqest.UserS != user.Email {
		// Optionally, you could return an error if the authenticated user isn't the sender (userS)
		return c.Status(403).JSON(map[string]string{"error": "You can only create requests as yourself"})
	}

	// Check if the receiver exists in the database
	var inUsers models.Users
	resouult := database.DB.Where("email= ?", reqest.UserR).First(&inUsers)
	if resouult.RowsAffected == 0 { // Corrected condition here
		return c.Status(400).JSON(map[string]string{"error": "There is no such user"})
	}

	// Check if a request already exists
	var existsReqest models.ChatReqest
	result := database.DBC.Where("userS = ? AND userR = ?", reqest.UserS, reqest.UserR).First(&existsReqest)
	if result.RowsAffected > 0 {
		return c.Status(400).JSON(map[string]string{"error": "Request already exists"})
	}

	// Add the request to the database
	resolt := database.DBC.Create(&reqest)
	if resolt.Error != nil {
		return c.Status(500).JSON(map[string]string{"error": "Error while adding request to a database"})
	}

	// Log and return success
	fmt.Println("New Request added.")
	return c.Status(201).JSON(map[string]interface{}{
		"message": "Request added successfully",
	})
}
