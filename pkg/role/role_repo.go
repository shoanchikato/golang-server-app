package role

import (
	e "app/pkg/errors"
	r "app/pkg/repo"
	"database/sql"
	"errors"
	"strconv"
	"sync"
)

type RoleRepo interface {
	Add(role *Role) error
	AddAll(roles *[]*Role) error
	Edit(id int, role *Role) error
	GetAll() (*[]Role, error)
	GetOne(id int) (*Role, error)
	Remove(id int) error
}

type roleRepo struct {
	db  *sql.DB
	rw  *sync.RWMutex
	dbU r.DBUtil
}

func NewRoleRepo(db *sql.DB, rw *sync.RWMutex, dbU r.DBUtil) RoleRepo {

	return &roleRepo{db, rw, dbU}
}

// Add
func (p *roleRepo) Add(role *Role) error {
	id, err := p.dbU.Transaction(ADD_ROLE_STMT, role.Name)
	if err != nil {
		return errors.Join(e.ErrRoleDomain, e.ErrOnAdd, err)
	}

	role.ID = int(id)

	return nil
}

// AddAll
func (p *roleRepo) AddAll(roles *[]*Role) error {
	newRoles := *roles
	for i := 0; i < len(newRoles); i++ {
		role := newRoles[i]
		err := p.Add(role)
		if err != nil {
			return err
		}
	}

	return nil
}

// Edit
func (p *roleRepo) Edit(id int, role *Role) error {
	idx, err := p.dbU.Transaction(EDIT_ROLE_STMT, role.Name, id)
	if err != nil {
		return errors.Join(e.ErrRoleDomain, e.ErrOnEdit, err)
	}

	role.ID = int(idx)

	return nil
}

// GetAll
func (p *roleRepo) GetAll() (*[]Role, error) {
	p.rw.RLock()
	defer p.rw.RUnlock()

	role := Role{}
	roles := []Role{}

	rows, err := p.db.Query(GET_ALL_ROLE_STMT)
	if err != nil {
		return nil, errors.Join(e.ErrRoleDomain, e.ErrOnGetAll, e.ErrRepoPreparingStmt, err)
	}
	defer rows.Close()

	for rows.Next() {
		err = rows.Scan(&role.ID, &role.Name)
		if err != nil {
			return nil, errors.Join(e.ErrRoleDomain, e.ErrOnGetAll, e.ErrRepoExecutingStmt, err)
		}

		roles = append(roles, role)
	}
	err = rows.Err()
	if err != nil {
		return nil, errors.Join(e.ErrRoleDomain, e.ErrOnGetAll, e.ErrRepoExecutingStmt, err)
	}

	return &roles, nil
}

// GetOne
func (p *roleRepo) GetOne(id int) (*Role, error) {
	p.rw.RLock()
	defer p.rw.RUnlock()

	role := Role{}

	row := p.db.QueryRow(GET_ONE_ROLE_STMT, id)
	err := row.Scan(&role.ID, &role.Name)
	if err == sql.ErrNoRows {
		return nil, errors.Join(e.ErrRoleDomain, e.ErrRepoExecutingStmt, e.NewErrRepoNotFound(strconv.Itoa(id)))
	}

	if err != nil {
		return nil, errors.Join(e.ErrRoleDomain, e.ErrOnGetOne, e.ErrRepoExecutingStmt, err)
	}

	return &role, nil
}

// Remove
func (p *roleRepo) Remove(id int) error {
	_, err := p.GetOne(id)
	if err != nil {
		return err
	}

	_, err = p.dbU.Transaction(REMOVE_ROLE_STMT, id)
	if err != nil {
		return errors.Join(e.ErrRoleDomain, e.ErrOnRemove, err)
	}

	return nil
}
