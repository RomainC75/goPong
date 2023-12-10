package services

import (
	"errors"
	"fmt"

	UserRequest "github.com/saegus/test-technique-romain-chenard/api/dto/requests"
	Responses "github.com/saegus/test-technique-romain-chenard/api/dto/responses"
	UserRepository "github.com/saegus/test-technique-romain-chenard/api/repositories"
	Models "github.com/saegus/test-technique-romain-chenard/data/models"
	"github.com/saegus/test-technique-romain-chenard/utils/encrypt"
)

type UserService struct {
	userRepository UserRepository.UserRepositoryInterface
}

func NewUserSrv() *UserService{
	return &UserService{
		userRepository: UserRepository.NewUserRepo(),
	}
}

func (userService *UserService) CreateUserSrv (user UserRequest.SignupRequest) (Models.User, error){
	
	_, err := userService.userRepository.FindUserByEmail(user.Email)
	if err == nil {
		return Models.User{}, errors.New("email already used")
	}
	
	hashedPassword, err := encrypt.HashAndSalt(user.Password)
	if err != nil {
		return Models.User{}, err
	}
	
	var newUser Models.User

	newUser.Email= user.Email
	newUser.Password= hashedPassword
	newUser.Pseudo= user.Pseudo
	

	createdUser, err := userService.userRepository.CreateUser(newUser)
	if err != nil {
		return Models.User{}, err
	}
	Responses.ToUser(createdUser)

	return createdUser, nil
}

func (userService *UserService) LoginSrv (user UserRequest.LoginRequest) (Responses.LoginResponse, error){
	foundUser, err := userService.userRepository.FindUserByEmail(user.Email)
	if err != nil {
		return Responses.LoginResponse{}, errors.New("wrong email/password 1")
	}
	
	err = encrypt.ComparePasswords(foundUser.Password, user.Password)
	if err != nil {
		return Responses.LoginResponse{}, errors.New("wrong email/password 2")
	}

	token, err := encrypt.Generate(foundUser)
	if err != nil {
		return Responses.LoginResponse{}, errors.New("error trying to generate the token")
	}
	fmt.Println("=> token : ", token)

	return Responses.LoginResponse{
		ID: foundUser.ID,
		Email: foundUser.Email,
		Token: token,
	}, nil
}