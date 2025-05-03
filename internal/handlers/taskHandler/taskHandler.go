package taskHandler

import (
	"net/http"
	"trello-clone/internal/initializer"

	"trello-clone/internal/models"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func CreateTask(c *gin.Context) {
	data := CreateTaskDto{}

	if err := c.ShouldBindJSON(&data); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "Fail to read create task data",
		})
		return
	}

	column := models.Column{}
	checkColumn := initializer.Db.Where(&models.BaseModel{Id: data.ColumnId}).Find(&column)
	if checkColumn.Error != nil {
		c.JSON(http.StatusNotFound, checkColumn.Error)

		return
	}
	var countTask int64
	initializer.Db.Model(&models.Task{}).Where(&models.Task{ColumnId: data.ColumnId}).Count(&countTask)

	newTask := models.Task{Name: data.Name, ColumnId: data.ColumnId, Position: int(countTask)}
	result := initializer.Db.Model(&models.Task{}).Create(&newTask)

	if result.Error != nil {
		c.JSON(http.StatusBadRequest, result.Error)
		return
	}

	c.JSON(http.StatusCreated, newTask)
}

func UpdateTaskPosition(c *gin.Context) {
	updateData := TaskPositionUpdate{}
	if err := c.ShouldBindJSON(&updateData); err != nil {
		c.JSON(http.StatusBadRequest, err.Error())
		return
	}

	task := models.Task{}
	findTask := initializer.Db.Where("id = ?", updateData.Id).First(&task)

	if err := findTask.Error; err != nil {
		c.JSON(http.StatusBadRequest, err.Error())
		return
	}

	updateColumn := models.Column{}
	findUpdateColumn := initializer.Db.Where(&models.BaseModel{Id: updateData.ColumnId}).Find(&updateColumn)
	if err := findUpdateColumn.Error; err != nil {
		c.JSON(http.StatusBadRequest, err.Error())
		return
	}
	if task.ColumnId == updateData.ColumnId {
		if task.Position < updateData.Position {
			resultUpdate := initializer.Db.Model(&models.Task{}).Where("? < position AND position <= ? AND column_id = ?", task.Position, updateData.Position, task.ColumnId).UpdateColumn("position", gorm.Expr("position - ?", 1))
			if err := resultUpdate.Error; err != nil {
				c.JSON(http.StatusBadRequest, err.Error())
				return
			}
		} else {
			resultUpdate := initializer.Db.Model(&models.Task{}).Where("? <= position AND position < ? AND column_id = ?", updateData.Position, task.Position, task.ColumnId).UpdateColumn("position", gorm.Expr("position + ?", 1))
			if err := resultUpdate.Error; err != nil {
				c.JSON(http.StatusBadRequest, err.Error())
				return
			}
		}
	} else {
		resultUpdateCurrentColumn := initializer.Db.Model(&models.Task{}).Where("position > ? AND column_id = ?", task.Position, task.ColumnId).UpdateColumn("position", gorm.Expr("position - ?", 1))
		if err := resultUpdateCurrentColumn.Error; err != nil {
			c.JSON(http.StatusBadRequest, err.Error())
			return
		}

		resultNewUpdate := initializer.Db.Model(&models.Task{}).Where("position >= ? AND column_id = ?", updateData.Position, updateData.ColumnId).UpdateColumn("position", gorm.Expr("position + ?", 1))
		if err := resultNewUpdate.Error; err != nil {
			c.JSON(http.StatusBadRequest, err.Error())
			return
		}
	}

	initializer.Db.Model(&task).Select("Position", "ColumnId").Updates(models.Task{Position: updateData.Position, ColumnId: updateData.ColumnId})

	c.JSON(http.StatusAccepted, task)
}
