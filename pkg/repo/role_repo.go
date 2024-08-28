package repo

import (
	e "app/pkg/errors"
	m "app/pkg/model"
	st "app/pkg/stmt"
	"database/sql"
	"errors"
	"strconv"
	"sync"
)

type RoleRepo interface {
	Add(role *m.Role) error
	AddAll(roles *[]*m.Role) error
	Edit(id int, role *m.Role) error
	GetAll() (*[]m.Role, error)
	GetOne(id int) (*m.Role, error)
	Remove(id int) error
}

type roleRepo struct {
	db  *sql.DB
	rw  *sync.RWMutex
	dbU DBUtil
}

func NewRoleRepo(db *sql.DB, rw *sync.RWMutex, dbU DBUtil) RoleRepo {

	return &roleRepo{db, rw, dbU}
}

// Add
func (p *roleRepo) Add(role *m.Role) error {
	id, err := p.dbU.Transaction(st.ADD_ROLE_STMT, role.Name)
	if err != nil {
		return errors.Join(e.ErrRoleDomain, e.ErrOnAdd, err)
	}

	role.ID = int(id)

	return nil
}

// AddAll
func (p *roleRepo) AddAll(roles *[]*m.Role) error {
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
func (p *roleRepo) Edit(id int, role *m.Role) error {
	idx, err := p.dbU.Transaction(st.EDIT_ROLE_STMT, role.Name, id)
	if err != nil {
		return errors.Join(e.ErrRoleDomain, e.ErrOnEdit, err)
	}

	role.ID = int(idx)

	return nil
}

// GetAll
func (p *roleRepo) GetAll() (*[]m.Role, error) {
	p.rw.RLock()
	defer p.rw.RUnlock()

	role := m.Role{}
	roles := []m.Role{}

	rows, err := p.db.Query(st.GET_ALL_ROLE_STMT)
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
func (p *roleRepo) GetOne(id int) (*m.Role, error) {
	p.rw.RLock()
	defer p.rw.RUnlock()

	role := m.Role{}

	row := p.db.QueryRow(st.GET_ONE_ROLE_STMT, id)
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

	_, err = p.dbU.Transaction(st.REMOVE_ROLE_STMT, id)
	if err != nil {
		return errors.Join(e.ErrRoleDomain, e.ErrOnRemove, err)
	}

	return nil
}
