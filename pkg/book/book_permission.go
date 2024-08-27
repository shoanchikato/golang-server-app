package book

import p "app/pkg/permission"

var (
	BookAdd    = p.NewPermission("book: add", "book", "add")
	BookAddAll = p.NewPermission("book: add all", "book", "get all")
	BookGetOne = p.NewPermission("book: get one", "book", "one")
	BookGetAll = p.NewPermission("book: get all", "book", "get all")
	BookEdit   = p.NewPermission("book: edit", "book", "edit")
	BookRemove = p.NewPermission("book: remove", "book", "remove")
)
