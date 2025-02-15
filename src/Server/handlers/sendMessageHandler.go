package handlers

import (
	"fmt"
	"strings"

	"github.com/MainFoldFunc/Web-sockets-time/src/database"
	"github.com/MainFoldFunc/Web-sockets-time/src/models"
	"github.com/gofiber/websocket/v2"
)

// SendMessage handles incoming WebSocket messages.
func SendMessage(c *websocket.Conn) {
	// Retrieve the authenticated user from Locals.
	user, ok := c.Locals("user").(models.Users)
	if !ok {
		fmt.Println("No authenticated user found")
		return
	}

	for {
		var message models.Conv
		if err := c.ReadJSON(&message); err != nil {
			fmt.Println("Invalid JSON body:", err)
			break
		}

		fmt.Println("Received message:\n", message)

		// Sanitize email addresses to remove dots and '@'
		sanitize := func(email string) string {
			return strings.ReplaceAll(strings.ReplaceAll(email, ".", "_"), "@", "_")
		}

		// Construct the conversation-specific table name.
		tableName := fmt.Sprintf("conv_%s_%s", sanitize(message.User1), sanitize(message.User2))

		// Populate the record to save. (Assumes models.SaveToDatabase is defined.)
		var saveToDatabase models.SaveToDatabase
		saveToDatabase.Content = message.Body
		// Give "1" to the user that sent the message, "0" to the other.
		if user.Email == message.User1 {
			saveToDatabase.Sender = 1
		} else {
			saveToDatabase.Sender = 0
		}

		// Save the message in the specific conversation table.
		if err := database.DBC.Table(tableName).Create(&saveToDatabase).Error; err != nil {
			fmt.Println("Error while adding message to the database:", err)
		} else {
			fmt.Println("Message stored successfully")
		}

		if err := c.WriteJSON(map[string]string{"status": "message received"}); err != nil {
			fmt.Println("Error while sending confirmation through the WebSocket:", err)
			break
		}
	}
}
