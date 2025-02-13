package main

import (
	"log"

	"github.com/MainFoldFunc/Web-sockets-time/src/database"
	"github.com/MainFoldFunc/Web-sockets-time/src/handlers"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
)

func main() {
	// ✅ Connect to the database before using it
	database.ConnDatabase()

	server := fiber.New()

	server.Use(cors.New(cors.Config{
		AllowOrigins:     "http://localhost:5173",
		AllowMethods:     "POST",
		AllowHeaders:     "Content-Type, Authorization",
		AllowCredentials: true,
	}))

	server.Post("/api/register", handlers.RegisterHandler)
	server.Post("/api/login", handlers.LoginHandler)
	server.Post("/api/logout", handlers.LogoutHandler)

	server.Post("/api/reqest", handlers.ChatReqestHandler)
	server.Post("/api/acceptReqest", handlers.AcceptChatRequestHandler)
	server.Post("/api/declineReqest", handlers.DeclineChatRequestHandler)

	port := "8080"
	log.Fatal(server.Listen(":" + port))
}
