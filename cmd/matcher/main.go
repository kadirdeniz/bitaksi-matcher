package main

import (
	"fmt"
	"matcher/tools/fiber"
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

	// Start router
	fiber.Router()
}
