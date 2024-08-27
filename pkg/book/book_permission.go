package book

import a "app/pkg/authorization"

var (
	BookAdd    a.AuthPermission = "book: add"
	BookAddAll a.AuthPermission = "book: add all"
	BookGetOne a.AuthPermission = "book: get one"
	BookGetAll a.AuthPermission = "book: get all"
	BookEdit   a.AuthPermission = "book: edit"
	BookRemove a.AuthPermission = "book: remove"
)
