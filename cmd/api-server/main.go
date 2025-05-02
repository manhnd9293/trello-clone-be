package main

import (
	"fmt"
	"log"
	"net/http"
	"trello-clone/internal/db"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func init() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Failed to load environment variables")
	}
	fmt.Println("Initialize project")
	db.ConnectDb()
}

func main() {
	router := gin.Default()
	router.GET("/ping", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "pong ping",
		})
	})

	// columnRouter := router.Group("column")

	router.Run()
}
