package main

import (
	"fmt"
	"log"
	"myGram/entities"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var (
	host     = "localhost"
	user     = "postgres"
	password = "admin"
	dbPort   = "5432"
	dbName   = "mygram"
	db       *gorm.DB
	err      error
)

func init() {
	config := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable",
		host, user, password, dbName, dbPort)

	db, err = gorm.Open(postgres.Open(config), &gorm.Config{})
	if err != nil {
		log.Fatal("Error connecting to database : ", err)
	}

	db.Debug().AutoMigrate(entities.User{}, entities.Photo{}, entities.Comment{}, entities.Socialmedia{})
}

func main() {
	// var PORT = ":4000"
	// routers.StartService(db).Run(PORT)
}
