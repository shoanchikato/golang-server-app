package permission

import (
	e "app/pkg/errors"
	"errors"
)

type AuthorizationService interface {
	CheckForAuthorization(userID int, permission string) error
}

type authorizationService struct {
	repo PermissionManagementRepo
}

func NewAuthorizationService(repo PermissionManagementRepo) AuthorizationService {
	return &authorizationService{repo}
}

func (a *authorizationService) hasPermission(permission string, permissions *[]Permission) bool {
	pp := *permissions
	for i := 0; i < len(pp); i++ {
		if string(permission) == pp[i].Name {
			return true
		}
	}

	return false
}

func (a *authorizationService) getPermissions(userID int) (*[]Permission, error) {
	permissions, err := a.repo.GetPermissonsByUserID(userID)
	if err != nil {
		return nil, err
	}

	return permissions, nil
}

func (a *authorizationService) CheckForAuthorization(userID int, permission string) error {
	permissions, err := a.getPermissions(userID)
	if err != nil {
		return errors.Join(e.ErrNotAuthorized, err)
	}

	isAuthorized := a.hasPermission(permission, permissions)
	if !isAuthorized {
		return e.ErrNotAuthorized
	}

	return nil
}
