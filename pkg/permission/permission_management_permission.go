package permission

import m "app/pkg/model"

var (
	PermissionManagementAddPermissionToRole       = m.NewPermission("permission management: addPermissionToRole", "permission management", "addPermissionToRole")
	PermissionManagementAddPermissionsToRole      = m.NewPermission("permission management: addPermissionsToRole", "permission management", "addPermissionsToRole")
	PermissionManagementAddRoleToUser             = m.NewPermission("permission management: addRoleToUser", "permission management", "addRoleToUser")
	PermissionManagementGetPermissionsByRoleId    = m.NewPermission("permission management: getPermissionsByRoleId", "permission management", "getPermissionsByRoleId")
	PermissionManagementGetPermissonsByUserId     = m.NewPermission("permission management: getPermissonsByUserId", "permission management", "getPermissonsByUserId")
	PermissionManagementGetRoleByUserId           = m.NewPermission("permission management: getRoleByUserId", "permission management", "getRoleByUserId")
	PermissionManagementRemovePermissionFromRole  = m.NewPermission("permission management: removePermissionFromRole", "permission management", "removePermissionFromRole")
	PermissionManagementRemovePermissionsFromRole = m.NewPermission("permission management: removePermissionsFromRole", "permission management", "removePermissionsFromRole")
	PermissionManagementRemoveRoleFromUser        = m.NewPermission("permission management: removeRoleFromUser", "permission management", "removeRoleFromUser")
)
