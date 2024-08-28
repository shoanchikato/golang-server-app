package permission

var (
	PermissionAdd    = NewPermission("permission: add", "permission", "add")
	PermissionAddAll = NewPermission("permission: add all", "permission", "get all")
	PermissionGetOne = NewPermission("permission: get one", "permission", "one")
	PermissionGetAll = NewPermission("permission: get all", "permission", "get all")
	PermissionEdit   = NewPermission("permission: edit", "permission", "edit")
	PermissionRemove = NewPermission("permission: remove", "permission", "remove")
)
