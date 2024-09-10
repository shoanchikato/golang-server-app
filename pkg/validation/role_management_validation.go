package validation

import (
	m "app/pkg/model"
	r "app/pkg/repo"
	s "app/pkg/service"
)

type RoleManagementValidator interface {
	AddRoleToUser(roleId, userId int) error
	GetRoleByUserId(userId int) (*m.Role, error)
	RemoveRoleFromUser(roleId, userId int) error
}

type rMValidator struct {
	repo r.RoleManagementRepo
	service s.ValidationService
}

func NewRoleManagementValidator(repo r.RoleManagementRepo, service s.ValidationService) RoleManagementValidator {
	return &rMValidator{repo, service}
}

// AddRoleToUser
func (r *rMValidator) AddRoleToUser(roleId int, userId int) error {
	return r.repo.AddRoleToUser(roleId, userId)
}

// GetRoleByUserId
func (r *rMValidator) GetRoleByUserId(userId int) (*m.Role, error) {
	return r.repo.GetRoleByUserId(userId)
}

// RemoveRoleFromUser
func (r *rMValidator) RemoveRoleFromUser(roleId int, userId int) error {
	return r.repo.RemoveRoleFromUser(roleId, userId)
}
