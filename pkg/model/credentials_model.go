package model

type Credentials struct {
	Username string `json:"username" validate:"required"`
	Password string `json:"password" validate:"required"`
}

func NewCredentials(username, password string) *Credentials {
	return &Credentials{username, password}
}
