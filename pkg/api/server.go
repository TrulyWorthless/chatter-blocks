package api

import (
	"fmt"
	"os"

	"github.com/joho/godotenv"
	"github.com/trulyworthless/chatter-blocks/pkg/api/controllers"
	"github.com/trulyworthless/chatter-blocks/pkg/api/seed"
)

var server = controllers.Server{}

func Run() {
	err := godotenv.Load()
	if err != nil {
		panic(err)
	} else {
		fmt.Println("env loaded")
	}

	server.Initialize(os.Getenv("DB_DRIVER"), os.Getenv("DB_USER"), os.Getenv("DB_PASSWORD"), os.Getenv("DB_PORT"), os.Getenv("DB_HOST"), os.Getenv("DB_NAME"))
	seed.Load(server.DB)
	server.Run(":8080")
}
