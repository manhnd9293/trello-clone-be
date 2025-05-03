package columnHandler

import (
	"net/http"

	"trello-clone/internal/initializer"

	"trello-clone/internal/models"

	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"gorm.io/gorm"
)

func CreateColumn(c *gin.Context) {
	data := CreateColumnDto{}

	if err := c.ShouldBindBodyWith(&data, binding.JSON); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "Fail to read column data",
		})
		return
	}
	var countColunm int64
	initializer.Db.Model(&models.Column{}).Count(&countColunm)

	newColumn := models.Column{Name: data.Name, Position: int(countColunm)}
	result := initializer.Db.Create(&newColumn)

	if result.Error != nil {
		c.JSON(http.StatusBadRequest, result.Error)
		return
	}

	c.JSON(http.StatusCreated, newColumn)
}

func GetAllColumns(c *gin.Context) {
	var columns []models.Column
	result := initializer.Db.Preload("Tasks", func(db *gorm.DB) *gorm.DB {
		return db.Order("tasks.position ASC")
	}).Order("position ASC").Find(&columns)
	if result.Error != nil {
		c.JSON(http.StatusBadRequest, result.Error.Error())
		return
	}

	c.JSON(http.StatusAccepted, columns)
}
