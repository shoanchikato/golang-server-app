package user

import (
	e "app/pkg/errors"
	"fmt"
	"strings"

	valid "github.com/asaskevich/govalidator"
)

type UserValidator interface {
	Add(user *User) error
	AddAll(users *[]*User) error
	Edit(id int, newUser *User) error
	GetAll() (*[]User, error)
	GetOne(id int) (*User, error)
	Remove(id int) error
}

type userValidator struct {
	encrypt UserEncryption
}

func NewUserValidator(encrypt UserEncryption) UserValidator {
	return &userValidator{encrypt}
}

// Add
func (v *userValidator) Add(user *User) error {
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
func (v *userValidator) AddAll(users *[]*User) error {
	newUsers := *users
	errs := make([]string, len(newUsers))
	for i := 0; i < len(newUsers); i++ {
		_, err := valid.ValidateStruct(newUsers[i])
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
func (v *userValidator) Edit(id int, newUser *User) error {
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
func (v *userValidator) GetAll() (*[]User, error) {
	return v.encrypt.GetAll()
}

// GetOne
func (v *userValidator) GetOne(id int) (*User, error) {
	return v.encrypt.GetOne(id)
}

// Remove
func (v *userValidator) Remove(id int) error {
	return v.encrypt.Remove(id)
}
