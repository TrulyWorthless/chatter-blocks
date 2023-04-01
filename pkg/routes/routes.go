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
	app.Delete("/identity/id/:id", handlers.DeleteIdentityByID)
	app.Delete("/identity/alias/:alias", handlers.DeleteIdentityByAlias)

	// Contacts Routes
	app.Post("/contact", handlers.CreateContact)
	app.Get("/contacts", handlers.GetContacts)
	app.Get("/contact/id/:id", handlers.GetContactById)
	app.Get("/contact/alias/:alias", handlers.GetContactByAlias)
	app.Put("/contact/id/:id", handlers.UpdateContactById)
	app.Put("/contact/:alias", handlers.UpdateContactByAlias)
	app.Delete("/contact/id/:id", handlers.DeleteContactById)
	app.Delete("/contact/alias/:alias", handlers.DeleteContactByAlias)

	//Export
	app.Post("/export/pubkey/id/:id", handlers.ExportPublicKeyById)
	app.Post("/export/pubkey/alias/:alias", handlers.ExportPublicKeyByAlias)
	app.Post("/export/address/id/:id", handlers.ExportAddressById)
	app.Post("/export/address/alias/:alias", handlers.ExportAddressByAlias)
}
