package validation

import (
	m "app/pkg/model"
	r "app/pkg/repo"
)

type PermissionManagementValidator interface {
	AddPermissionToRole(permissionId, roleId int) error
	AddPermissionsToRole(permissionIds []int, roleId int) error
	AddRoleToUser(roleId, userId int) error
	GetPermissionsByRoleId(roleId int) (*[]m.Permission, error)
	GetPermissonsByUserId(userId int) (*[]m.Permission, error)
	GetRoleByUserId(userId int) (*m.Role, error)
	RemovePermissionFromRole(roleId, permissionId int) error
	RemovePermissionsFromRole(roleId int, permissionIds []int) error
	RemoveRoleFromUser(roleId, userId int) error
}

type pMValidator struct {
	repo r.PermissionManagementRepo
}

func NewPermissionManagementValidator(repo r.PermissionManagementRepo) PermissionManagementValidator {
	return &pMValidator{repo}
}

// AddPermissionToRole
func (p *pMValidator) AddPermissionToRole(permissionId, roleId int) error {
	return p.repo.AddPermissionToRole(permissionId, roleId)
}

// AddPermissionsToRole
func (p *pMValidator) AddPermissionsToRole(permissionIds []int, roleId int) error {
	return p.repo.AddPermissionsToRole(permissionIds, roleId)
}

// AddRoleToUser
func (p *pMValidator) AddRoleToUser(roleId int, userId int) error {
	return p.repo.AddRoleToUser(roleId, userId)
}

// GetPermissionsByRoleId
func (p *pMValidator) GetPermissionsByRoleId(roleId int) (*[]m.Permission, error) {
	return p.repo.GetPermissionsByRoleId(roleId)
}

// GetPermissonsByUserId
func (p *pMValidator) GetPermissonsByUserId(userId int) (*[]m.Permission, error) {
	return p.repo.GetPermissonsByUserId(userId)
}

// GetRoleByUserId
func (p *pMValidator) GetRoleByUserId(userId int) (*m.Role, error) {
	return p.repo.GetRoleByUserId(userId)
}

// RemovePermissionFromRole
func (p *pMValidator) RemovePermissionFromRole(roleId int, permissionId int) error {
	return p.repo.RemovePermissionFromRole(roleId, permissionId)
}

// RemovePermissionsFromRole
func (p *pMValidator) RemovePermissionsFromRole(roleId int, permissionIds []int) error {
	return p.repo.RemovePermissionsFromRole(roleId, permissionIds)
}

// RemoveRoleFromUser
func (p *pMValidator) RemoveRoleFromUser(roleId int, userId int) error {
	return p.repo.RemoveRoleFromUser(roleId, userId)
}
