package app

import "github.com/orkungursel/hey-taxi-identity-api/internal/domain/model"

type HTTPError struct {
	Code     int         `json:"-"`
	Message  interface{} `json:"message"`
	Internal error       `json:"-"` // Stores the error returned by an external dependency
} // name: "HTTPError"

// SuccessAuthResponse is the response of LoginRequest
type SuccessAuthResponse struct {
	UserDto               UserResponse `json:"user"`
	AccessToken           string       `json:"access_token"`
	AccessTokenExpiresIn  int          `json:"access_token_expires_in"`
	RefreshToken          string       `json:"refresh_token"`
	RefreshTokenExpiresIn int          `json:"refresh_token_expires_in"`
} // @name SuccessAuthResponse

type UserResponse struct {
	Id        string `json:"id"`
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
	Email     string `json:"email"`
	Role      string `json:"role"`
	Avatar    string `json:"avatar"`
} // @name UserResponse

func (r *UserResponse) fromUser(u *model.User) {
	r.Id = u.GetIdString()
	r.FirstName = u.FirstName
	r.LastName = u.LastName
	r.Email = u.Email
	r.Avatar = u.Avatar
	r.Role = u.GetRole()
}

func UserResponseFromUser(u *model.User) *UserResponse {
	r := &UserResponse{}
	r.fromUser(u)
	return r
}
