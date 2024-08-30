package main

import (
	m "app/pkg/model"
	pe "app/pkg/permission"
)

func Data() (
	[]*m.Author,
	[]*m.Book,
	[]*m.Post,
	[]*m.User,
	[]*m.Permission,
) {
	posts := []*m.Post{
		m.NewPost("one", "one body", 1),
		m.NewPost("two", "two body", 1),
		m.NewPost("three", "three body", 1),
		m.NewPost("four", "four body", 1),
		m.NewPost("five", "five body", 1),
	}

	authors := []*m.Author{
		m.NewAuthor("John", "Doe"),
		m.NewAuthor("Jane", "Doe"),
		m.NewAuthor("James", "Doe"),
	}

	books := []*m.Book{
		m.NewBook("one book", 2010, 1),
		m.NewBook("two book", 2018, 2),
		m.NewBook("three book", 2004, 3),
		m.NewBook("four book", 2014, 1),
	}

	users := []*m.User{
		m.NewUser("John", "Doe", "john_doe", "john@doe.com", "password1"),
		m.NewUser("Jenny", "Doe", "jenny_doe", "jenny@doe.com", "password2"),
		m.NewUser("Jim", "Doe", "jim_doe", "jim@doe.com", "password3"),
	}

	permissions := []*m.Permission{
		pe.AuthLogin,
		pe.AuthResetPassword,

		pe.AuthorAdd,
		pe.AuthorAddAll,
		pe.AuthorEdit,
		pe.AuthorGetAll,
		pe.AuthorGetOne,
		pe.AuthorRemove,

		pe.BookAdd,
		pe.BookAddAll,
		pe.BookEdit,
		pe.BookGetAll,
		pe.BookGetOne,
		pe.BookRemove,

		pe.PermissionAdd,
		pe.PermissionAddAll,
		pe.PermissionEdit,
		pe.PermissionGetAll,
		pe.PermissionGetByEntity,
		pe.PermissionGetOne,
		pe.PermissionRemove,

		pe.PostAdd,
		pe.PostAddAll,
		pe.PostEdit,
		pe.PostGetAll,
		pe.PostGetOne,
		pe.PostRemove,

		pe.RoleAdd,
		pe.RoleAddAll,
		pe.RoleEdit,
		pe.RoleGetAll,
		pe.RoleGetOne,
		pe.RoleRemove,

		pe.UserAdd,
		pe.UserAddAll,
		pe.UserEdit,
		pe.UserGetAll,
		pe.UserGetOne,
		pe.UserRemove,

		pe.PermissionManagementAddPermissionToRole,
		pe.PermissionManagementAddPermissionsToRole,
		pe.PermissionManagementGetPermissionsByRoleId,
		pe.PermissionManagementGetPermissonsByUserId,
		pe.PermissionManagementRemovePermissionFromRole,
		pe.PermissionManagementRemovePermissionsFromRole,

		pe.RoleManagementAddRoleToUser,
		pe.RoleManagementGetRoleByUserId,
		pe.RoleManagementRemoveRoleFromUser,
	}

	return authors, books, posts, users, permissions
}
