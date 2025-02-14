package main

import (
	"log"

	"github.com/MainFoldFunc/Web-sockets-time/src/database"
	"github.com/MainFoldFunc/Web-sockets-time/src/handlers"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
)

func main() {
	// ✅ Connect to the database
	database.ConnDatabase()
	database.ConnDatabaseConv()

	server := fiber.New()

	server.Use(cors.New(cors.Config{
		AllowOrigins:     "http://localhost:5173",
		AllowMethods:     "POST",
		AllowHeaders:     "Content-Type, Authorization",
		AllowCredentials: true,
	}))

	// REST API routes
	server.Post("/api/register", handlers.RegisterHandler)
	server.Post("/api/login", handlers.LoginHandler)
	server.Post("/api/logout", handlers.LogoutHandler)

	server.Post("/api/chatReqest", handlers.ChatReqest)

	port := "8080"
	log.Println("Server running on http://localhost:" + port)
	log.Fatal(server.Listen(":" + port))
}
