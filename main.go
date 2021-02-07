package main

import (
	"github.com/gofiber/fiber/v2"
)

func helloWorld(c *fiber.Ctx) error {
	return c.SendString("Hello, World!")
}

func setupRoutes(app *fiber.App) {
	app.Get("/", helloWorld)
}

func main() {
	app := fiber.New()

	setupRoutes(app)
	app.Listen(":3000")
}
