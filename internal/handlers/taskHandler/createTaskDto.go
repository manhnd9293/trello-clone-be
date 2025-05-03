package taskHandler

type CreateTaskDto struct {
	Name     string `json:"name" binding:"required"`
	ColumnId string `json:"columnId" binding:"required"`
}
