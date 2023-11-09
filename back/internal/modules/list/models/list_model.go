package models

import (
	"github.com/google/uuid"
	"github.com/saegus/test-technique-romain-chenard/internal/modules/task/models"
	"gorm.io/gorm"
)

type List struct {
	gorm.Model
	// -> ID, CreatedAt, UpdatedAt, DeletedAt
	ID       uuid.UUID `gorm:"type:uuid;default:uuid_generate_v4()"`
	Name    string    `gorm:"varchar:191"`
	UserId	uuid.UUID
	Tasks  []models.Task `gorm:"foreignKey:ListId"`
}
