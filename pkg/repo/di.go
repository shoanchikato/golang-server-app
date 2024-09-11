package repo

import (
	"database/sql"
	"sync"
)

type Repos struct {
	User                 UserRepo
	Auth                 AuthRepo
	Role                 RoleRepo
	Permission           PermissionRepo
	Author               AuthorRepo
	Book                 BookRepo
	Post                 PostRepo
	RoleManagement       RoleManagementRepo
	PermissionManagement PermissionManagementRepo
}

func RepoDi(db *sql.DB, rw *sync.RWMutex, dbU DBUtil) *Repos {
	user := NewUserRepo(db, rw, dbU)
	auth := NewAuthRepo(db, rw, dbU)
	role := NewRoleRepo(db, rw, dbU)
	permission := NewPermissionRepo(db, rw, dbU)
	author := NewAuthorRepo(db, rw, dbU)
	book := NewBookRepo(db, rw, dbU, author)
	post := NewPostRepo(db, rw, dbU)
	roleManagment := NewRoleManagementRepo(db, rw, dbU, user, role, permission)
	permissionManagement := NewPermissionManagementRepo(db, rw, dbU, user, role, permission)

	return &Repos{user, auth, role, permission, author, book, post, roleManagment, permissionManagement}
}
