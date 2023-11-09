package models

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Task struct {
	gorm.Model
	// -> ID, CreatedAt, UpdatedAt, DeletedAt
	ID       uuid.UUID `gorm:"type:uuid;default:uuid_generate_v4()"`
	Name    string    `gorm:"varchar:191"`
	Description string `gorm:"varchar:300"`
	DeadLine time.Time `gorm:"type:date"`
	IsDone bool `gorm:"type:boolean"`
	ListId	uuid.UUID 
}
