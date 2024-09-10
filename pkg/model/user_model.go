package model

import (
	"fmt"
	valid "github.com/go-ozzo/ozzo-validation/v4"
)

type User struct {
	Id        int    `json:"id,omitempty"`
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
	Username  string `json:"username"`
	Email     string `json:"email"`
	Password  string `json:"password,omitempty"`
}

func NewUser(firstName, lastName, username, email, password string) *User {
	return &User{0, firstName, lastName, username, email, password}
}

func (u User) String() string {
	return fmt.Sprintf(
		"{%d, %s, %s, %s, %s}",
		u.Id, u.FirstName, u.LastName, u.Username, u.Email,
	)
}

func (u *User) Validate() error {
	return valid.ValidateStruct(u,
		valid.Field(&u.FirstName, valid.Required.Error("first_name is required")),
		valid.Field(&u.LastName, valid.Required.Error("last_name is required")),
		valid.Field(&u.Username, valid.Required.Error("username is required")),
		valid.Field(&u.Email, valid.Required.Error("email is required"),),
		valid.Field(&u.Password, valid.Required.Error("password is required")),
	)
}

type EditUser struct {
	Id        int    `json:"id,omitempty"`
	FirstName string `json:"first_name" validate:"required~first_name is required"`
	LastName  string `json:"last_name" validate:"required~last_name is required"`
	Username  string `json:"username" validate:"required~username is required"`
	Email     string `json:"email" validate:"email~is not a valid email,required~email is required"`
}

func (e EditUser) ToUser() *User {
	return &User{e.Id, e.FirstName, e.LastName, e.Username, e.Email, ""}
}

func (e *EditUser) Validate() error {
	return valid.ValidateStruct(&e,
		valid.Field(&e.FirstName, valid.Required.Error("first_name is required")),
		valid.Field(&e.LastName, valid.Required.Error("last_name is required")),
		valid.Field(&e.Username, valid.Required.Error("username is required")),
		valid.Field(&e.Email, valid.Required.Error("email is required"),),
	)
}
