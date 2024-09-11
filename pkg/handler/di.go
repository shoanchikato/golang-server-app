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
	Author               AuthorHandler
	Book                 BookHandler
	Post                 PostHandler
}

func HandlerDi(httpErrorFmts *ef.HttpErrorFmts, jwt s.JWTService, logger s.Logger) *Handlers {
	auth := NewAuthHandler(httpErrorFmts.Auth, jwt)
	user := NewUserHandler(httpErrorFmts.User, logger)
	role := NewRoleHandler(httpErrorFmts.Role, logger)
	roleManagement := NewRoleManagementHandler(httpErrorFmts.RoleManagement, logger)
	permission := NewPermissionHandler(httpErrorFmts.Permission, logger)
	permissionManagement := NewPermissionManagementHandler(httpErrorFmts.PermissionManagement, logger)
	author := NewAuthorHandler(httpErrorFmts.Author, logger)
	book := NewBookHandler(httpErrorFmts.Book, logger)
	post := NewPostHandler(httpErrorFmts.Post, logger)

	return &Handlers{auth, user, role, roleManagement, permission, permissionManagement, author, book, post}
}
