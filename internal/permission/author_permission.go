package permission

import m "app/internal/model"

var (
	AuthorAdd     = m.NewPermission("author: add", "author", "add")
	AuthorAddAll  = m.NewPermission("author: add all", "author", "add all")
	AuthorGetOne  = m.NewPermission("author: get one", "author", "get one")
	AuthorGetAll  = m.NewPermission("author: get all", "author", "get all")
	AuthorGetMore = m.NewPermission("author: get more", "author", "get more")
	AuthorEdit    = m.NewPermission("author: edit", "author", "edit")
	AuthorRemove  = m.NewPermission("author: remove", "author", "remove")
)
