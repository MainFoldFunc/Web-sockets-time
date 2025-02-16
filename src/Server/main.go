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

	// ✅ Improved CORS Configuration
	app.Use(cors.New(cors.Config{
		AllowOrigins:     "http://192.168.1.19:5173", // No trailing slash
		AllowMethods:     "GET,POST,OPTIONS",         // Added OPTIONS
		AllowHeaders:     "Content-Type, Authorization",
		AllowCredentials: true,
	}))

	// ✅ REST API Endpoints
	app.Post("/api/register", handlers.RegisterHandler)
	app.Post("/api/login", handlers.LoginHandler)
	app.Post("/api/logout", handlers.LogoutHandler)

	app.Post("/api/searchUsers", handlers.SearchForUsersHandler)

	app.Post("/api/chatRequest", handlers.ChatReqest)
	app.Post("/api/seeChatRequests", handlers.SeeChatReqestsHandler)
	app.Post("/api/acceptChatRequest", handlers.AcceptChatReqest)
	app.Post("/api/declineChatRequest", handlers.DeclineChatReqestHandler)

	// ✅ WebSocket Authentication Middleware for Sending Messages
	app.Use("/api/sendMessage", func(c *fiber.Ctx) error {
		if err := handlers.Authenticate(c); err != nil {
			return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{"error": "unauthorized"})
		}
		return c.Next()
	})

	// ✅ WebSocket Authentication Middleware for Receiving Messages
	app.Use("/api/getMessages", func(c *fiber.Ctx) error {
		if err := handlers.Authenticate(c); err != nil {
			return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{"error": "unauthorized"})
		}
		return c.Next()
	})

	// ✅ WebSocket Endpoints
	app.Get("/api/sendMessage", websocket.New(handlers.SendMessage))
	app.Get("/api/getMessages", websocket.New(handlers.GetMessage))

	// ✅ Start Server
	port := "8080"
	log.Println("🚀 Server running on http://localhost:" + port)
	log.Fatal(app.Listen("192.168.1.19:" + port))
}
