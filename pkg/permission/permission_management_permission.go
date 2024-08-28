package permission

var (
	PermissionManagementAdd    = NewPermission("permission management: add", "permission management", "add")
	PermissionManagementAddAll = NewPermission("permission management: add all", "permission management", "get all")
	PermissionManagementGetOne = NewPermission("permission management: get one", "permission management", "one")
	PermissionManagementGetAll = NewPermission("permission management: get all", "permission management", "get all")
	PermissionManagementEdit   = NewPermission("permission management: edit", "permission management", "edit")
	PermissionManagementRemove = NewPermission("permission management: remove", "permission management", "remove")
)
