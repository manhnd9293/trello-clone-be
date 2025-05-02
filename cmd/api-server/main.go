package main

import (
	"fmt"
	"log"
	"net/http"
	"trello-clone/internal/db"
	"trello-clone/internal/handlers/columnHandler"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"trello-clone/internal/models"
)

func init() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Failed to load environment variables")
	}
	fmt.Println("Initialize project")
	db.ConnectDb()
	db.Connection.AutoMigrate(&models.Column{})
}

func main() {
	router := gin.Default()
	router.GET("/ping", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "pong ping",
		})
	})

	columnRouter := router.Group("/column")
	columnRouter.POST("", columnHandler.CreateColumn)

	router.Run()
}
