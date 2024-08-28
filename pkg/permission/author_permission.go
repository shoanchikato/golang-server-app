package permission

import m "app/pkg/model"

var (
	AuthorAdd    = m.NewPermission("author: add", "author", "add")
	AuthorAddAll = m.NewPermission("author: add all", "author", "get all")
	AuthorGetOne = m.NewPermission("author: get one", "author", "one")
	AuthorGetAll = m.NewPermission("author: get all", "author", "get all")
	AuthorEdit   = m.NewPermission("author: edit", "author", "edit")
	AuthorRemove = m.NewPermission("author: remove", "author", "remove")
)
