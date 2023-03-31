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

	// Identities Routes
	app.Post("/identity", handlers.CreateIdentity)
	app.Get("/identities", handlers.GetIdentities)
	app.Get("/identity/id/:id", handlers.GetIdentityById)
	app.Get("/identity/alias/:alias", handlers.GetIdentityByAlias)
	app.Put("/identity/id/:id", handlers.UpdateIdentityByID)
	app.Put("/identity/alias/:alias", handlers.UpdateIdentityByAlias)
	app.Delete("/identity/id/:id", handlers.DeletIdentityByID)
	app.Delete("/identity/alias/:alias", handlers.DeletIdentityByAlias)

	// Contacts Routes
	app.Post("/contacts", handlers.CreateContact)
	app.Get("/contacts", handlers.GetContacts)
	app.Get("/identity/id/:id", handlers.GetContactById)
	app.Get("/identity/alias/:alias", handlers.GetContactByAlias)
	app.Put("/contacts/id/:id", handlers.UpdateContactById)
	app.Put("/contacts/:alias", handlers.UpdateContactByAlias)
	app.Delete("/contacts/id/:id", handlers.DeleteContactById)
	app.Delete("/contacts/alias/:alias", handlers.DeleteContactByAlias)

	//Export
}
