package role

import (
	e "app/pkg/errors"
	"fmt"
	"strings"

	valid "github.com/asaskevich/govalidator"
)

type RoleValidator interface {
	Add(role *Role) error
	AddAll(roles *[]*Role) error
	Edit(id int, newRole *Role) error
	GetAll() (*[]Role, error)
	GetOne(id int) (*Role, error)
	Remove(id int) error
}

type roleValidator struct {
	repo RoleRepo
}

func NewRoleValidator(repo RoleRepo) RoleValidator {
	return &roleValidator{repo}
}

// Add
func (r *roleValidator) Add(role *Role) error {
	_, err := valid.ValidateStruct(role)
	if err != nil {
		return e.NewValidationError(e.ErrAddValidation, err.Error())
	}

	err = r.repo.Add(role)
	if err != nil {
		return err
	}

	return nil
}

// AddAll
func (r *roleValidator) AddAll(roles *[]*Role) error {
	newRoles := *roles
	errs := make([]string, len(newRoles))
	for i := 0; i < len(newRoles); i++ {
		_, err := valid.ValidateStruct(newRoles[i])
		if err != nil {
			errStr := fmt.Sprintf("\n[%d] %s", i, err.Error())
			errs[i] = errStr
		}
	}

	for i := len(newRoles)-1; i > 0; i-- {
		if errs[i] != "" {
			newErrors := strings.Join(errs, "")
			return e.NewValidationError(e.ErrAddAllValidation, newErrors)
		}
	}

	err := r.repo.AddAll(roles)
	if err != nil {
		return err
	}

	return nil
}

// Edit
func (r *roleValidator) Edit(id int, newRole *Role) error {
	_, err := valid.ValidateStruct(newRole)
	if err != nil {
		return e.NewValidationError(e.ErrEditValidation, err.Error())
	}

	err = r.repo.Edit(id, newRole)
	if err != nil {
		return err
	}

	return nil
}

// GetAll
func (r *roleValidator) GetAll() (*[]Role, error) {
	return r.repo.GetAll()
}

// GetOne
func (r *roleValidator) GetOne(id int) (*Role, error) {
	return r.repo.GetOne(id)
}

// Remove
func (r *roleValidator) Remove(id int) error {
	return r.repo.Remove(id)
}
