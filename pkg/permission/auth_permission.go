package permission

import m "app/pkg/model"

var (
	AuthAdd    = m.NewPermission("auth: add", "auth", "add")
	AuthAddAll = m.NewPermission("auth: add all", "auth", "get all")
	AuthGetOne = m.NewPermission("auth: get one", "auth", "one")
	AuthGetAll = m.NewPermission("auth: get all", "auth", "get all")
	AuthEdit   = m.NewPermission("auth: edit", "auth", "edit")
	AuthRemove = m.NewPermission("auth: remove", "auth", "remove")
)
