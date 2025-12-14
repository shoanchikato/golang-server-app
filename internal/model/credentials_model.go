package model

import valid "github.com/go-ozzo/ozzo-validation/v4"

type Credentials struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

func NewCredentials(username, password string) *Credentials {
	return &Credentials{username, password}
}

func (c *Credentials) Validate() error {
	return valid.ValidateStruct(c,
		valid.Field(&c.Username, valid.Required.Error("username is required")),
		valid.Field(&c.Password, valid.Required.Error("password is required")),
	)
}
