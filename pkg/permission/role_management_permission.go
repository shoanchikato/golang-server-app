package permission

import m "app/pkg/model"

var (
	RoleManagementAddRoleToUser      = m.NewPermission("role management: add role to user", "role management", "add role to user")
	RoleManagementGetRolesByUserId   = m.NewPermission("role management: get role by user id", "role management", "get role by user id")
	RoleManagementRemoveRoleFromUser = m.NewPermission("role management: remove role from user", "role management", "remove role from user")
)
