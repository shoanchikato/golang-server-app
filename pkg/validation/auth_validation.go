package validation

import (
	c "app/pkg/constants"
	en "app/pkg/encrypt"
	e "app/pkg/errors"
	m "app/pkg/model"

	valid "github.com/asaskevich/govalidator"
)

type AuthValidator interface {
	Login(credentials *m.Credentials) (userId *int, err error)
	ResetPassword(username, newPassword string) error
}

type authValidator struct {
	service en.AuthEncryption
}

func NewAuthValidator(service en.AuthEncryption) AuthValidator {
	return &authValidator{service}
}

// Login
func (a *authValidator) Login(credentials *m.Credentials) (userId *int, err error) {
	_, err = valid.ValidateStruct(credentials)
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
