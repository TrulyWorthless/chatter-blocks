package main

import (
	"github.com/gofiber/fiber/v2"
	"github.com/trulyworthless/chatter-blocks/database"
)

func main() {
	database.ConnectDb()
	app := fiber.New()

	setupRoutes(app)

	app.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("Div Rhino Trivia App!")
	})

	app.Listen(":3000")
}
