package httperrorfmt

import (
	a "app/pkg/authorization"
	s "app/pkg/service"
)

type HttpErrorFmts struct {
	Auth                 AuthHttpErrorFmt
	User                 UserHttpErrorFmt
	Permission           PermissionHttpErrorFmt
	PermissionManagement PermissionManagementHttpErrorFmt
	Role                 RoleHttpErrorFmt
	RoleManagement       RoleManagementHttpErrorFmt
}

func HttpErrorFmtDi(httpErrorFmt s.HttpErrorFmt, jwt s.JWTService, authorizations *a.Authorizations) *HttpErrorFmts {
	user := NewUserHttpErrorFmt(authorizations.User, httpErrorFmt)
	auth := NewAuthHttpErrorFmt(authorizations.Auth, jwt, httpErrorFmt)
	permission := NewPermissionHttpErrorFmt(authorizations.Permission, httpErrorFmt)
	permissionManagement := NewPermissionManagementHttpErrorFmt(authorizations.PermissionManagement, httpErrorFmt)
	role := NewRoleHttpErrorFmt(authorizations.Role, httpErrorFmt)
	roleManagement := NewRoleManagementHttpErrorFmt(authorizations.RoleManagement, httpErrorFmt)

	return &HttpErrorFmts{auth, user, permission, permissionManagement, role, roleManagement}
}
