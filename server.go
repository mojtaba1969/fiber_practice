package main

import "github.com/gofiber/fiber/v2"

func main() {
	app := fiber.New(fiber.Config{
		CaseSensitive: true,
		StrictRouting: false,
	})

	app.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("Hello, World!")
	})

	app.Get("/test", func(c *fiber.Ctx) error {
		return c.SendString("This is a test!")
	})

	app.Listen(":3000")
}

// go mod init test
// go mod tidy
// go get -u github.com/gofiber/fiber/v2
