package authorization

import (
	e "app/pkg/errors"
	m "app/pkg/model"
	p "app/pkg/permission"
	s "app/pkg/service"
	v "app/pkg/validation"
	"errors"
)

type AuthAuthorization interface {
	Login(userId int, credentials *m.Credentials) (bool, error)
	ResetPassword(userId int, credentials *m.Credentials, newPassword string) error
}

type authAuthorization struct {
	auth      s.AuthorizationService
	validator v.AuthValidator
}

func NewAuthAuthorization(
	auth s.AuthorizationService,
	validator v.AuthValidator,
) AuthAuthorization {
	return &authAuthorization{auth, validator}
}

// Login
func (a authAuthorization) Login(userId int, credentials *m.Credentials) (bool, error) {
	err := a.auth.CheckForAuthorization(userId, p.AuthLogin.Name)
	if err != nil {
		return false, errors.Join(e.ErrAuthDomain, e.ErrOnLogin, err)
	}

	return a.validator.Login(credentials)
}

// ResetPassword
func (a authAuthorization) ResetPassword(userId int, credentials *m.Credentials, newPassword string) error {
	err := a.auth.CheckForAuthorization(userId, p.AuthResetPassword.Name)
	if err != nil {
		return errors.Join(e.ErrAuthDomain, e.ErrOnResetPassword, err)
	}

	return a.validator.ResetPassword(credentials, newPassword)
}
