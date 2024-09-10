package validation

import (
	en "app/pkg/encrypt"
	e "app/pkg/errors"
	m "app/pkg/model"
	s "app/pkg/service"
	"fmt"
	"strings"
)

type UserValidator interface {
	Add(user *m.User) error
	AddAll(users *[]*m.User) error
	Edit(id int, newUser *m.EditUser) error
	GetAll(lastId, limit int) (*[]m.User, error)
	GetOne(id int) (*m.User, error)
	Remove(id int) error
}

type userValidator struct {
	encrypt en.UserEncryption
	service s.ValidationService
}

func NewUserValidator(encrypt en.UserEncryption, service s.ValidationService) UserValidator {
	return &userValidator{encrypt, service}
}

// Add
func (v *userValidator) Add(user *m.User) error {
	err := v.service.Validate(user)
	if err != nil {
		return e.NewValidationError(e.ErrAddValidation, err.Error())
	}

	err = v.encrypt.Add(user)
	if err != nil {
		return err
	}

	return nil
}

// AddAll
func (v *userValidator) AddAll(users *[]*m.User) error {
	newUsers := *users
	errs := make([]string, len(newUsers))
	for i := range newUsers {
		err := v.service.Validate(newUsers[i])
		if err != nil {
			errStr := fmt.Sprintf("\n[%d] %s", i, err.Error())
			errs[i] = errStr
		}
	}

	for i := len(newUsers) - 1; i > 0; i-- {
		if errs[i] != "" {
			newErrors := strings.Join(errs, "")
			return e.NewValidationError(e.ErrAddAllValidation, newErrors)
		}
	}

	err := v.encrypt.AddAll(users)
	if err != nil {
		return err
	}

	return nil
}

// Edit
func (v *userValidator) Edit(id int, newUser *m.EditUser) error {
	err := v.service.Validate(newUser)
	if err != nil {
		return e.NewValidationError(e.ErrEditValidation, err.Error())
	}

	err = v.encrypt.Edit(id, newUser)
	if err != nil {
		return err
	}

	return nil
}

// GetAll
func (v *userValidator) GetAll(lastId, limit int) (*[]m.User, error) {
	return v.encrypt.GetAll(lastId, limit)
}

// GetOne
func (v *userValidator) GetOne(id int) (*m.User, error) {
	return v.encrypt.GetOne(id)
}

// Remove
func (v *userValidator) Remove(id int) error {
	return v.encrypt.Remove(id)
}
