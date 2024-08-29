package errors

import "errors"

var (
	ErrOnAdd     = errors.New("add: ")
	ErrOnAddAll  = errors.New("add app: ")
	ErrOnEdit    = errors.New("edit: ")
	ErrOnGetAll  = errors.New("get all: ")
	ErrOnGetOne  = errors.New("get one: ")
	ErrOnGetMore = errors.New("get more: ")
	ErrOnRemove  = errors.New("remove post: ")

	ErrOnGetByEntity = errors.New("get by entity: ")

	ErrOnLogin         = errors.New("login: ")
	ErrOnResetPassword = errors.New("reset password: ")

	ErrOnAddPermissionToRole       = errors.New("add permission to role: ")
	ErrOnAddPermissionsToRole      = errors.New("add permissions to role: ")
	ErrOnAddRoleToUser             = errors.New("add role to user: ")
	ErrOnGetPermissionsByRoleId    = errors.New("get permissions by role id: ")
	ErrOnGetPermissonsByUserId     = errors.New("get permissons by user id: ")
	ErrOnGetRoleByUserId           = errors.New("get role by user id: ")
	ErrOnRemovePermissionFromRole  = errors.New("remove permission from role: ")
	ErrOnRemovePermissionsFromRole = errors.New("remove permissions from role: ")
	ErrOnRemoveRoleFromUser        = errors.New("remove role from user: ")
)
