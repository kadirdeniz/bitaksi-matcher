package fiber

import (
	"fmt"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"log"
	"matcher/internal"
	"matcher/tools/fiber/handler"
	"matcher/tools/fiber/middleware"
)

func Router() {
	err := StartServer(8000)
	if err != nil {
		log.Fatal(err)
	}
	log.Println("User service started")
}

func StartServer(port int) error {

	// Create repository
	repository := internal.NewRepository()

	// Create handler
	handler := handler.NewHandler(repository)

	app := fiber.New()

	// Cors
	app.Use(cors.New())

	app.Group("/api/v1")

	// Create routes
	app.Get("/drivers/nearest", middleware.IsAuthenticated, handler.GetNearestDriver)
	// Health check
	app.Get("/health", func(c *fiber.Ctx) error {
		return c.Status(fiber.StatusOK).SendString("OK")
	})

	return app.Listen(fmt.Sprintf(":%d", port))
}
