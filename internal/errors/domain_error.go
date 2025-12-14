package errors

import "errors"

var (
	ErrAuthDomain                 = errors.New("auth: ")
	ErrAuthorDomain               = errors.New("author: ")
	ErrBookDomain                 = errors.New("book: ")
	ErrPermissionDomain           = errors.New("permission: ")
	ErrPostDomain                 = errors.New("post: ")
	ErrRoleDomain                 = errors.New("role: ")
	ErrUserDomain                 = errors.New("user: ")
	ErrPermissionManagementDomain = errors.New("permission management: ")
	ErrRoleManagementDomain       = errors.New("role management: ")
)
