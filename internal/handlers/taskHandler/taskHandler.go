package taskHandler

import (
	"net/http"
	"trello-clone/internal/db"

	"trello-clone/internal/models"

	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
)

func CreateTask(c *gin.Context) {
	data := CreateTaskDto{}

	if err := c.ShouldBindBodyWith(&data, binding.JSON); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "Fail to read create task data",
		})
		return
	}

	column := models.Column{}
	checkColumn := db.Connection.Where(&models.BaseModel{ID: data.ColumnID}).Find(&column)
	if checkColumn.Error != nil {
		c.JSON(http.StatusNotFound, checkColumn.Error)

		return
	}

	newTask := models.Task{Name: data.Name, ColumnID: data.ColumnID}
	result := db.Connection.Model(&models.Task{}).Create(&newTask)

	if result.Error != nil {
		c.JSON(http.StatusBadRequest, result.Error)
		return
	}

	c.JSON(http.StatusCreated, newTask)
}
