package model

import "fmt"

type User struct {
	Id        int    `json:"id,omitempty"`
	FirstName string `json:"first_name" valid:"required~first_name is required"`
	LastName  string `json:"last_name" valid:"required~last_name is required"`
	Username  string `json:"username" valid:"required~username is required"`
	Email     string `json:"email" valid:"email~is not a valid email,required~email is required"`
	Password  string `json:"password,omitempty" valid:"required~password is required"`
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
