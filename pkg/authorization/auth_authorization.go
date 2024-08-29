package authorization

import (
	m "app/pkg/model"
	s "app/pkg/service"
	v "app/pkg/validation"
)

type AuthAuthorization interface {
	Login(credentials *m.Credentials) (userId *int, err error)
	ResetPassword(username, newPassword string) error
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
func (a authAuthorization) Login(credentials *m.Credentials) (userId *int, err error) {
	return a.validator.Login(credentials)
}

// ResetPassword
func (a authAuthorization) ResetPassword(username, newPassword string) error {
	return a.validator.ResetPassword(username, newPassword)
}
