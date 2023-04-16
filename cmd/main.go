package main

import (
	"fmt"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/trulyworthless/chatter-blocks/pkg/database"
	"github.com/trulyworthless/chatter-blocks/pkg/routes"
	"github.com/trulyworthless/chatter-blocks/pkg/web3"
)

func main() {
	//create db connection
	database.InitDb()

	//create fiber
	app := fiber.New()
	app.Use(cors.New())

	//init routes
	routes.InitRoutes(app)

	//start rpc client
	web3.Init()

	//start app
	fmt.Println("Listening to port 3000")
	app.Listen(":3000")
}
