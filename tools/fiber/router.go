package fiber

import (
	"fmt"
	swagger "github.com/arsmn/fiber-swagger/v2"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/recover"

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

	app := fiber.New(
		fiber.Config{
			ErrorHandler: handler.ErrorHandler,
		},
	)

	// Cors
	app.Use(cors.New())

	// Swagger
	app.Get("/swagger/*", swagger.HandlerDefault)

	// Recovery
	app.Use(recover.New())

	api := app.Group("/api/v1")

	// Create routes
	api.Get("/drivers/nearest", middleware.IsAuthenticated, handler.GetNearestDriver)
	// Health check
	api.Get("/health", func(c *fiber.Ctx) error {
		return c.Status(fiber.StatusOK).SendString("OK")
	})

	return app.Listen(fmt.Sprintf(":%d", port))
}
