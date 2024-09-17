package httperrorfmt

import (
	a "app/pkg/authorization"
	m "app/pkg/model"
	s "app/pkg/service"
)

type RoleManagementHttpErrorFmt interface {
	AddRoleToUser(adminId int, roleId, userId int) error
	GetRolesByUserId(adminId int, userId int) (*[]m.Role, error)
	RemoveRoleFromUser(adminId int, roleId, userId int) error
}

type roleManagementHttpErrorFmt struct {
	authorization a.RoleManagementAuthorization
	service       s.HttpErrorFmt
}

func NewRoleManagementHttpErrorFmt(authorization a.RoleManagementAuthorization, service s.HttpErrorFmt) RoleManagementHttpErrorFmt {
	return &roleManagementHttpErrorFmt{authorization, service}
}

func (r *roleManagementHttpErrorFmt) AddRoleToUser(adminId int, roleId, userId int) error {
	err := r.authorization.AddRoleToUser(adminId, roleId, userId)
	if err != nil {
		return r.service.GetError(err)
	}

	return nil
}

func (r *roleManagementHttpErrorFmt) GetRolesByUserId(adminId int, userId int) (*[]m.Role, error) {
	roles, err := r.authorization.GetRolesByUserId(adminId, userId)
	if err != nil {
		return nil, r.service.GetError(err)
	}

	return roles, nil
}

func (r *roleManagementHttpErrorFmt) RemoveRoleFromUser(adminId int, roleId, userId int) error {
	err := r.authorization.RemoveRoleFromUser(adminId, roleId, userId)
	if err != nil {
		return r.service.GetError(err)
	}

	return nil
}
