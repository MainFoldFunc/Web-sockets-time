package main

import (
	"log"
	"os"
	"path/filepath"

	"github.com/MainFoldFunc/Web-sockets-time/src/database"
	"github.com/MainFoldFunc/Web-sockets-time/src/handlers"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/websocket/v2"
	"github.com/joho/godotenv"
)

func main() {
	// ✅ Load .env from two directories up
	envPath, _ := filepath.Abs("../../.env")
	err := godotenv.Load(envPath)
	if err != nil {
		log.Fatal("Error while opening the .env file:", err)
	}

	// ✅ Connect to the databases
	database.ConnDatabase()
	database.ConnDatabaseConv()

	app := fiber.New()

	allowOrigins := "http://" + os.Getenv("PORT_ORIGIN")

	// ✅ Improved CORS Configuration
	app.Use(cors.New(cors.Config{
		AllowOrigins:     allowOrigins,       // No trailing slash
		AllowMethods:     "GET,POST,OPTIONS", // Added OPTIONS
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
	port := os.Getenv("PORT")
	log.Println("🚀 Server running on:", port)
	log.Fatal(app.Listen(port)) // ✅ Correct
}
