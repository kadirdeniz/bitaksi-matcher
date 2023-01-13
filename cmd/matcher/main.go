package main

import (
	"fmt"
	_ "matcher/docs"
	"matcher/tools/fiber"
)

func main() {
	fmt.Println("Matcher Service")

	// Start router
	fiber.Router()
}
