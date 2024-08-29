package encrypt

import (
	e "app/pkg/errors"
	m "app/pkg/model"
	r "app/pkg/repo"
	s "app/pkg/service"
	"errors"
)

type AuthEncryption interface {
	Login(credentials m.Credentials) (userId *int, err error)
	ResetPassword(credentials m.Credentials, newPassword string) error
}

type authEncryption struct {
	repo    r.AuthRepo
	encrypt s.EncryptionService
}

func NewAuthEncryption(repo r.AuthRepo, encrypt s.EncryptionService) AuthEncryption {
	return &authEncryption{repo, encrypt}
}

// Login
func (a *authEncryption) Login(credentials m.Credentials) (userId *int, err error) {
	authDetails, err := a.repo.GetByUsername(credentials.Username)
	if err != nil {
		return nil, err
	}

	isMatch, err := a.encrypt.CheckPassword(&authDetails.Password, &credentials.Password)
	if err != nil {
		return nil, errors.Join(e.ErrIncorrectCredentials, err)
	}

	if !isMatch {
		return nil, e.ErrIncorrectCredentials
	}

	return &authDetails.UserId, nil
}

// ResetPassword
func (a *authEncryption) ResetPassword(credentials m.Credentials, newPassword string) error {
	userId, err := a.Login(credentials)
	if err != nil {
		return err
	}

	if userId == nil {
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
