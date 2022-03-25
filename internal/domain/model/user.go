package model

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type User struct {
	Id        primitive.ObjectID `json:"id" bson:"_id,omitempty" redis:"id" validate:"omitempty"`
	FirstName string             `json:"first_name" bson:"first_name,omitempty" redis:"first_name" validate:"required,lte=30"`
	LastName  string             `json:"last_name" bson:"last_name,omitempty" redis:"last_name" validate:"required,lte=30"`
	Email     string             `json:"email,omitempty" bson:"email,omitempty" redis:"email" validate:"omitempty,lte=100,email"`
	Password  string             `json:"-,omitempty" bson:"password,omitempty" redis:"password" validate:"omitempty,required,gte=6,lte=60"`
	Role      string             `json:"role,omitempty" bson:"role,omitempty" redis:"role" validate:"omitempty,lte=10"`
	Avatar    string             `json:"avatar,omitempty" bson:"avatar,omitempty" redis:"avatar" validate:"omitempty"`
	CreatedAt time.Time          `json:"created_at,omitempty" bson:"created_at,omitempty" redis:"created_at"`
	UpdatedAt time.Time          `json:"updated_at,omitempty" bson:"updated_at,omitempty" redis:"updated_at"`
} // @name User

// GetId returns the user id
func (u *User) GetId() primitive.ObjectID {
	return u.Id
}

// GetIdString returns the user id as a string
func (u *User) GetIdString() string {
	if u.Id.IsZero() {
		return ""
	}

	return u.GetId().Hex()
}

// GetRole returns the role of the user
func (u *User) GetRole() string {
	if u.Role != "" {
		return u.Role
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
