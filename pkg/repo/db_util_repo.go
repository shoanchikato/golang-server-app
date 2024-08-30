package repo

import (
	c "app/pkg/constants"
	e "app/pkg/errors"
	"database/sql"
	"errors"
	"strings"
	"sync"

	"github.com/mattn/go-sqlite3"
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
	if *limit > c.REPO_UPPER_LIMIT {
		*limit = c.REPO_UPPER_LIMIT
	}

	if *limit < c.REPO_LOWER_LIMIT {
		*limit = c.REPO_LOWER_LIMIT
	}
}

func (d *dbUtil) getField(err error) string {
	firstPart := strings.SplitAfter(err.Error(), ": ")[1]
	secondPart := strings.SplitAfter(firstPart, ".")[1]

	return secondPart
}

func (d *dbUtil) getDuplicateColumn(err *sqlite3.Error) error {
	if err.ExtendedCode == sqlite3.ErrConstraintUnique {
		field := d.getField(err)
		return e.NewErrRepoDuplicate(field)
	}

	return err
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
	perr := &sqlite3.Error{}
	if errors.As(err, perr) {
		err = d.getDuplicateColumn(perr)
		return 0, err
	}

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
