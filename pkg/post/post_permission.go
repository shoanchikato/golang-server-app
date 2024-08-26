package post

import a "app/pkg/authorization"

const (
	PostAdd    a.AuthPermission = "post: add"
	PostAddAll a.AuthPermission = "post: add all"
	PostGetOne a.AuthPermission = "post: get one"
	PostGetAll a.AuthPermission = "post: get all"
	PostEdit   a.AuthPermission = "post: edit"
	PostRemove a.AuthPermission = "post: remove"
)
