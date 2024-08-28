package authorization

import (
	m "app/pkg/model"
	s "app/pkg/service"
	v "app/pkg/validation"
)

type PermissionManagementAuthorization interface {
	AddPermissionToRole(adminID int, permissionID, roleID int) error
	AddPermissionsToRole(adminID int, permissionIDs []int, roleID int) error
	AddRoleToUser(adminID int, roleID, userID int) error
	GetPermissionsByRoleID(adminID int, roleID int) (*[]m.Permission, error)
	GetPermissonsByUserID(adminID int, userID int) (*[]m.Permission, error)
	GetRoleByUserID(adminID int, userID int) (*m.Role, error)
	RemovePermissionFromRole(adminID int, roleID, permissionID int) error
	RemovePermissionsFromRole(adminID int, roleID int, permissionIDs []int) error
	RemoveRoleFromUser(adminID int, roleID, userID int) error
}

type permissionManagementAuthorization struct {
	auth      s.AuthorizationService
	validator v.PermissionManagementValidator
}

func NewPermissionManagementAuthorization(
	auth s.AuthorizationService,
	validator v.PermissionManagementValidator,
) PermissionManagementAuthorization {
	return &permissionManagementAuthorization{auth, validator}
}

// AddPermissionToRole
func (p *permissionManagementAuthorization) AddPermissionToRole(
	adminID int,
	permissionID int,
	roleID int,
) error {
	panic("unimplemented")
}

// AddPermissionsToRole
func (p *permissionManagementAuthorization) AddPermissionsToRole(
	adminID int,
	permissionIDs []int,
	roleID int,
) error {
	panic("unimplemented")
}

// AddRoleToUser
func (p *permissionManagementAuthorization) AddRoleToUser(
	adminID int,
	roleID int,
	userID int,
) error {
	panic("unimplemented")
}

// GetPermissionsByRoleID
func (p *permissionManagementAuthorization) GetPermissionsByRoleID(
	adminID int,
	roleID int,
) (*[]m.Permission, error) {
	panic("unimplemented")
}

// GetPermissonsByUserID
func (p *permissionManagementAuthorization) GetPermissonsByUserID(
	adminID int,
	userID int,
) (*[]m.Permission, error) {
	panic("unimplemented")
}

// GetRoleByUserID
func (p *permissionManagementAuthorization) GetRoleByUserID(
	adminID int,
	userID int,
) (*m.Role, error) {
	panic("unimplemented")
}

// RemovePermissionFromRole
func (p *permissionManagementAuthorization) RemovePermissionFromRole(
	adminID int,
	roleID int,
	permissionID int,
) error {
	panic("unimplemented")
}

// RemovePermissionsFromRole
func (p *permissionManagementAuthorization) RemovePermissionsFromRole(
	adminID int,
	roleID int,
	permissionIDs []int,
) error {
	panic("unimplemented")
}

// RemoveRoleFromUser
func (p *permissionManagementAuthorization) RemoveRoleFromUser(
	adminID int,
	roleID int,
	userID int,
) error {
	panic("unimplemented")
}
