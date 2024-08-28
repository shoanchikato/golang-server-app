package authorization

import (
	e "app/pkg/errors"
	m "app/pkg/model"
	p "app/pkg/permission"
	s "app/pkg/service"
	v "app/pkg/validation"
	"errors"
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
func (pe *permissionManagementAuthorization) AddPermissionToRole(
	adminId int,
	permissionId int,
	roleId int,
) error {
	err := pe.auth.CheckForAuthorization(adminId, p.PermissionManagementAddPermissionToRole.Name)
	if err != nil {
		return errors.Join(e.ErrAuthorDomain, e.ErrOnAddAll, err)
	}

	return pe.validator.AddPermissionToRole(permissionId, roleId)
}

// AddPermissionsToRole
func (pe *permissionManagementAuthorization) AddPermissionsToRole(
	adminId int,
	permissionIds []int,
	roleId int,
) error {
	err := pe.auth.CheckForAuthorization(adminId, p.PermissionManagementAddPermissionsToRole.Name)
	if err != nil {
		return errors.Join(e.ErrAuthorDomain, e.ErrOnAddPermissionToRole, err)
	}

	return pe.validator.AddPermissionsToRole(permissionIds, roleId)
}

// AddRoleToUser
func (pe *permissionManagementAuthorization) AddRoleToUser(
	adminId int,
	roleId int,
	userId int,
) error {
	err := pe.auth.CheckForAuthorization(adminId, p.PermissionManagementAddPermissionsToRole.Name)
	if err != nil {
		return errors.Join(e.ErrAuthorDomain, e.ErrOnAddRoleToUser, err)
	}

	return pe.validator.AddRoleToUser(roleId, userId)
}

// GetPermissionsByRoleId
func (pe *permissionManagementAuthorization) GetPermissionsByRoleId(
	adminId int,
	roleId int,
) (*[]m.Permission, error) {
	err := pe.auth.CheckForAuthorization(adminId, p.PermissionManagementGetPermissionsByRoleId.Name)
	if err != nil {
		return nil, errors.Join(e.ErrAuthorDomain, e.ErrOnGetPermissionsByRoleId, err)
	}

	return pe.validator.GetPermissionsByRoleId(roleId)
}

// GetPermissonsByUserId
func (pe *permissionManagementAuthorization) GetPermissonsByUserId(
	adminId int,
	userId int,
) (*[]m.Permission, error) {
	err := pe.auth.CheckForAuthorization(adminId, p.PermissionManagementGetPermissonsByUserId.Name)
	if err != nil {
		return nil, errors.Join(e.ErrAuthorDomain, e.ErrOnGetPermissonsByUserId, err)
	}

	return pe.validator.GetPermissonsByUserId(userId)
}

// GetRoleByUserId
func (pe *permissionManagementAuthorization) GetRoleByUserId(
	adminId int,
	userId int,
) (*m.Role, error) {
	err := pe.auth.CheckForAuthorization(adminId, p.PermissionManagementGetRoleByUserId.Name)
	if err != nil {
		return nil, errors.Join(e.ErrAuthorDomain, e.ErrOnGetRoleByUserId, err)
	}

	return pe.validator.GetRoleByUserId(userId)
}

// RemovePermissionFromRole
func (pe *permissionManagementAuthorization) RemovePermissionFromRole(
	adminId int,
	roleId int,
	permissionId int,
) error {
	err := pe.auth.CheckForAuthorization(adminId, p.PermissionManagementRemovePermissionFromRole.Name)
	if err != nil {
		return errors.Join(e.ErrAuthorDomain, e.ErrOnRemovePermissionFromRole, err)
	}

	return pe.validator.RemovePermissionFromRole(roleId, permissionId)
}

// RemovePermissionsFromRole
func (pe *permissionManagementAuthorization) RemovePermissionsFromRole(
	adminId int,
	roleId int,
	permissionIds []int,
) error {
	err := pe.auth.CheckForAuthorization(adminId, p.PermissionManagementRemovePermissionsFromRole.Name)
	if err != nil {
		return errors.Join(e.ErrAuthorDomain, e.ErrOnRemovePermissionsFromRole, err)
	}

	return pe.validator.RemovePermissionsFromRole(roleId, permissionIds)
}

// RemoveRoleFromUser
func (pe *permissionManagementAuthorization) RemoveRoleFromUser(
	adminId int,
	roleId int,
	userId int,
) error {
	err := pe.auth.CheckForAuthorization(adminId, p.PermissionManagementRemoveRoleFromUser.Name)
	if err != nil {
		return errors.Join(e.ErrAuthorDomain, e.ErrOnRemoveRoleFromUser, err)
	}

	return pe.validator.RemoveRoleFromUser(roleId, userId)
}
