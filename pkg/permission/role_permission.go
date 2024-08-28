package permission

import m "app/pkg/model"

var (
	RoleAdd    = m.NewPermission("role: add", "role", "add")
	RoleAddAll = m.NewPermission("role: add all", "role", "get all")
	RoleGetOne = m.NewPermission("role: get one", "role", "one")
	RoleGetAll = m.NewPermission("role: get all", "role", "get all")
	RoleEdit   = m.NewPermission("role: edit", "role", "edit")
	RoleRemove = m.NewPermission("role: remove", "role", "remove")
)
