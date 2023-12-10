package requests

type LoginRequest struct {
	Email string `json:"email" binding:"required,email"`
	Password string `json:"password" binding:"required,min=8"`
}

type SignupRequest struct {
	LoginRequest
	Pseudo string `json:"pseudo" binding:"required,min=2"`
}

