package book

import (
	r "app/pkg/repo"
	"database/sql"
	"sync"
)

func Di(
	db *sql.DB,
	rw *sync.RWMutex,
	dbU r.DBUtil,
) BookValidator {
	repo := NewBookRepo(db, rw, dbU)
	val := NewBookValidator(repo)

	return val
}
