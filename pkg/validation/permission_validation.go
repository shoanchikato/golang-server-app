package validation

import (
	e "app/pkg/errors"
	m "app/pkg/model"
	r "app/pkg/repo"
	s "app/pkg/service"
	"fmt"
	"strings"
)

type PermissionValidator interface {
	Add(permission *m.Permission) error
	AddAll(permissions *[]*m.Permission) error
	Edit(id int, newPermission *m.Permission) error
	GetAll(lastId, limit int) (*[]m.Permission, error)
	GetByEntity(entity string) (*[]m.Permission, error)
	GetOne(id int) (*m.Permission, error)
	Remove(id int) error
}

type permissionValidator struct {
	repo r.PermissionRepo
	service s.ValidationService
}

func NewPermissionValidator(repo r.PermissionRepo, service s.ValidationService) PermissionValidator {
	return &permissionValidator{repo, service}
}

// Add
func (v *permissionValidator) Add(permission *m.Permission) error {
	err := v.service.Validate(permission)
	if err != nil {
		return e.NewValidationError(e.ErrAddValidation, err.Error())
	}

	err = v.repo.Add(permission)
	if err != nil {
		return err
	}

	return nil
}

// AddAll
func (v *permissionValidator) AddAll(permissions *[]*m.Permission) error {
	newPermissions := *permissions
	errs := make([]string, len(newPermissions))
	for i := range newPermissions {
		err := v.service.Validate(newPermissions[i])
		if err != nil {
			errStr := fmt.Sprintf("\n[%d] %s", i, err.Error())
			errs[i] = errStr
		}
	}

	for i := len(newPermissions) - 1; i > 0; i-- {
		if errs[i] != "" {
			newErrors := strings.Join(errs, "")
			return e.NewValidationError(e.ErrAddAllValidation, newErrors)
		}
	}

	err := v.repo.AddAll(permissions)
	if err != nil {
		return err
	}

	return nil
}

// Edit
func (v *permissionValidator) Edit(id int, newPermission *m.Permission) error {
	err := v.service.Validate(newPermission)
	if err != nil {
		return e.NewValidationError(e.ErrEditValidation, err.Error())
	}

	err = v.repo.Edit(id, newPermission)
	if err != nil {
		return err
	}

	return nil
}

// GetAll
func (v *permissionValidator) GetAll(lastId, limit int) (*[]m.Permission, error) {
	return v.repo.GetAll(lastId, limit)
}

// GetByEntity
func (v *permissionValidator) GetByEntity(entity string) (*[]m.Permission, error) {
	return v.repo.GetByEntity(entity)
}

// GetOne
func (v *permissionValidator) GetOne(id int) (*m.Permission, error) {
	return v.repo.GetOne(id)
}

// Remove
func (v *permissionValidator) Remove(id int) error {
	return v.repo.Remove(id)
}
