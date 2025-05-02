package columnHandler

import (
	"net/http"
	"trello-clone/internal/db"
	"trello-clone/internal/models"

	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
)

func CreateColumn(c *gin.Context) {
	data := CreateColumnDto{}

	if err := c.ShouldBindBodyWith(&data, binding.JSON); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "Fail to read column data",
		})
		return
	}

	newColumn := models.Column{Name: data.Name}
	result := db.Connection.Create(&newColumn)

	if result.Error != nil {
		c.JSON(http.StatusBadRequest, result.Error)
		return
	}

	c.JSON(http.StatusCreated, newColumn)
}
