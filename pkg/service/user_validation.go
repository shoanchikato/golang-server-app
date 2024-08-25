package service

import (
	e "app/pkg/errors"
	m "app/pkg/model"
	"fmt"
	"strings"

	valid "github.com/asaskevich/govalidator"
)

type UserValidator interface {
	Add(user *m.User) error
	AddAll(users *[]*m.User) error
	Edit(id int, newUser *m.User) error
	GetAll() (*[]m.User, error)
	GetOne(id int) (*m.User, error)
	Remove(id int) error
}

type userValidator struct {
	encrypt UserEncryption
}

func NewUserValidator(encrypt UserEncryption) UserValidator {
	return &userValidator{encrypt}
}

// Add
func (v *userValidator) Add(user *m.User) error {
	_, err := valid.ValidateStruct(user)
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
	errs := []string{}
	newUsers := *users
	for i := 0; i < len(newUsers); i++ {
		_, err := valid.ValidateStruct(newUsers[i])
		if err != nil {
			errStr := fmt.Sprintf("\n[%d] %s", i, err.Error())
			errs = append(errs, errStr)
		}
	}

	if len(errs) > 0 {
		newErrors := strings.Join(errs, "")
		return e.NewValidationError(e.ErrAddAllValidation, newErrors)
	}

	err := v.encrypt.AddAll(users)
	if err != nil {
		return err
	}

	return nil
}

// Edit
func (v *userValidator) Edit(id int, newUser *m.User) error {
	_, err := valid.ValidateStruct(newUser)
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
func (v *userValidator) GetAll() (*[]m.User, error) {
	return v.encrypt.GetAll()
}

// GetOne
func (v *userValidator) GetOne(id int) (*m.User, error) {
	return v.encrypt.GetOne(id)
}

// Remove
func (v *userValidator) Remove(id int) error {
	return v.encrypt.Remove(id)
}
