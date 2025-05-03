package models

type Task struct {
	BaseModel
	Name     string `json:"name"`
	ColumnId string `json:"columnId"`
	Column   Column `gorm:"foreignKey:ColumnId;references:Id" json:"column"`
	Position int    `json:"position" gorm:"default:0"`
}
