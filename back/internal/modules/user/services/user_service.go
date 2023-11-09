package services

import (
	"errors"
	"fmt"

	UserModel "github.com/saegus/test-technique-romain-chenard/internal/modules/user/models"
	UserRepository "github.com/saegus/test-technique-romain-chenard/internal/modules/user/repositories"
	UserRequest "github.com/saegus/test-technique-romain-chenard/internal/modules/user/requests"
	UserResponse "github.com/saegus/test-technique-romain-chenard/internal/modules/user/responses"
	"github.com/saegus/test-technique-romain-chenard/pkg/encrypt"
)

type UserService struct {
	userRepository UserRepository.UserRepositoryInterface
}

func New() *UserService{
	return &UserService{
		userRepository: UserRepository.New(),
	}
}

func (userService *UserService) CreateUserSrv (user UserRequest.SignupRequest) (UserModel.User, error){
	
	_, err := userService.userRepository.FindUserByEmail(user.Email)
	if err == nil {
		return UserModel.User{}, errors.New("email already used")
	}
	
	hashedPassword, err := encrypt.HashAndSalt(user.Password)
	if err != nil {
		return UserModel.User{}, err
	}
	
	var newUser UserModel.User

	newUser.Email= user.Email
	newUser.Password= hashedPassword
	newUser.FirstName= user.FirstName
	newUser.LastName= user.LastName

	createdUser, err := userService.userRepository.CreateUser(newUser)
	if err != nil {
		return UserModel.User{}, err
	}
	UserResponse.ToUser(createdUser)

	return createdUser, nil
}

func (userService *UserService) LoginSrv (user UserRequest.LoginRequest) (UserResponse.LoginResponse, error){
	foundUser, err := userService.userRepository.FindUserByEmail(user.Email)
	if err != nil {
		return UserResponse.LoginResponse{}, errors.New("wrong email/password 1")
	}
	
	err = encrypt.ComparePasswords(foundUser.Password, user.Password)
	if err != nil {
		return UserResponse.LoginResponse{}, errors.New("wrong email/password 2")
	}

	token, err := encrypt.Generate(foundUser)
	if err != nil {
		return UserResponse.LoginResponse{}, errors.New("error trying to generate the token")
	}
	fmt.Println("=> token : ", token)

	return UserResponse.LoginResponse{
		ID: foundUser.ID,
		Email: foundUser.Email,
		Token: token,
	}, nil
}