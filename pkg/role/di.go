package role

import (
	r "app/pkg/repo"
	"database/sql"
	"sync"
)

func Di(
	db *sql.DB,
	rw *sync.RWMutex,
	dbU r.DBUtil,
) (
	RoleValidator,
	RoleRepo,
) {
	repo := NewRoleRepo(db, rw, dbU)
	val := NewRoleValidator(repo)

	return val, repo
}
