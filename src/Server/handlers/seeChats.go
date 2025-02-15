package handlers

import (
	"fmt"
	"strings"
	"time"

	"github.com/MainFoldFunc/Web-sockets-time/src/database"
	"github.com/MainFoldFunc/Web-sockets-time/src/models"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/websocket/v2"
)

func GetMessage(c *websocket.Conn) {
	// Authenticate user
	_, ok := c.Locals("user").(models.Users)
	if !ok {
		fmt.Println("No authenticated user found")
		c.WriteJSON(fiber.Map{"error": "unauthorized"})
		return
	}

	// Extract conversation details
	var query struct {
		User1 string `json:"user1"`
		User2 string `json:"user2"`
	}

	if err := c.ReadJSON(&query); err != nil {
		fmt.Println("Invalid JSON body:", err)
		c.WriteJSON(fiber.Map{"error": "invalid request"})
		return
	}

	// Sanitize email addresses for table name
	sanitize := func(email string) string {
		return strings.ReplaceAll(strings.ReplaceAll(email, ".", "_"), "@", "_")
	}

	// Construct table name based on sanitized emails
	tableName := fmt.Sprintf("conv_%s_%s", sanitize(query.User1), sanitize(query.User2))

	// Ensure WebSocket connection is closed once done
	defer c.Close()

	// Continuously check for new messages
	for {
		var messages []models.Conv

		// Query the database for messages
		if err := database.DBC.Table(tableName).Find(&messages).Error; err != nil {
			// Log error and break if there's an issue fetching messages
			fmt.Println("Error fetching messages:", err)
			c.WriteJSON(fiber.Map{"error": "error fetching messages"})
			break
		}

		// Send messages to the WebSocket if any were found
		if len(messages) > 0 {
			if err := c.WriteJSON(messages); err != nil {
				// Log error and break if there's an issue sending the messages
				fmt.Println("Error sending messages:", err)
				c.WriteJSON(fiber.Map{"error": "error sending messages"})
				break
			}
		}

		// Sleep for 2 seconds before checking for new messages
		time.Sleep(2 * time.Second)
	}
}
