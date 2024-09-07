package handler

import (
	ef "app/pkg/httperrorfmt"
	s "app/pkg/service"
)

type Handlers struct {
	Auth                 AuthHandler
	User                 UserHandler
	Role                 RoleHandler
	RoleManagement       RoleManagementHandler
	Permission           PermissionHandler
	PermissionManagement PermissionManagementHandler
}

func HandlerDi(httpErrorFmts *ef.HttpErrorFmts, jwt s.JWTService) *Handlers {
	auth := NewAuthHandler(httpErrorFmts.Auth, jwt)
	user := NewUserHandler(httpErrorFmts.User)
	role := NewRoleHandler(httpErrorFmts.Role)
	roleManagement := NewRoleManagementHandler(httpErrorFmts.RoleManagement)
	permission := NewPermissionHandler(httpErrorFmts.Permission)
	permissionManagement := NewPermissionManagementHandler(httpErrorFmts.PermissionManagement)

	return &Handlers{auth, user, role, roleManagement, permission, permissionManagement}
}
