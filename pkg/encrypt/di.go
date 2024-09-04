package encrypt

import (
	r "app/pkg/repo"
	s "app/pkg/service"
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
