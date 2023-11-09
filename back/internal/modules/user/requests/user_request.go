package requests

type LoginRequest struct {
	Email string `json:"email" binding:"required,email"`
	Password string `json:"password" binding:"required,min=8"`
}

type SignupRequest struct {
	LoginRequest
	FirstName string `json:"firstname" binding:"required,min=2"`
	LastName string `json:"lastname" binding:"required,min=2"`
}

