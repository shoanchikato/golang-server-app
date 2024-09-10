package validation

import (
	m "app/pkg/model"
	r "app/pkg/repo"
	s "app/pkg/service"
)

type PermissionManagementValidator interface {
	AddPermissionToRole(permissionId, roleId int) error
	AddPermissionsToRole(permissionIds []int, roleId int) error
	GetPermissionsByRoleId(roleId int) (*[]m.Permission, error)
	GetPermissonsByUserId(userId int) (*[]m.Permission, error)
	RemovePermissionFromRole(roleId, permissionId int) error
	RemovePermissionsFromRole(roleId int, permissionIds []int) error
}

type pMValidator struct {
	repo r.PermissionManagementRepo
	service s.ValidationService
}

func NewPermissionManagementValidator(repo r.PermissionManagementRepo, validation s.ValidationService) PermissionManagementValidator {
	return &pMValidator{repo, validation}
}

// AddPermissionToRole
func (p *pMValidator) AddPermissionToRole(permissionId, roleId int) error {
	return p.repo.AddPermissionToRole(permissionId, roleId)
}

// AddPermissionsToRole
func (p *pMValidator) AddPermissionsToRole(permissionIds []int, roleId int) error {
	return p.repo.AddPermissionsToRole(permissionIds, roleId)
}

// GetPermissionsByRoleId
func (p *pMValidator) GetPermissionsByRoleId(roleId int) (*[]m.Permission, error) {
	return p.repo.GetPermissionsByRoleId(roleId)
}

// GetPermissonsByUserId
func (p *pMValidator) GetPermissonsByUserId(userId int) (*[]m.Permission, error) {
	return p.repo.GetPermissonsByUserId(userId)
}

// RemovePermissionFromRole
func (p *pMValidator) RemovePermissionFromRole(roleId int, permissionId int) error {
	return p.repo.RemovePermissionFromRole(roleId, permissionId)
}

// RemovePermissionsFromRole
func (p *pMValidator) RemovePermissionsFromRole(roleId int, permissionIds []int) error {
	return p.repo.RemovePermissionsFromRole(roleId, permissionIds)
}
