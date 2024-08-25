package auth

type Credentials struct {
	Username string `json:"username" valid:"required~username is required"`
	Password string `json:"password" valid:"required~password is required"`
}

func NewCredentials(username, password string) *Credentials {
	return &Credentials{username, password}
}
