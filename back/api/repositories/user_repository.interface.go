package repositories

import (
	models "github.com/saegus/test-technique-romain-chenard/internal/modules/user/models"
)

type UserRepositoryInterface interface {
	CreateUser(user models.User) (models.User, error)
	FindUserByEmail(email string) (models.User, error)
}
