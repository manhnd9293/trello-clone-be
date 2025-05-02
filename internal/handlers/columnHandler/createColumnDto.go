package columnHandler

type CreateColumnDto struct {
	Name string `json:"name" binding:"required"`
}
