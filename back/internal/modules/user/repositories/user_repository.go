package repositories

import (
	"errors"

	models "github.com/saegus/test-technique-romain-chenard/internal/modules/user/models"
	database "github.com/saegus/test-technique-romain-chenard/pkg/database"
	"gorm.io/gorm"
)

type UserRepository struct {
	DB *gorm.DB
}

func New() *UserRepository{
	return  &UserRepository{
		DB: database.Connection(),
	}
}

func (UserRepository *UserRepository) CreateUser(user models.User) (models.User, error){
	var newUser models.User
	result := UserRepository.DB.Create(&user).Scan(&newUser)
	if result.RowsAffected == 0 {
		return models.User{}, errors.New("error trying to creat a new user")
	}
	return newUser, nil
}

func (UserRepository *UserRepository) FindUserByEmail(email string) (models.User, error){
	var foundUser models.User
	result := UserRepository.DB.Where("email = ?", email).First(&foundUser)
	if result.RowsAffected == 0 {
		return models.User{}, errors.New("error trying to creat a new user")
	}
	return foundUser, nil
}