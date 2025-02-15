package main

import (
	"log"

	"github.com/MainFoldFunc/Web-sockets-time/src/database"
	"github.com/MainFoldFunc/Web-sockets-time/src/handlers"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/websocket/v2"
)

func main() {
	// ✅ Connect to the databases
	database.ConnDatabase()
	database.ConnDatabaseConv()

	app := fiber.New()

	// ✅ CORS Configuration
	app.Use(cors.New(cors.Config{
		AllowOrigins:     "http://localhost:5173",
		AllowMethods:     "GET,POST",
		AllowHeaders:     "Content-Type, Authorization",
		AllowCredentials: true,
	}))

	// ✅ REST API Endpoints
	app.Post("/api/register", handlers.RegisterHandler)
	app.Post("/api/login", handlers.LoginHandler)
	app.Post("/api/logout", handlers.LogoutHandler)

	app.Post("/api/chatReqest", handlers.ChatReqest)
	app.Post("/api/seeChatReqests", handlers.SeeChatReqestsHandler)
	app.Post("/api/acceptChatReqest", handlers.AcceptChatReqest)
	app.Post("/api/declineChatReqest", handlers.DeclineChatReqestHandler)

	// ✅ WebSocket Authentication Middleware
	app.Use("/api/sendMessage", func(c *fiber.Ctx) error {
		if err := handlers.Authenticate(c); err != nil {
			return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{"error": "unauthorized"})
		}
		return c.Next()
	})

	// ✅ WebSocket Endpoint
	app.Get("/api/sendMessage", websocket.New(handlers.SendMessage))

	// ✅ Start Server
	port := "8080"
	log.Println("Server running on http://localhost:" + port)
	log.Fatal(app.Listen(":" + port))
}
