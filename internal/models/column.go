package models

type Column struct {
	BaseModel
	Name     string `json:"name"`
	Tasks    []Task `json:"tasks"`
	Position int    `json:"position" gorm:"default:0"`
}
