package models

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type List struct {
	gorm.Model
	// -> ID, CreatedAt, UpdatedAt, DeletedAt
	ID     uuid.UUID `gorm:"type:uuid;default:uuid_generate_v4()"`
	Name   string    `gorm:"varchar:191"`
	UserId uuid.UUID
	Tasks  []Task `gorm:"foreignKey:ListId"`
}
