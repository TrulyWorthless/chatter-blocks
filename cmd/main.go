package main

import (
	"fmt"

	"github.com/gofiber/fiber/v2"
	"github.com/trulyworthless/chatter-blocks/pkg/database"
	"github.com/trulyworthless/chatter-blocks/pkg/routes"
)

func main() {
	//create db connection
	database.InitDb()

	//create fiber
	app := fiber.New()

	//init routes
	routes.InitRoutes(app)

	//TODO Seed DB

	//start app
	fmt.Println("Listening to port 3000")
	app.Listen(":3000")
}
