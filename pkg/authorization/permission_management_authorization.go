package authorization

import (
	m "app/pkg/model"
	s "app/pkg/service"
	v "app/pkg/validation"
)

type PermissionManagementAuthorization interface {
	AddPermissionToRole(adminId int, permissionId, roleId int) error
	AddPermissionsToRole(adminId int, permissionIds []int, roleId int) error
	AddRoleToUser(adminId int, roleId, userId int) error
	GetPermissionsByRoleId(adminId int, roleId int) (*[]m.Permission, error)
	GetPermissonsByUserId(adminId int, userId int) (*[]m.Permission, error)
	GetRoleByUserId(adminId int, userId int) (*m.Role, error)
	RemovePermissionFromRole(adminId int, roleId, permissionId int) error
	RemovePermissionsFromRole(adminId int, roleId int, permissionIds []int) error
	RemoveRoleFromUser(adminId int, roleId, userId int) error
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
	adminId int,
	permissionId int,
	roleId int,
) error {
	panic("unimplemented")
}

// AddPermissionsToRole
func (p *permissionManagementAuthorization) AddPermissionsToRole(
	adminId int,
	permissionIds []int,
	roleId int,
) error {
	panic("unimplemented")
}

// AddRoleToUser
func (p *permissionManagementAuthorization) AddRoleToUser(
	adminId int,
	roleId int,
	userId int,
) error {
	panic("unimplemented")
}

// GetPermissionsByRoleId
func (p *permissionManagementAuthorization) GetPermissionsByRoleId(
	adminId int,
	roleId int,
) (*[]m.Permission, error) {
	panic("unimplemented")
}

// GetPermissonsByUserId
func (p *permissionManagementAuthorization) GetPermissonsByUserId(
	adminId int,
	userId int,
) (*[]m.Permission, error) {
	panic("unimplemented")
}

// GetRoleByUserId
func (p *permissionManagementAuthorization) GetRoleByUserId(
	adminId int,
	userId int,
) (*m.Role, error) {
	panic("unimplemented")
}

// RemovePermissionFromRole
func (p *permissionManagementAuthorization) RemovePermissionFromRole(
	adminId int,
	roleId int,
	permissionId int,
) error {
	panic("unimplemented")
}

// RemovePermissionsFromRole
func (p *permissionManagementAuthorization) RemovePermissionsFromRole(
	adminId int,
	roleId int,
	permissionIds []int,
) error {
	panic("unimplemented")
}

// RemoveRoleFromUser
func (p *permissionManagementAuthorization) RemoveRoleFromUser(
	adminId int,
	roleId int,
	userId int,
) error {
	panic("unimplemented")
}
