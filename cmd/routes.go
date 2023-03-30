package main

import (
	"github.com/gofiber/fiber/v2"
	"github.com/trulyworthless/chatter-blocks/handlers"
)

func setupRoutes(app *fiber.App) {
	app.Get("/", handlers.Home)
}
