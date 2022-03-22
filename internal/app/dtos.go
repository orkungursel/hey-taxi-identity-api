package app

type LoginRequest struct {
	Email    string `json:"email" validate:"required,email"`
	Password string `json:"password" validate:"required"`
} // @name LoginRequest

// LoginResponse is the response of LoginRequest
type LoginResponse struct {
	Token        string `json:"token" validate:"required"`
	RefreshToken string `json:"refresh_token" validate:"required"`
	ExpiresIn    int    `json:"expires_in" validate:"required"`
} // @name LoginResponse

type RegisterRequest struct {
	Email    string `json:"email" validate:"required,email"`
	Password string `json:"password" validate:"required"`
} // @name RegisterResponse
