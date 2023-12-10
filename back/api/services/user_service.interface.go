package services

import (
	UserRequest "github.com/saegus/test-technique-romain-chenard/api/dto/requests"
	UserResponse "github.com/saegus/test-technique-romain-chenard/api/dto/responses"
	UserModel "github.com/saegus/test-technique-romain-chenard/data/models"
)

type UserServiceInterface interface {
	CreateUserSrv (user UserRequest.SignupRequest) (UserModel.User, error)
	LoginSrv (user UserRequest.LoginRequest) (UserResponse.LoginResponse, error)
}
