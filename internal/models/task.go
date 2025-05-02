package models

type Task struct {
	BaseModel
	Name     string `json:"name"`
	ColumnID string `json:"columnId"`
	Column   Column `gorm:"foreignKey:ColumnID;references:ID" json:"column"`
}
