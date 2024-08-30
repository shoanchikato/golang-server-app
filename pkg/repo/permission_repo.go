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

type PermissionRepo interface {
	Add(permission *m.Permission) error
	AddAll(permissions *[]*m.Permission) error
	Edit(id int, permission *m.Permission) error
	GetAll(lastId, limit int) (*[]m.Permission, error)
	GetByEntity(entity string) (*[]m.Permission, error)
	GetOne(id int) (*m.Permission, error)
	Remove(id int) error
}

type permissionRepo struct {
	db  *sql.DB
	rw  *sync.RWMutex
	dbU DBUtil
}

func NewPermissionRepo(db *sql.DB, rw *sync.RWMutex, dbU DBUtil) PermissionRepo {
	_, err := db.Exec(st.CREATE_PERMISSION_TABLE_STMT)
	if err != nil {
		log.Fatalf("error creating permission table: %v: %s\n", err, st.CREATE_PERMISSION_TABLE_STMT)
	}

	return &permissionRepo{db, rw, dbU}
}

// Add
func (p *permissionRepo) Add(permission *m.Permission) error {
	id, err := p.dbU.Transaction(st.ADD_PERMISSION_STMT, permission.Name, permission.Entity, permission.Operation)
	if err != nil {
		return errors.Join(e.ErrPermissionDomain, e.ErrOnAdd, err)
	}

	permission.Id = int(id)

	return nil
}

// AddAll
func (p *permissionRepo) AddAll(permissions *[]*m.Permission) error {
	for _, permission := range *permissions {
		err := p.Add(permission)
		if err != nil {
			return err
		}
	}

	return nil
}

// Edit
func (p *permissionRepo) Edit(id int, permission *m.Permission) error {
	idx, err := p.dbU.Transaction(st.EDIT_PERMISSION_STMT, permission.Name, permission.Entity, permission.Operation, id)
	if err != nil {
		return errors.Join(e.ErrPermissionDomain, e.ErrOnEdit, err)
	}

	permission.Id = int(idx)

	return nil
}

// GetAll
func (p *permissionRepo) GetAll(lastId, limit int) (*[]m.Permission, error) {
	p.rw.RLock()
	defer p.rw.RUnlock()

	p.dbU.CheckLimit(&limit)

	rows, err := p.db.Query(st.GET_ALL_PERMISSION_STMT, lastId, limit)
	if err != nil {
		return nil, errors.Join(e.ErrPermissionDomain, e.ErrOnGetAll, e.ErrRepoPreparingStmt, err)
	}
	defer rows.Close()

	permission := m.Permission{}
	permissions := []m.Permission{}

	for rows.Next() {
		err = rows.Scan(&permission.Id, &permission.Name, &permission.Entity, &permission.Operation)
		if err != nil {
			return nil, errors.Join(e.ErrPermissionDomain, e.ErrOnGetAll, e.ErrRepoExecutingStmt, err)
		}

		permissions = append(permissions, permission)
	}
	err = rows.Err()
	if err != nil {
		return nil, errors.Join(e.ErrPermissionDomain, e.ErrOnGetAll, e.ErrRepoExecutingStmt, err)
	}

	return &permissions, nil
}

// GetAll
func (p *permissionRepo) GetByEntity(entity string) (*[]m.Permission, error) {
	p.rw.RLock()
	defer p.rw.RUnlock()

	rows, err := p.db.Query(st.GET_BY_ENTITY_PERMISSION_STMT, entity)
	if err != nil {
		return nil, errors.Join(e.ErrPermissionDomain, e.ErrOnGetByEntity, e.ErrRepoPreparingStmt, err)
	}
	defer rows.Close()

	permission := m.Permission{}
	permissions := []m.Permission{}

	for rows.Next() {
		err = rows.Scan(&permission.Id, &permission.Name, &permission.Entity, &permission.Operation)
		if err != nil {
			return nil, errors.Join(e.ErrPermissionDomain, e.ErrOnGetByEntity, e.ErrRepoExecutingStmt, err)
		}

		permissions = append(permissions, permission)
	}
	err = rows.Err()
	if err != nil {
		return nil, errors.Join(e.ErrPermissionDomain, e.ErrOnGetByEntity, e.ErrRepoExecutingStmt, err)
	}

	return &permissions, nil
}

// GetOne
func (p *permissionRepo) GetOne(id int) (*m.Permission, error) {
	p.rw.RLock()
	defer p.rw.RUnlock()

	permission := m.Permission{}

	row := p.db.QueryRow(st.GET_ONE_PERMISSION_STMT, id)
	err := row.Scan(&permission.Id, &permission.Name, &permission.Entity, &permission.Operation)
	if err == sql.ErrNoRows {
		return nil, errors.Join(e.ErrPermissionDomain, e.ErrRepoExecutingStmt, e.NewErrRepoNotFound("permission id", strconv.Itoa(id)))
	}

	if err != nil {
		return nil, errors.Join(e.ErrPermissionDomain, e.ErrOnGetOne, e.ErrRepoExecutingStmt, err)
	}

	return &permission, nil
}

// Remove
func (p *permissionRepo) Remove(id int) error {
	_, err := p.GetOne(id)
	if err != nil {
		return err
	}
	_, err = p.dbU.Transaction(st.REMOVE_PERMISSION_STMT, id, id)
	if err != nil {
		return errors.Join(e.ErrPermissionDomain, e.ErrOnRemove, err)
	}

	return nil
}
