package main

import (
	"log"
	"trello-clone/internal/db"
	"trello-clone/internal/models"

	"github.com/joho/godotenv"
)

func init() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Failed to load environment variables")
	}
	db.ConnectDb()
}

func main() {
	db.Connection.AutoMigrate(&models.Column{})
}
