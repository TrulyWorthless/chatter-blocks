package main

import (
	"github.com/gofiber/fiber/v2"
	"github.com/trulyworthless/chatter-blocks/pkg/database"
	"github.com/trulyworthless/chatter-blocks/pkg/routes"
	"github.com/trulyworthless/chatter-blocks/pkg/web3"
)

func main() {
	//create db connection
	database.New()

	//create ethclient
	web3.New()

	//create fiber
	app := fiber.New()

	//init routes
	routes.InitRoutes(app)

	//start app
	app.Listen(":3000")
}
