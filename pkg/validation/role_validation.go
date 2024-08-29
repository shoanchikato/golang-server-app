package validation

import (
	e "app/pkg/errors"
	m "app/pkg/model"
	r "app/pkg/repo"
	"fmt"
	"strings"

	valid "github.com/asaskevich/govalidator"
)

type RoleValidator interface {
	Add(role *m.Role) error
	AddAll(roles *[]*m.Role) error
	Edit(id int, newRole *m.Role) error
	GetAll(lastId, limit int) (*[]m.Role, error)
	GetOne(id int) (*m.Role, error)
	Remove(id int) error
}

type roleValidator struct {
	repo r.RoleRepo
}

func NewRoleValidator(repo r.RoleRepo) RoleValidator {
	return &roleValidator{repo}
}

// Add
func (r *roleValidator) Add(role *m.Role) error {
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
func (r *roleValidator) AddAll(roles *[]*m.Role) error {
	newRoles := *roles
	errs := make([]string, len(newRoles))
	for i := 0; i < len(newRoles); i++ {
		_, err := valid.ValidateStruct(newRoles[i])
		if err != nil {
			errStr := fmt.Sprintf("\n[%d] %s", i, err.Error())
			errs[i] = errStr
		}
	}

	for i := len(newRoles) - 1; i > 0; i-- {
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
func (r *roleValidator) Edit(id int, newRole *m.Role) error {
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
func (r *roleValidator) GetAll(lastId, limit int) (*[]m.Role, error) {
	return r.repo.GetAll(lastId, limit)
}

// GetOne
func (r *roleValidator) GetOne(id int) (*m.Role, error) {
	return r.repo.GetOne(id)
}

// Remove
func (r *roleValidator) Remove(id int) error {
	return r.repo.Remove(id)
}
