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
	Author               AuthorHttpErrorFmt
	Book                 BookHttpErrorFmt
	Post                 PostHttpErrorFmt
}

func HttpErrorFmtDi(httpErrorFmt s.HttpErrorFmt, jwt s.JWTService, authorizations *a.Authorizations) *HttpErrorFmts {
	user := NewUserHttpErrorFmt(authorizations.User, httpErrorFmt)
	auth := NewAuthHttpErrorFmt(authorizations.Auth, jwt, httpErrorFmt)
	permission := NewPermissionHttpErrorFmt(authorizations.Permission, httpErrorFmt)
	permissionManagement := NewPermissionManagementHttpErrorFmt(authorizations.PermissionManagement, httpErrorFmt)
	role := NewRoleHttpErrorFmt(authorizations.Role, httpErrorFmt)
	roleManagement := NewRoleManagementHttpErrorFmt(authorizations.RoleManagement, httpErrorFmt)
	author := NewAuthorHttpErrorFmt(authorizations.Author, httpErrorFmt)
	book := NewBookHttpErrorFmt(authorizations.Book, httpErrorFmt)
	post := NewPostHttpErrorFmt(authorizations.Post, httpErrorFmt)

	return &HttpErrorFmts{auth, user, permission, permissionManagement, role, roleManagement, author, book, post}
}
