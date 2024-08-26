package auth

import (
	e "app/pkg/errors"
	r "app/pkg/repo"
	"database/sql"
	"errors"
	"sync"
)

type AuthRepo interface {
	GetByUsername(username string) (*Auth, error)
	ResetPassword(username, newPassword string) error
}

type authRepo struct {
	db  *sql.DB
	rw  *sync.RWMutex
	dbU r.DBUtil
}

func NewAuthRepo(db *sql.DB, rw *sync.RWMutex, dbU r.DBUtil) AuthRepo {
	return &authRepo{db, rw, dbU}
}

// GetByUsername
func (a *authRepo) GetByUsername(username string) (*Auth, error) {
	a.rw.RLock()
	defer a.rw.RUnlock()

	auth := Auth{}

	row := a.db.QueryRow(GET_AUTH_DETAILS_BY_USERNAME, username)
	err := row.Scan(&auth.Username, &auth.Email, &auth.Password, &auth.UserID)
	if err == sql.ErrNoRows {
		return nil, errors.Join(e.ErrAuthDomain, e.ErrRepoExecutingStmt, e.NewErrRepoNotFound(username))
	}

	if err != nil {
		return nil, errors.Join(e.ErrAuthDomain, e.ErrRepoGetOne, e.ErrRepoExecutingStmt, err)
	}

	return &auth, nil
}

// ResetPassword
func (a *authRepo) ResetPassword(username string, newPassword string) error {
	_, err := a.dbU.Transaction(RESET_PASSWORD, newPassword, username)
	if err != nil {
		return err
	}

	return nil
}
