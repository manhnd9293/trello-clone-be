package taskHandler

type CreateTaskDto struct {
	Name     string `json:"name" binding:"required"`
	ColumnID string `json:"columnId" binding:"required"`
}
