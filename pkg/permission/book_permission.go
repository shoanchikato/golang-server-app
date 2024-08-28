package permission

import m "app/pkg/model"

var (
	BookAdd    = m.NewPermission("book: add", "book", "add")
	BookAddAll = m.NewPermission("book: add all", "book", "add all")
	BookGetOne = m.NewPermission("book: get one", "book", "get one")
	BookGetAll = m.NewPermission("book: get all", "book", "get all")
	BookEdit   = m.NewPermission("book: edit", "book", "edit")
	BookRemove = m.NewPermission("book: remove", "book", "remove")
)
