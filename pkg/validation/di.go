package validation

import (
	en "app/pkg/encrypt"
	r "app/pkg/repo"
	s "app/pkg/service"
)

type Validators struct {
	User                 UserValidator
	Auth                 AuthValidator
	Role                 RoleValidator
	Permission           PermissionValidator
	Author               AuthorValidator
	Book                 BookValidator
	Post                 PostValidator
	RoleManagement       RoleManagementValidator
	PermissionManagement PermissionManagementValidator
}

func ValidationDi(repos *r.Repos, encryptions *en.Encryptions, validation s.ValidationService) *Validators {
	user := NewUserValidator(encryptions.User, validation)
	auth := NewAuthValidator(encryptions.Auth, validation)
	role := NewRoleValidator(repos.Role, validation)
	permission := NewPermissionValidator(repos.Permission, validation)
	author := NewAuthorValidator(repos.Author, validation)
	book := NewBookValidator(repos.Book, validation)
	post := NewPostValidator(repos.Post, validation)
	roleManagment := NewRoleManagementValidator(repos.RoleManagement, validation)
	permissionManagement := NewPermissionManagementValidator(repos.PermissionManagement, validation)

	return &Validators{user, auth, role, permission, author, book, post, roleManagment, permissionManagement}
}
