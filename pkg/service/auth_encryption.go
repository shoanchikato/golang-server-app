package service

import (
	e "app/pkg/errors"
	m "app/pkg/model"
	r "app/pkg/repo"
	"errors"
)

type AuthEncryption interface {
	Login(credentials *m.Credentials) (bool, error)
	ResetPassword(credentials *m.Credentials, newPassword string) error
}

type authEncryption struct {
	repo    r.AuthRepo
	encrypt EncryptionService
}

func NewAuthEncryption(repo r.AuthRepo, encrypt EncryptionService) AuthEncryption {
	return &authEncryption{repo, encrypt}
}

// Login
func (a *authEncryption) Login(credentials *m.Credentials) (bool, error) {
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
func (a *authEncryption) ResetPassword(credentials *m.Credentials, newPassword string) error {
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
