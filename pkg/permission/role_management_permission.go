package permission

import m "app/pkg/model"

var (
	RoleManagementAddRoleToUser      = m.NewPermission("role management: addRoleToUser", "role management", "addRoleToUser")
	RoleManagementGetRoleByUserId    = m.NewPermission("role management: getRoleByUserId", "role management", "getRoleByUserId")
	RoleManagementRemoveRoleFromUser = m.NewPermission("role management: removeRoleFromUser", "role management", "removeRoleFromUser")
)
