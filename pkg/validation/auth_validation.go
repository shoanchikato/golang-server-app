package validation

import (
	c "app/pkg/constants"
	en "app/pkg/encrypt"
	e "app/pkg/errors"
	m "app/pkg/model"
	s "app/pkg/service"
)

type AuthValidator interface {
	Login(credentials *m.Credentials) (userId *int, err error)
	ResetPassword(username, newPassword string) error
}

type authValidator struct {
	service    en.AuthEncryption
	validation s.ValidationService
}

func NewAuthValidator(service en.AuthEncryption, validation s.ValidationService) AuthValidator {
	return &authValidator{service, validation}
}

// Login
func (a *authValidator) Login(credentials *m.Credentials) (userId *int, err error) {
	err = a.validation.Validate(credentials)
	if err != nil {
		return nil, e.NewValidationError(e.ErrLoginValidation, err.Error())
	}

	return a.service.Login(*credentials)
}

// ResetPassword
func (a *authValidator) ResetPassword(username, newPassword string) error {
	if len(username) < c.USERNAME_LENGTH {
		return e.NewValidationError(e.ErrResetPasswordValidation, e.ErrUsernameLength.Error())
	}

	if len(newPassword) < c.PASSWORD_LENGTH {
		return e.NewValidationError(e.ErrResetPasswordValidation, e.ErrPasswordLength.Error())
	}
	return a.service.ResetPassword(username, newPassword)
}
