package app

type LoginRequest struct {
	Email    string `json:"email" validate:"required,email,lte=100"`
	Password string `json:"password" validate:"required,gte=6,lte=60"`
} // @name LoginRequest

type RegisterRequest struct {
	Email    string `json:"email" validate:"required,email,lte=100"`
	Password string `json:"password" validate:"required,gte=6,lte=60"`
} // @name RegisterResponse

type RefreshTokenRequest struct {
	Token string `json:"token" validate:"required"`
} // @name RefreshTokenRequest
