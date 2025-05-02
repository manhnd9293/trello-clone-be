package main

import (
	"fmt"
	"log"
	"net/http"
	"trello-clone/internal/db"
	"trello-clone/internal/handlers/columnHandler"
	"trello-clone/internal/handlers/taskHandler"

	"trello-clone/internal/models"

	"github.com/gin-contrib/cors"
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
	db.Connection.AutoMigrate(&models.Column{}, &models.Task{})
}

func main() {
	router := gin.Default()
	router.GET("/ping", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "pong ping",
		})
	})
	router.Use(cors.Default())

	columnRouter := router.Group("/columns")
	columnRouter.POST("", columnHandler.CreateColumn)
	columnRouter.GET("", columnHandler.GetAllColumns)

	taskRouter := router.Group("/tasks")
	taskRouter.POST("", taskHandler.CreateTask)

	router.Run()
}
