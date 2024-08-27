package permission

import (
	e "app/pkg/errors"
	"fmt"
	"strings"

	valid "github.com/asaskevich/govalidator"
)

type PermissionValidator interface {
	Add(permission *Permission) error
	AddAll(permissions *[]*Permission) error
	Edit(id int, newPermission *Permission) error
	GetAll() (*[]Permission, error)
	GetOne(id int) (*Permission, error)
	Remove(id int) error
}

type permissionValidator struct {
	repo PermissionRepo
}

func NewPermissionValidator(repo PermissionRepo) PermissionValidator {
	return &permissionValidator{repo}
}

// Add
func (v *permissionValidator) Add(permission *Permission) error {
	_, err := valid.ValidateStruct(permission)
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
func (v *permissionValidator) AddAll(permissions *[]*Permission) error {
	newPermissions := *permissions
	errs := make([]string, len(newPermissions))
	for i := 0; i < len(newPermissions); i++ {
		_, err := valid.ValidateStruct(newPermissions[i])
		if err != nil {
			errStr := fmt.Sprintf("\n[%d] %s", i, err.Error())
			errs = append(errs, errStr)
		}
	}

	if len(errs) > 0 {
		newErrors := strings.Join(errs, "")
		return e.NewValidationError(e.ErrAddAllValidation, newErrors)
	}

	err := v.repo.AddAll(permissions)
	if err != nil {
		return err
	}

	return nil
}

// Edit
func (v *permissionValidator) Edit(id int, newPermission *Permission) error {
	_, err := valid.ValidateStruct(newPermission)
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
func (v *permissionValidator) GetAll() (*[]Permission, error) {
	return v.repo.GetAll()
}

// GetOne
func (v *permissionValidator) GetOne(id int) (*Permission, error) {
	return v.repo.GetOne(id)
}

// Remove
func (v *permissionValidator) Remove(id int) error {
	return v.repo.Remove(id)
}
