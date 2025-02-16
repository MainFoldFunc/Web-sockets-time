package handlers

import (
	"fmt"
	"github.com/MainFoldFunc/Web-sockets-time/src/database"
	"github.com/MainFoldFunc/Web-sockets-time/src/models"
	"github.com/gofiber/fiber/v2"
)

func SearchForUsersHandler(c *fiber.Ctx) error {
	fmt.Println("Handler search for users called")

	// Authenticate the user
	if err := Authenticate(c); err != nil {
		fmt.Println("Error while authenticating")
		return c.Status(403).JSON(map[string]string{"error": "No auth"})
	}

	// Parse the search query from the body
	userToSearch := new(models.SearchForUsersBar)
	err := c.BodyParser(userToSearch)
	if err != nil {
		return c.Status(400).JSON(map[string]string{"error": "Invalid JSON body"})
	}

	// Fetch users matching the search query
	var users []models.Users
	result := database.DB.Where("email LIKE ?", "%"+userToSearch.Email+"%").Find(&users)
	if result.Error != nil {
		return c.Status(404).JSON(map[string]string{"error": "No users found matching the search"})
	}

	// Return the list of users as a JSON response
	return c.Status(200).JSON(map[string]interface{}{
		"message": "Users found",
		"users":   users,
	})
}
