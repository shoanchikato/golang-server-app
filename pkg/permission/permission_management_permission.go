package permission

import m "app/pkg/model"

var (
	PermissionManagementAdd    = m.NewPermission("permission management: add", "permission management", "add")
	PermissionManagementAddAll = m.NewPermission("permission management: add all", "permission management", "add all")
	PermissionManagementGetOne = m.NewPermission("permission management: get one", "permission management", "get one")
	PermissionManagementGetAll = m.NewPermission("permission management: get all", "permission management", "get all")
	PermissionManagementEdit   = m.NewPermission("permission management: edit", "permission management", "edit")
	PermissionManagementRemove = m.NewPermission("permission management: remove", "permission management", "remove")
)
