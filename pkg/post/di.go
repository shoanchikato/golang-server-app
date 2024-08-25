package post

import (
	r "app/pkg/repo"
	"database/sql"
	"sync"
)

func Di(
	db *sql.DB,
	rw *sync.RWMutex,
	dbU r.DBUtil,
) PostValidator {
	repo := NewPostRepo(db, rw, dbU)
	val := NewPostValidator(repo)

	return val
}
