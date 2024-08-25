package author

import (
	r "app/pkg/repo"
	"database/sql"
	"sync"
)

func Di(
	db *sql.DB,
	rw *sync.RWMutex,
	dbU r.DBUtil,
) AuthorValidator {
	repo := NewAuthorRepo(db, rw, dbU)
	val := NewAuthorValidator(repo)

	return val
}
