package model

import "fmt"

type User struct {
	Id        int    `json:"id,omitempty"`
	FirstName string `json:"first_name" validate:"required~first_name is required"`
	LastName  string `json:"last_name" validate:"required~last_name is required"`
	Username  string `json:"username" validate:"required~username is required"`
	Email     string `json:"email" validate:"email~is not a valid email,required~email is required"`
	Password  string `json:"password,omitempty" validate:"required~password is required"`
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
