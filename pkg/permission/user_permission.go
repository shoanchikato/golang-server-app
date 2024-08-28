package permission

import m "app/pkg/model"

var (
	UserAdd    = m.NewPermission("user: add", "user", "add")
	UserAddAll = m.NewPermission("user: add all", "user", "add all")
	UserGetOne = m.NewPermission("user: get one", "user", "get one")
	UserGetAll = m.NewPermission("user: get all", "user", "get all")
	UserEdit   = m.NewPermission("user: edit", "user", "edit")
	UserRemove = m.NewPermission("user: remove", "user", "remove")
)
