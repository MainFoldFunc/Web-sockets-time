package handlers

import (
	"fmt"

	"github.com/MainFoldFunc/Web-sockets-time/src/database"
	"github.com/MainFoldFunc/Web-sockets-time/src/models"
	"github.com/gofiber/fiber/v2"
)

func DeclineChatReqestHandler(c *fiber.Ctx) error {
	fmt.Println("Reqest chat handler called")

	if err := Authenticate(c); err != nil {
		fmt.Println("Error while auth")
		return err
	}

	user := c.Locals("user").(models.Users)

	reqest := new(models.AcceptChatReqest)
	if err := c.BodyParser(reqest); err != nil {
		return c.Status(400).JSON(map[string]string{"error": "Invalid reqest body"})
	}

	if reqest.UserEmail != user.Email {
		return c.Status(403).JSON(map[string]string{"error": "You cannot accept reqest as anotherperson"})
	}

	var chatReq models.ChatReqest
	result := database.DBC.Where("user_s = ? AND user_r = ?", reqest.UserS, reqest.UserEmail).First(&chatReq)
	if result.RowsAffected == 0 {
		return c.Status(404).JSON(map[string]string{"error": "reqest not found"})
	}

	chatReq.Status = "accepted"
	if err := database.DBC.Save(&chatReq).Error; err != nil {
		return c.Status(500).JSON(map[string]string{"error": "Error while updating status of the reqest"})
	}

	return c.Status(200).JSON(map[string]interface{}{
		"message": "Reqest updated succesfully",
	})
}
