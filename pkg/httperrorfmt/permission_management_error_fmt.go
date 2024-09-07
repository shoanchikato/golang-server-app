package httperrorfmt

import (
	a "app/pkg/authorization"
	m "app/pkg/model"
	s "app/pkg/service"
)

type PermissionManagementHttpErrorFmt interface {
	AddPermissionToRole(adminId int, permissionId, roleId int) error
	AddPermissionsToRole(adminId int, permissionIds []int, roleId int) error
	GetPermissionsByRoleId(adminId int, roleId int) (*[]m.Permission, error)
	GetPermissionsByUserId(adminId int, userId int) (*[]m.Permission, error)
	RemovePermissionFromRole(adminId int, roleId, permissionId int) error
	RemovePermissionsFromRole(adminId int, roleId int, permissionIds []int) error
}

type permissionManagementHttpErrorFmt struct {
	authorization a.PermissionManagementAuthorization
	service       s.HttpErrorFmt
}

func NewPermissionManagementHttpErrorFmt(authorization a.PermissionManagementAuthorization, service s.HttpErrorFmt) PermissionManagementHttpErrorFmt {
	return &permissionManagementHttpErrorFmt{authorization, service}
}

// AddPermissionToRole
func (p *permissionManagementHttpErrorFmt) AddPermissionToRole(adminId int, permissionId int, roleId int) error {
	err := p.authorization.AddPermissionToRole(adminId, permissionId, roleId)
	if err != nil {
		return p.service.GetError(err)
	}

	return nil
}

// AddPermissionsToRole
func (p *permissionManagementHttpErrorFmt) AddPermissionsToRole(adminId int, permissionIds []int, roleId int) error {
	err := p.authorization.AddPermissionsToRole(adminId, permissionIds, roleId)
	if err != nil {
		return p.service.GetError(err)
	}

	return nil
}

// GetPermissionsByRoleId
func (p *permissionManagementHttpErrorFmt) GetPermissionsByRoleId(adminId int, roleId int) (*[]m.Permission, error) {
	permissions, err := p.authorization.GetPermissionsByRoleId(adminId, roleId)
	if err != nil {
		return nil, p.service.GetError(err)
	}

	return permissions, nil
}

// GetPermissonsByUserId
func (p *permissionManagementHttpErrorFmt) GetPermissionsByUserId(adminId int, userId int) (*[]m.Permission, error) {
	permissions, err := p.authorization.GetPermissonsByUserId(adminId, userId)
	if err != nil {
		return nil, p.service.GetError(err)
	}

	return permissions, nil
}

// RemovePermissionFromRole
func (p *permissionManagementHttpErrorFmt) RemovePermissionFromRole(adminId int, roleId int, permissionId int) error {
	err := p.authorization.RemovePermissionFromRole(adminId, roleId, permissionId)
	if err != nil {
		return p.service.GetError(err)
	}

	return nil
}

// RemovePermissionsFromRole
func (p *permissionManagementHttpErrorFmt) RemovePermissionsFromRole(adminId int, roleId int, permissionIds []int) error {
	err := p.authorization.RemovePermissionsFromRole(adminId, roleId, permissionIds)
	if err != nil {
		return p.service.GetError(err)
	}

	return nil
}
