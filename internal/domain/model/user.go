package model

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type User struct {
	UserID    primitive.ObjectID `json:"user_id" bson:"_id,omitempty" redis:"user_id" validate:"omitempty"`
	FirstName string             `json:"first_name" bson:"first_name,omitempty" redis:"first_name" validate:"required,lte=30"`
	LastName  string             `json:"last_name" bson:"last_name,omitempty" redis:"last_name" validate:"required,lte=30"`
	Email     string             `json:"email,omitempty" bson:"email,omitempty" redis:"email" validate:"omitempty,lte=60,email"`
	Password  string             `json:"password,omitempty" bson:"password,omitempty" redis:"password" validate:"omitempty,required,gte=6"`
	Role      *string            `json:"role,omitempty" bson:"role,omitempty" redis:"role" validate:"omitempty,lte=10"`
	CreatedAt time.Time          `json:"created_at,omitempty" bson:"created_at,omitempty" redis:"created_at"`
	UpdatedAt time.Time          `json:"updated_at,omitempty" bson:"updated_at,omitempty" redis:"updated_at"`
} // @name User

// GetUserID returns the user id
func (u *User) GetUserID() primitive.ObjectID {
	return u.UserID
}

// GetUserIDString returns the user id as a string
func (u *User) GetUserIDString() string {
	if u.UserID.IsZero() {
		return ""
	}

	return u.GetUserID().Hex()
}

// GetRole returns the role of the user
func (u *User) GetRole() string {
	if u.Role != nil {
		return *u.Role
	}

	return RoleUser
}

// IsAdmin returns true if the user is an admin
func (u *User) IsAdmin() bool {
	return u.GetRole() == RoleAdmin
}

// IsUser returns true if the user is a user
func (u *User) IsUser() bool {
	return u.GetRole() == RoleUser
}
