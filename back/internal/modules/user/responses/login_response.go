package responses

import (
	"github.com/google/uuid"
)

type LoginResponse struct {
	ID uuid.UUID `json:"id"`
	Email string `json:"email"`
	Token string `json:"token"`
}