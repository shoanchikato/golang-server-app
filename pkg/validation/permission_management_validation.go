package validation

import (
	r "app/pkg/repo"
	m "app/pkg/model"
)

type PermissionManagementValidator interface {
	AddPermissionToRole(permissionID, roleID int) error
	AddPermissionsToRole(permissionIDs []int, roleID int) error
	AddRoleToUser(roleID, userID int) error
	GetPermissionsByRoleID(roleID int) (*[]m.Permission, error)
	GetPermissonsByUserID(userID int) (*[]m.Permission, error)
	GetRoleByUserID(userID int) (*m.Role, error)
	RemovePermissionFromRole(roleID, permissionID int) error
	RemovePermissionsFromRole(roleID int, permissionIDs []int) error
	RemoveRoleFromUser(roleID, userID int) error
}

type pMValidator struct {
	repo r.PermissionManagementRepo
}

func NewPermissionManagementValidator(repo r.PermissionManagementRepo) PermissionManagementValidator {
	return &pMValidator{repo}
}

// AddPermissionToRole
func (p *pMValidator) AddPermissionToRole(permissionID, roleID int) error {
	return p.repo.AddPermissionToRole(permissionID, roleID)
}

// AddPermissionsToRole
func (p *pMValidator) AddPermissionsToRole(permissionIDs []int, roleID int) error {
	return p.repo.AddPermissionsToRole(permissionIDs, roleID)
}

// AddRoleToUser
func (p *pMValidator) AddRoleToUser(roleID int, userID int) error {
	return p.repo.AddRoleToUser(roleID, userID)
}

// GetPermissionsByRoleID
func (p *pMValidator) GetPermissionsByRoleID(roleID int) (*[]m.Permission, error) {
	return p.repo.GetPermissionsByRoleID(roleID)
}

// GetPermissonsByUserID
func (p *pMValidator) GetPermissonsByUserID(userID int) (*[]m.Permission, error) {
	return p.repo.GetPermissonsByUserID(userID)
}

// GetRoleByUserID
func (p *pMValidator) GetRoleByUserID(userID int) (*m.Role, error) {
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
