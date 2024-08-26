package errors

import "errors"

var (
	ErrAuthDomain       = errors.New("auth: ")
	ErrAuthorDomain     = errors.New("author: ")
	ErrBookDomain       = errors.New("book: ")
	ErrPermissionDomain = errors.New("permission: ")
	ErrPostDomain       = errors.New("post: ")
	ErrRoleDomain       = errors.New("role: ")
	ErrUserDomain       = errors.New("user: ")
	ErrPermissionManagement = errors.New("permission management: ")
)
