package repo

import (
	e "app/pkg/errors"
	"database/sql"
	"errors"
	"sync"
)

type DBUtil interface {
	Transaction(statement string, args ...any) (int64, error)
	CheckLimit(limit *int)
}

type dbUtil struct {
	db *sql.DB
	rw *sync.RWMutex
}

func NewDBUtil(db *sql.DB, rw *sync.RWMutex) DBUtil {
	return &dbUtil{db, rw}
}

// CheckLimit
func (d *dbUtil) CheckLimit(limit *int) {
	if *limit > UPPER_LIMIT {
		*limit = UPPER_LIMIT
	}

	if *limit < LOWER_LIMIT {
		*limit = LOWER_LIMIT
	}
}

// Transaction
func (d *dbUtil) Transaction(statement string, args ...any) (int64, error) {
	d.rw.Lock()
	defer d.rw.Unlock()

	tx, err := d.db.Begin()
	if err != nil {
		tx.Rollback()
		return 0, errors.Join(e.ErrRepoBeginTx, err)
	}

	result, err := tx.Exec(statement, args...)
	if err != nil {
		tx.Rollback()
		return 0, errors.Join(e.ErrRepoExecutingStmt, err)
	}

	err = tx.Commit()
	if err != nil {
		tx.Rollback()
		return 0, errors.Join(e.ErrRepoCommitTx, err)
	}

	id, err := result.LastInsertId()
	if err != nil {
		return id, errors.Join(e.ErrRepoLastInsertedId, err)
	}

	return id, nil
}
