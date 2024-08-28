package validation

import (
	en "app/pkg/encrypt"
	e "app/pkg/errors"
	m "app/pkg/model"

	valid "github.com/asaskevich/govalidator"
)

type AuthValidator interface {
	Login(credentials *m.Credentials) (bool, error)
	ResetPassword(credentials *m.Credentials, newPassword string) error
}

type authValidator struct {
	service en.AuthEncryption
}

func NewAuthValidator(service en.AuthEncryption) AuthValidator {
	return &authValidator{service}
}

// Login implements AuthValidator.
func (a *authValidator) Login(credentials *m.Credentials) (bool, error) {
	_, err := valid.ValidateStruct(credentials)
	if err != nil {
		return false, e.NewValidationError(e.ErrLoginValidation, err.Error())
	}

	return a.service.Login(*credentials)
}

// ResetPassword implements AuthValidator.
func (a *authValidator) ResetPassword(credentials *m.Credentials, newPassword string) error {
	_, err := valid.ValidateStruct(credentials)
	if err != nil {
		return e.NewValidationError(e.ErrResetPasswordValidation, err.Error())
	}

	return a.service.ResetPassword(*credentials, newPassword)
}
