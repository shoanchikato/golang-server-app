package validation

import (
	e "app/pkg/errors"
	m "app/pkg/model"
	r "app/pkg/repo"
	s "app/pkg/service"
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
	repo    r.RoleRepo
	service s.ValidationService
}

func NewRoleValidator(repo r.RoleRepo, service s.ValidationService) RoleValidator {
	return &roleValidator{repo, service}
}

// Add
func (r *roleValidator) Add(role *m.Role) error {
	err := r.service.Validate(role)
	if err != nil {
		return e.NewValidationError(e.ErrAddValidation, err)
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
	errs := make([]error, 0, len(newRoles))
	for i := range newRoles {
		err := r.service.Validate(newRoles[i])
		if err != nil {
			errs = append(errs, err)
		}
	}

	if len(errs) > 0 {
		return e.NewValidationError(e.ErrAddAllValidation, errs...)
	}

	err := r.repo.AddAll(roles)
	if err != nil {
		return err
	}

	return nil
}

// Edit
func (r *roleValidator) Edit(id int, newRole *m.Role) error {
	err := r.service.Validate(newRole)
	if err != nil {
		return e.NewValidationError(e.ErrEditValidation, err)
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
