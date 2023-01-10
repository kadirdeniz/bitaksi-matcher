package main

import (
	"fmt"
	swagger "github.com/arsmn/fiber-swagger/v2"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"matcher/internal"

	"matcher/tools/fiber/handler"
	"matcher/tools/fiber/middleware"
)

// @title Matcher API
// @description This is a sample server for a matcher service.
// @version 1.0.0
// @BasePath /api/v1
// @schemes http
// @host localhost:8000
// @Accept json
// @Produce json
func main() {
	fmt.Println("Matcher Service")

	// Create repository
	repository := internal.NewRepository()

	// Create handler
	handler := handler.NewHandler(repository)

	app := fiber.New()

	// Cors
	app.Use(cors.New())

	// Swagger
	app.Get("/swagger/*", swagger.HandlerDefault)

	api := app.Group("/api/v1")

	// Create routes
	api.Get("/drivers/nearest", middleware.IsAuthenticated, handler.GetNearestDriver)
	// Health check
	api.Get("/health", func(c *fiber.Ctx) error {
		return c.Status(fiber.StatusOK).SendString("OK")
	})

	app.Listen(fmt.Sprintf(":%d", 8000))
}
