package routes

import (
	"github.com/gofiber/fiber/v2"
	"github.com/trulyworthless/chatter-blocks/pkg/handlers"
)

func InitRoutes(app *fiber.App) {
	// Home Route
	app.Get("/", handlers.Home)

	//TODO
	// Login Route
	app.Post("/login", handlers.Login)

	//Facts Routes
	app.Get("/listfacts", handlers.ListFacts)
	app.Post("/createfact", handlers.CreateFact)

	// Identities Routes
	app.Post("/identities", handlers.CreateIdentity)
	app.Get("/identities", handlers.GetIdentities)
	app.Get("/identities/{uuid}", handlers.GetIdentity)
	app.Put("/identities/{uuid}", handlers.UpdateIdentity)
	app.Delete("/identities/{uuid}", handlers.DeletIdentity)

	// Contacts Routes
	app.Post("/contacts", handlers.CreateContact)
	app.Get("/contacts", handlers.GetContacts)
	app.Get("/contacts/{uuid}", handlers.GetContact)
	app.Put("/contacts/{uuid}", handlers.UpdateContact)
	app.Delete("/contacts/{uuid}", handlers.DeleteContact)
}
