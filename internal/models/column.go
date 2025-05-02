package models

type Column struct {
	BaseModel
	Name  string `json:"name"`
	Tasks []Task `json:"tasks"`
}
