package main

import (
	"github.com/caiostarke/goAndSvelte/frontend"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/filesystem"
)

func main() {
	app := fiber.New()
	app.Get("/hello.json", handleHello)
	app.Use("/", filesystem.New(filesystem.Config{
		Root:         frontend.BuildHTTPFS(),
		NotFoundFile: "index.html",
	}))
	app.Listen(":8080")
}

func handleHello(c *fiber.Ctx) error {
	return c.JSON(fiber.Map{"message": "hello from fiber server"})
}
