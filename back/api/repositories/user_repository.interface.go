package repositories

import (
	models "github.com/saegus/test-technique-romain-chenard/data/models"
)

type UserRepositoryInterface interface {
	CreateUser(user models.User) (models.User, error)
	FindUserByEmail(email string) (models.User, error)
}
