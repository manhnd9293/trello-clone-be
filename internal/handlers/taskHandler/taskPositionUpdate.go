package taskHandler

type TaskPositionUpdate struct {
	Id       string `json:"id" binding:"required"`
	ColumnId string `json:"columnId" binding:"required"`
	Position int    `json:"position" binding:"min=0"`
}
