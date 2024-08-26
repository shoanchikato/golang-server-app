package permission

import (
	r "app/pkg/repo"
	rr "app/pkg/role"
	u "app/pkg/user"
	"database/sql"
	"sync"
)

func Di(
	db *sql.DB,
	rw *sync.RWMutex,
	dbU r.DBUtil,
	ur u.UserRepo,
	rr rr.RoleRepo,
) (
	PermissionValidator,
	PermissionManagementValidator,
) {
	pr := NewPermissionRepo(db, rw, dbU)
	pmr := NewPermissonManagementRepo(db, rw, dbU, ur, rr, pr)

	permissionVal := NewPermissionValidator(pr)
	permissionMngVal := NewPermissionManagementValidator(pmr)

	return permissionVal, permissionMngVal
}
