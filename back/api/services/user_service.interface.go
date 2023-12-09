package services

import (
	UserModel "github.com/saegus/test-technique-romain-chenard/internal/modules/user/models"
	UserRequest "github.com/saegus/test-technique-romain-chenard/internal/modules/user/requests"
	UserResponse "github.com/saegus/test-technique-romain-chenard/internal/modules/user/responses"
)

type UserServiceInterface interface {
	CreateUserSrv (user UserRequest.SignupRequest) (UserModel.User, error)
	LoginSrv (user UserRequest.LoginRequest) (UserResponse.LoginResponse, error)
}
