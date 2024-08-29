package permission

import m "app/pkg/model"

var (
	RoleManagementAddRoleToUser      = m.NewPermission("role management: add role to user", "role management", "add role to user")
	RoleManagementGetRoleByUserId    = m.NewPermission("role management: get role by user id", "role management", "get role by user id")
	RoleManagementRemoveRoleFromUser = m.NewPermission("role management: remove role from user", "", "remove role from user")
)
