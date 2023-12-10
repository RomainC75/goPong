package requests

import "time"

type CreateTaskRequest struct {
	Name string `json:"name" binding:"required"`
	Description string `json:"description" binding:"required"`
	DeadLine time.Time `json:"deadLine" binding:"required"`
}

