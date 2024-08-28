package permission

import m "app/pkg/model"

var (
	PermissionManagementAddPermissionToRole       = m.NewPermission("permission management: addPermissionToRole", "permission management", "addPermissionToRole")
	PermissionManagementAddPermissionsToRole      = m.NewPermission("permission management: addPermissionsToRole", "permission management", "addPermissionsToRole")
	PermissionManagementGetPermissionsByRoleId    = m.NewPermission("permission management: getPermissionsByRoleId", "permission management", "getPermissionsByRoleId")
	PermissionManagementGetPermissonsByUserId     = m.NewPermission("permission management: getPermissonsByUserId", "permission management", "getPermissonsByUserId")
	PermissionManagementRemovePermissionFromRole  = m.NewPermission("permission management: removePermissionFromRole", "permission management", "removePermissionFromRole")
	PermissionManagementRemovePermissionsFromRole = m.NewPermission("permission management: removePermissionsFromRole", "permission management", "removePermissionsFromRole")
)
