package handlers

import (
	"fmt"
	"github.com/MainFoldFunc/Web-sockets-time/src/database"
	"github.com/MainFoldFunc/Web-sockets-time/src/models"
	"github.com/gofiber/fiber/v2"
)

func ChatReqestHandler(c *fiber.Ctx) error {
	reqest := new(models.ChatReqest)

	err := c.BodyParser(reqest)
	if err != nil {
		return c.Status(400).JSON(map[string]string{"error": "invalid reqest body"})
	}

	var exsistsReqest models.ChatReqest
	resoult := database.DB.Where("userS = ? AND userR = ? AND status = status quo", reqest.UserS, reqest.UserR).First(&exsistsReqest)
	if resoult.RowsAffected > 0 {
		return c.Status(409).JSON(map[string]string{"error": "reqest already sent"})
	}
	resolt := database.DB.Create(&reqest)
	if resolt.Error != nil {
		return c.Status(500).JSON(map[string]string{"error": "Failed to create a reqest"})
	}

	fmt.Println("Reqest recived")

	return c.Status(201).JSON(map[string]string{"message": "Reqest created"})
}
