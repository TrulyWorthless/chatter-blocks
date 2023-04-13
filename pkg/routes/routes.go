package routes

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/trulyworthless/chatter-blocks/pkg/handlers"
	"github.com/trulyworthless/chatter-blocks/pkg/middleware"
)

func InitRoutes(app *fiber.App) {
	// Middleware
	api := app.Group("/", logger.New())

	// Home Route
	api.Get("/home", handlers.Home)

	// Auth
	auth := api.Group("/auth")
	auth.Post("/login", handlers.Login)
	auth.Post("/logout", middleware.Protected(), handlers.Logout)

	// User
	user := api.Group("/user")
	user.Post("/", handlers.CreateUser)
	user.Get("/:username", handlers.GetUser)
	user.Patch("/:username", middleware.Protected(), handlers.UpdateUser)
	user.Delete("/:username", middleware.Protected(), handlers.DeleteUser)

	// Identities Routes
	identity := api.Group("/identity")
	identity.Post("/", handlers.CreateIdentity)
	identity.Get("/list", handlers.GetIdentities)
	identity.Get("/alias/:alias", handlers.GetIdentityByAlias)
	identity.Put("/alias/:alias", handlers.UpdateIdentityByAlias)
	// identity.Delete("/alias/:alias", middleware.Protected(), handlers.DeleteIdentityByAlias)
	identity.Delete("/alias/:alias", handlers.DeleteIdentityByAlias)

	// Contacts Routes
	contact := api.Group("/contact")
	contact.Post("/", handlers.CreateContact)
	contact.Get("/list", handlers.GetContacts)
	contact.Get("/correspondent/:correspondent", handlers.GetContactByCorrespondent)
	contact.Put("/correspondent/:correspondent", handlers.UpdateContactByCorrespondent)
	// contact.Delete("/alias/:alias", middleware.Protected(), handlers.DeleteContactByAlias)
	contact.Delete("/correspondent/:correspondent", handlers.DeleteContactByCorrespondent)

	//Export
	//TODO Add remove files
	export := api.Group("/export")
	export.Post("/pubkey/alias/:alias", handlers.ExportPublicKeyByAlias)
	export.Post("/address/alias/:alias", handlers.ExportAddressByAlias)
}
