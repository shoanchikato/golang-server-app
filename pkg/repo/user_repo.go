package repo

import (
	e "app/pkg/errors"
	m "app/pkg/model"
	st "app/pkg/stmt"
	"database/sql"
	"errors"
	"log"
	"strconv"
	"sync"
)

type UserRepo interface {
	Add(user *m.User) error
	AddAll(user *[]*m.User) error
	Edit(id int, user *m.User) error
	GetAll(lastId, limit int) (*[]m.User, error)
	GetOne(id int) (*m.User, error)
	Remove(id int) error
}

type userRepo struct {
	db  *sql.DB
	rw  *sync.RWMutex
	dbU DBUtil
}

func NewUserRepo(db *sql.DB, rw *sync.RWMutex, dbU DBUtil) UserRepo {
	_, err := db.Exec(st.CREATE_USER_TABLE_STMT)
	if err != nil {
		log.Fatalf("error creating post table: %v: %s\n", err, st.CREATE_USER_TABLE_STMT)
	}

	return &userRepo{db, rw, dbU}
}

// Add
func (p *userRepo) Add(user *m.User) error {
	id, err := p.dbU.Transaction(
		st.ADD_USER_STMT,
		user.FirstName,
		user.LastName,
		user.Username,
		user.Email,
		user.Password,
	)
	if err != nil {
		return errors.Join(e.ErrUserDomain, e.ErrOnAdd, err)
	}

	user.Id = int(id)

	return nil
}

// AddAll
func (p *userRepo) AddAll(users *[]*m.User) error {
	newUsers := *users
	for i := 0; i < len(newUsers); i++ {
		user := newUsers[i]
		err := p.Add(user)
		if err != nil {
			return err
		}
	}

	return nil
}

// Edit
func (p *userRepo) Edit(id int, user *m.User) error {
	idx, err := p.dbU.Transaction(
		st.EDIT_USER_STMT,
		user.FirstName,
		user.LastName,
		user.Username,
		user.Email,
		id,
	)
	if err != nil {
		return errors.Join(e.ErrUserDomain, e.ErrOnEdit, err)
	}

	user.Id = int(idx)

	return nil
}

// GetAll
func (p *userRepo) GetAll(lastId, limit int) (*[]m.User, error) {
	p.rw.RLock()
	defer p.rw.RUnlock()

	user := m.User{}
	users := []m.User{}

	rows, err := p.db.Query(st.GET_ALL_USER_STMT, lastId, limit)
	if err != nil {
		return nil, errors.Join(e.ErrUserDomain, e.ErrOnGetAll, e.ErrRepoPreparingStmt, err)
	}
	defer rows.Close()

	for rows.Next() {
		err = rows.Scan(
			&user.Id,
			&user.FirstName,
			&user.LastName,
			&user.Username,
			&user.Email,
		)
		if err != nil {
			return nil, errors.Join(e.ErrUserDomain, e.ErrOnGetAll, e.ErrRepoExecutingStmt, err)
		}

		users = append(users, user)
	}
	err = rows.Err()
	if err != nil {
		return nil, errors.Join(e.ErrUserDomain, e.ErrOnGetAll, e.ErrRepoExecutingStmt, err)
	}

	return &users, nil
}

// GetOne
func (p *userRepo) GetOne(id int) (*m.User, error) {
	p.rw.RLock()
	defer p.rw.RUnlock()

	user := m.User{}

	row := p.db.QueryRow(st.GET_ONE_USER_STMT, id)
	err := row.Scan(
		&user.Id,
		&user.FirstName,
		&user.LastName,
		&user.Username,
		&user.Email,
	)
	if err == sql.ErrNoRows {
		return nil, errors.Join(e.ErrUserDomain, e.ErrRepoExecutingStmt, e.NewErrRepoNotFound(strconv.Itoa(id)))
	}

	if err != nil {
		return nil, errors.Join(e.ErrUserDomain, e.ErrOnGetOne, e.ErrRepoExecutingStmt, err)
	}

	return &user, nil
}

// Remove
func (p *userRepo) Remove(id int) error {
	_, err := p.GetOne(id)
	if err != nil {
		return err
	}

	_, err = p.dbU.Transaction(st.REMOVE_USER_STMT, id)
	if err != nil {
		return errors.Join(e.ErrUserDomain, e.ErrOnRemove, err)
	}

	return nil
}
