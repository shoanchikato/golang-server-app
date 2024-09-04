package validation

import (
	en "app/pkg/encrypt"
	r "app/pkg/repo"
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

func ValidationDi(repos *r.Repos, encryptions *en.Encryptions) *Validators {
	user := NewUserValidator(encryptions.User)
	auth := NewAuthValidator(encryptions.Auth)
	role := NewRoleValidator(repos.Role)
	permission := NewPermissionValidator(repos.Permission)
	author := NewAuthorValidator(repos.Author)
	book := NewBookValidator(repos.Book)
	post := NewPostValidator(repos.Post)
	roleManagment := NewRoleManagementValidator(repos.RoleManagement)
	permissionManagement := NewPermissionManagementValidator(repos.PermissionManagement)

	return &Validators{user, auth, role, permission, author, book, post, roleManagment, permissionManagement}
}
