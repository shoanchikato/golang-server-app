package service

import (
	e "app/pkg/errors"
	m "app/pkg/model"
	r "app/pkg/repo"
	"errors"
)

type AuthorizationService interface {
	CheckForAuthorization(userId int, permission string) error
}

type authorizationService struct {
	repo r.PermissionManagementRepo
}

func NewAuthorizationService(repo r.PermissionManagementRepo) AuthorizationService {
	return &authorizationService{repo}
}

func (a *authorizationService) hasPermission(permissionName string, permissions *[]m.Permission) bool {
	for _, permission := range *permissions {
		if permissionName == permission.Name {
			return true
		}
	}

	return false
}

func (a *authorizationService) getPermissions(userId int) (*[]m.Permission, error) {
	permissions, err := a.repo.GetPermissonsByUserId(userId)
	if err != nil {
		return nil, err
	}

	return permissions, nil
}

func (a *authorizationService) CheckForAuthorization(userId int, permission string) error {
	permissions, err := a.getPermissions(userId)
	if err != nil {
		return errors.Join(e.ErrNotAuthorized, err)
	}

	isAuthorized := a.hasPermission(permission, permissions)
	if !isAuthorized {
		return e.ErrNotAuthorized
	}

	return nil
}
