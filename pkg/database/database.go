package database

import (
	"fmt"
	"log"
	"os"

	"github.com/trulyworthless/chatter-blocks/pkg/config"
	"github.com/trulyworthless/chatter-blocks/pkg/database/models"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

var Db *gorm.DB

func InitDb() {
	fmt.Println("create db connection")
	//configs
	dsn := fmt.Sprintf(
		"host=db user=%s password=%s dbname=%s port=5432 sslmode=disable TimeZone=America/Los_Angeles",
		config.Config("DB_USER"),
		config.Config("DB_PASSWORD"),
		config.Config("DB_NAME"),
	)

	//open gorm
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Info),
	})

	//error check
	if err != nil {
		log.Fatal("Failed to connect to database. \n", err)
		os.Exit(1)
	}

	//add logger
	log.Println("connected to db")
	db.Logger = logger.Default.LogMode(logger.Info)

	//migrate
	log.Println("running migrations")
	db.AutoMigrate(&models.User{}, &models.Identity{}, &models.Contact{})

	Db = db
}
