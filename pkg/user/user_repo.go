package user

import (
	e "app/pkg/errors"
	r "app/pkg/repo"
	"database/sql"
	"errors"
	"log"
	"strconv"
	"sync"
)

type UserRepo interface {
	Add(user *User) error
	AddAll(user *[]*User) error
	Edit(id int, user *User) error
	GetAll() (*[]User, error)
	GetOne(id int) (*User, error)
	Remove(id int) error
}

type userRepo struct {
	db  *sql.DB
	rw  *sync.RWMutex
	dbU r.DBUtil
}

func NewUserRepo(db *sql.DB, rw *sync.RWMutex, dbU r.DBUtil) UserRepo {
	_, err := db.Exec(CREATE_USER_TABLE_STMT)
	if err != nil {
		log.Fatalf("error creating post table: %v: %s\n", err, CREATE_USER_TABLE_STMT)
	}

	return &userRepo{db, rw, dbU}
}

// Add
func (p *userRepo) Add(user *User) error {
	id, err := p.dbU.Transaction(
		ADD_USER_STMT,
		user.FirstName,
		user.LastName,
		user.Username,
		user.Email,
		user.Password,
	)
	if err != nil {
		return errors.Join(e.ErrUserDomain, e.ErrRepoAdd, err)
	}

	user.ID = int(id)

	return nil
}

// AddAll
func (p *userRepo) AddAll(users *[]*User) error {
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
func (p *userRepo) Edit(id int, user *User) error {
	_, err := p.dbU.Transaction(
		EDIT_USER_STMT,
		user.FirstName,
		user.LastName,
		user.Username,
		user.Email,
		id,
	)
	if err != nil {
		return errors.Join(e.ErrUserDomain, e.ErrRepoEdit, err)
	}

	return nil
}

// GetAll
func (p *userRepo) GetAll() (*[]User, error) {
	p.rw.RLock()
	defer p.rw.RUnlock()

	user := User{}
	users := []User{}

	rows, err := p.db.Query(GET_ALL_USER_STMT)
	if err != nil {
		return nil, errors.Join(e.ErrUserDomain, e.ErrRepoGetAll, e.ErrRepoPreparingStmt, err)
	}
	defer rows.Close()

	for rows.Next() {
		err = rows.Scan(
			&user.ID,
			&user.FirstName,
			&user.LastName,
			&user.Username,
			&user.Email,
		)
		if err != nil {
			return nil, errors.Join(e.ErrUserDomain, e.ErrRepoGetAll, e.ErrRepoExecutingStmt, err)
		}

		users = append(users, user)
	}
	err = rows.Err()
	if err != nil {
		return nil, errors.Join(e.ErrUserDomain, e.ErrRepoGetAll, e.ErrRepoExecutingStmt, err)
	}

	return &users, nil
}

// GetOne
func (p *userRepo) GetOne(id int) (*User, error) {
	p.rw.RLock()
	defer p.rw.RUnlock()

	user := User{}

	row := p.db.QueryRow(GET_ONE_USER_STMT, id)
	err := row.Scan(
		&user.ID,
		&user.FirstName,
		&user.LastName,
		&user.Username,
		&user.Email,
	)
	if err == sql.ErrNoRows {
		return nil, errors.Join(e.ErrUserDomain, e.ErrRepoExecutingStmt, e.NewErrRepoNotFound(strconv.Itoa(id)))
	}

	if err != nil {
		return nil, errors.Join(e.ErrUserDomain, e.ErrRepoGetOne, e.ErrRepoExecutingStmt, err)
	}

	return &user, nil
}

// Remove
func (p *userRepo) Remove(id int) error {
	_, err := p.GetOne(id)
	if err != nil {
		return err
	}

	_, err = p.dbU.Transaction(REMOVE_USER_STMT, id)
	if err != nil {
		return errors.Join(e.ErrUserDomain, e.ErrRepoRemove, err)
	}

	return nil
}
