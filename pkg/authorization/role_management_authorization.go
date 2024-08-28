package authorization

import (
	e "app/pkg/errors"
	m "app/pkg/model"
	p "app/pkg/permission"
	s "app/pkg/service"
	v "app/pkg/validation"
	"errors"
)

type RoleManagementAuthorization interface {
	AddRoleToUser(adminId int, roleId, userId int) error
	GetRoleByUserId(adminId int, userId int) (*m.Role, error)
	RemoveRoleFromUser(adminId int, roleId, userId int) error
}

type roleManagementAuthorization struct {
	auth      s.AuthorizationService
	validator v.RoleManagementValidator
}

func NewRoleManagementAuthorization(
	auth s.AuthorizationService,
	validator v.RoleManagementValidator,
) RoleManagementAuthorization {
	return &roleManagementAuthorization{auth, validator}
}

// AddRoleToUser
func (r *roleManagementAuthorization) AddRoleToUser(adminId int, roleId int, userId int) error {
	err := r.auth.CheckForAuthorization(adminId, p.RoleManagementAddRoleToUser.Name)
	if err != nil {
		return errors.Join(e.ErrAuthorDomain, e.ErrOnAddRoleToUser, err)
	}

	return r.validator.AddRoleToUser(roleId, userId)
}

// GetRoleByUserId
func (r *roleManagementAuthorization) GetRoleByUserId(adminId int, userId int) (*m.Role, error) {
	err := r.auth.CheckForAuthorization(adminId, p.RoleManagementGetRoleByUserId.Name)
	if err != nil {
		return nil, errors.Join(e.ErrAuthorDomain, e.ErrOnGetRoleByUserId, err)
	}

	return r.validator.GetRoleByUserId(userId)
}

// RemoveRoleFromUser
func (r *roleManagementAuthorization) RemoveRoleFromUser(adminId int, roleId int, userId int) error {
	err := r.auth.CheckForAuthorization(adminId, p.RoleManagementRemoveRoleFromUser.Name)
	if err != nil {
		return errors.Join(e.ErrAuthorDomain, e.ErrOnRemoveRoleFromUser, err)
	}

	return r.validator.RemoveRoleFromUser(roleId, userId)
}
