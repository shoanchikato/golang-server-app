package permission

import (
	e "app/pkg/errors"
	rr "app/pkg/role"
	"fmt"
	"strings"

	valid "github.com/asaskevich/govalidator"
)

type PermissionManagementValidator interface {
	AddPermissionToRole(permission *Permission, roleID int) error
	AddPermissionsToRole(permissions *[]*Permission, roleID int) error
	AddRoleToUser(roleID, userID int) error
	GetPermissionsByRoleID(roleID int) (*[]Permission, error)
	GetPermissonsByUserID(userID int) (*[]Permission, error)
	GetRoleByUserID(userID int) (*rr.Role, error)
	RemovePermissionFromRole(roleID, permissionID int) error
	RemovePermissionsFromRole(roleID int, permissionIDs []int) error
	RemoveRoleFromUser(roleID, userID int) error
}

type pMValidator struct {
	repo PermissionManagementRepo
}

func NewPermissionManagementValidator(repo PermissionManagementRepo) PermissionManagementValidator {
	return &pMValidator{repo}
}

// AddPermissionToRole
func (p *pMValidator) AddPermissionToRole(permission *Permission, roleID int) error {
	_, err := valid.ValidateStruct(permission)
	if err != nil {
		return e.NewValidationError(e.ErrAddValidation, err.Error())
	}

	err = p.repo.AddPermissionToRole(permission, roleID)
	if err != nil {
		return err
	}

	return nil
}

// AddPermissionsToRole
func (p *pMValidator) AddPermissionsToRole(permissions *[]*Permission, roleID int) error {
	errs := []string{}
	newPermissions := *permissions
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

	err := p.repo.AddPermissionsToRole(permissions, roleID)
	if err != nil {
		return err
	}

	return nil
}

// AddRoleToUser
func (p *pMValidator) AddRoleToUser(roleID int, userID int) error {
	return p.repo.AddRoleToUser(roleID, userID)
}

// GetPermissionsByRoleID
func (p *pMValidator) GetPermissionsByRoleID(roleID int) (*[]Permission, error) {
	return p.repo.GetPermissionsByRoleID(roleID)
}

// GetPermissonsByUserID
func (p *pMValidator) GetPermissonsByUserID(userID int) (*[]Permission, error) {
	return p.repo.GetPermissonsByUserID(userID)
}

// GetRoleByUserID
func (p *pMValidator) GetRoleByUserID(userID int) (*rr.Role, error) {
	return p.repo.GetRoleByUserID(userID)
}

// RemovePermissionFromRole
func (p *pMValidator) RemovePermissionFromRole(roleID int, permissionID int) error {
	return p.repo.RemovePermissionFromRole(roleID, permissionID)
}

// RemovePermissionsFromRole
func (p *pMValidator) RemovePermissionsFromRole(roleID int, permissionIDs []int) error {
	return p.repo.RemovePermissionsFromRole(roleID, permissionIDs)
}

// RemoveRoleFromUser
func (p *pMValidator) RemoveRoleFromUser(roleID int, userID int) error {
	return p.repo.RemoveRoleFromUser(roleID, userID)
}
