package post

import p "app/pkg/permission"

var (
	PostAdd    = p.NewPermission("post: add", "post", "add")
	PostAddAll = p.NewPermission("post: add all", "post", "add all")
	PostGetOne = p.NewPermission("post: get one", "post", "one")
	PostGetAll = p.NewPermission("post: get all", "post", "get all")
	PostEdit   = p.NewPermission("post: edit", "post", "edit")
	PostRemove = p.NewPermission("post: remove", "post", "remove")
)
