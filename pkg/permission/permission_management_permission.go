package permission

import m "app/pkg/model"

var (
	PermissionManagementAddPermissionToRole       = m.NewPermission("permission management: add permission to role", "permission management", "add permission to role")
	PermissionManagementAddPermissionsToRole      = m.NewPermission("permission management: add permissions to role", "permission management", "add permissions to role")
	PermissionManagementGetPermissionsByRoleId    = m.NewPermission("permission management: get permissions by role id", "permission management", "get permissions by role id")
	PermissionManagementGetPermissonsByUserId     = m.NewPermission("permission management: get permissions by user id", "permission management", "get permissions by user id")
	PermissionManagementRemovePermissionFromRole  = m.NewPermission("permission management: remove permission from role", "permission management", "remove permission from role")
	PermissionManagementRemovePermissionsFromRole = m.NewPermission("permission management: remove permissions from role", "permission management", "remove permissions from role")
)
