package permission

import m "app/pkg/model"

var (
	PermissionAdd    = m.NewPermission("permission: add", "permission", "add")
	PermissionAddAll = m.NewPermission("permission: add all", "permission", "get all")
	PermissionGetOne = m.NewPermission("permission: get one", "permission", "one")
	PermissionGetAll = m.NewPermission("permission: get all", "permission", "get all")
	PermissionEdit   = m.NewPermission("permission: edit", "permission", "edit")
	PermissionRemove = m.NewPermission("permission: remove", "permission", "remove")
)
