package encrypt

import (
	r "app/internal/repo"
	s "app/internal/service"
)

type Encryptions struct {
	Auth AuthEncryption
	User UserEncryption
}

func EncryptDi(encrypt s.EncryptionService, repos *r.Repos) *Encryptions {
	auth := NewAuthEncryption(repos.Auth, encrypt)
	user := NewUserEncryption(repos.User, encrypt)

	return &Encryptions{auth, user}
}
