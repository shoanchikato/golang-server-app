package auth

import (
	e "app/pkg/errors"
	s "app/pkg/service"
	"errors"
)

type AuthEncryption interface {
	Login(credentials Credentials) (bool, error)
	ResetPassword(credentials Credentials, newPassword string) error
}

type authEncryption struct {
	repo    AuthRepo
	encrypt s.EncryptionService
}

func NewAuthEncryption(repo AuthRepo, encrypt s.EncryptionService) AuthEncryption {
	return &authEncryption{repo, encrypt}
}

// Login
func (a *authEncryption) Login(credentials Credentials) (bool, error) {
	authDetails, err := a.repo.GetByUsername(credentials.Username)
	if err != nil {
		return false, err
	}

	isMatch, err := a.encrypt.CheckPassword(&authDetails.Password, &credentials.Password)
	if err != nil {
		return false, errors.Join(e.ErrIncorrectCredentials, err)
	}

	if !isMatch {
		return false, e.ErrIncorrectCredentials
	}

	return true, nil
}

// ResetPassword
func (a *authEncryption) ResetPassword(credentials Credentials, newPassword string) error {
	isMatch, err := a.Login(credentials)
	if err != nil {
		return err
	}

	if !isMatch {
		return e.ErrIncorrectCredentials
	}

	err = a.encrypt.HashPassword(&newPassword)
	if err != nil {
		return err
	}

	err = a.repo.ResetPassword(credentials.Username, newPassword)
	if err != nil {
		return err
	}

	return nil
}
