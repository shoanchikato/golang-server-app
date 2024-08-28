package permission

import m "app/pkg/model"

var (
	PostAdd    = m.NewPermission("post: add", "post", "add")
	PostAddAll = m.NewPermission("post: add all", "post", "add all")
	PostGetOne = m.NewPermission("post: get one", "post", "one")
	PostGetAll = m.NewPermission("post: get all", "post", "get all")
	PostEdit   = m.NewPermission("post: edit", "post", "edit")
	PostRemove = m.NewPermission("post: remove", "post", "remove")
)
