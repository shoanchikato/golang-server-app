package authorization

import (
	s "app/pkg/service"
	v "app/pkg/validation"
)

type Authorizations struct {
	User                 UserAuthorization
	Auth                 AuthAuthorization
	Role                 RoleAuthorization
	Permission           PermissionAuthorization
	Author               AuthorAuthorization
	Book                 BookAuthorization
	Post                 PostAuthorization
	RoleManagement       RoleManagementAuthorization
	PermissionManagement PermissionManagementAuthorization
}

func AuthorizationDi(authService s.AuthorizationService, validators *v.Validators) *Authorizations {
	user := NewUserAuthorization(authService, validators.User)
	auth := NewAuthAuthorization(authService, validators.Auth)
	role := NewRoleAuthorization(authService, validators.Role)
	permission := NewPermissionAuthorization(authService, validators.Permission)
	author := NewAuthorAuthorization(authService, validators.Author)
	book := NewBookAuthorization(authService, validators.Book)
	post := NewPostAuthorization(authService, validators.Post)
	roleManagement := NewRoleManagementAuthorization(authService, validators.RoleManagement)
	permissionManagement := NewPermissionManagementAuthorization(authService, validators.PermissionManagement)

	return &Authorizations{user, auth, role, permission, author, book, post, roleManagement, permissionManagement}
}
