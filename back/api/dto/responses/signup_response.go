package responses

import (
	"github.com/google/uuid"
	UserModel "github.com/saegus/test-technique-romain-chenard/internal/modules/user/models"
)

type User struct {
	ID    uuid.UUID
	Email string
	Pseudo string
}

func ToUser(user UserModel.User) User {
	return User{
		ID:    user.ID,
		Email: user.Email, 
		Pseudo:  user.Pseudo,
		
	}
}