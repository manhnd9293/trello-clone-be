package models

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type BaseModel struct {
	Id        string         `gorm:"type:uuid;primary_key;" json:"id"`
	CreatedAt time.Time      `json:"created_at"`
	UpdatedAt time.Time      `json:"updated_at"`
	DeletedAt gorm.DeletedAt `gorm:"index" json:"deleted_at"`
}

func (b *BaseModel) BeforeCreate(tx *gorm.DB) (err error) {
	b.Id = uuid.New().String()
	return
}
