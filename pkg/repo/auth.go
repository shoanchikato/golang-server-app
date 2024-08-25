package repo

import (
	e "app/pkg/errors"
	m "app/pkg/model"
	s "app/pkg/repo/stmt"
	"database/sql"
	"errors"
	"sync"
)

type AuthRepo interface {
	GetByUsername(username string) (*m.Auth, error)
	ResetPassword(username, newPassword string) error
}

type authRepo struct {
	db  *sql.DB
	rw  *sync.RWMutex
	dbU DBUtil
}

func NewAuthRepo(db *sql.DB, rw *sync.RWMutex, dbU DBUtil) AuthRepo {
	return &authRepo{db, rw, dbU}
}

// GetByUsername
func (a *authRepo) GetByUsername(username string) (*m.Auth, error) {
	a.rw.RLock()
	defer a.rw.RUnlock()

	auth := m.Auth{}

	row := a.db.QueryRow(s.GET_AUTH_DETAILS_BY_USERNAME, username)
	err := row.Scan(&auth.Username, &auth.Email, &auth.Password, &auth.UserID)
	if err == sql.ErrNoRows {
		return nil, errors.Join(e.ErrRepoExecutingStmt, e.NewErrRepoNotFound(username))
	}

	if err != nil {
		return nil, errors.Join(e.ErrRepoGetOne, e.ErrRepoExecutingStmt, err)
	}

	return &auth, nil
}

// ResetPassword
func (a *authRepo) ResetPassword(username string, newPassword string) error {
	_, err := a.dbU.Transaction(s.RESET_PASSWORD, newPassword, username)
	if err != nil {
		return err
	}

	return nil
}
