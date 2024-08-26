package permission

import (
	e "app/pkg/errors"
	r "app/pkg/repo"
	"database/sql"
	"errors"
	"log"
	"strconv"
	"sync"
)

type PermissionRepo interface {
	Add(permission *Permission) error
	AddAll(permissions *[]*Permission) error
	Edit(id int, permission *Permission) error
	GetAll() (*[]Permission, error)
	GetOne(id int) (*Permission, error)
	Remove(id int) error
}

type permissionRepo struct {
	db  *sql.DB
	rw  *sync.RWMutex
	dbU r.DBUtil
}

func NewPermissionRepo(db *sql.DB, rw *sync.RWMutex, dbU r.DBUtil) PermissionRepo {
	_, err := db.Exec(CREATE_PERMISSION_TABLE_STMT)
	if err != nil {
		log.Fatalf("error creating permission table: %v: %s\n", err, CREATE_PERMISSION_TABLE_STMT)
	}

	return &permissionRepo{db, rw, dbU}
}

// Add
func (p *permissionRepo) Add(permission *Permission) error {
	id, err := p.dbU.Transaction(ADD_PERMISSION_STMT, permission.Name)
	if err != nil {
		return errors.Join(e.ErrPermissionDomain, e.ErrRepoAdd, err)
	}

	permission.ID = int(id)

	return nil
}

// AddAll
func (p *permissionRepo) AddAll(permissions *[]*Permission) error {
	newPermissions := *permissions
	for i := 0; i < len(newPermissions); i++ {
		permission := newPermissions[i]
		err := p.Add(permission)
		if err != nil {
			return err
		}
	}

	return nil
}

// Edit
func (p *permissionRepo) Edit(id int, permission *Permission) error {
	_, err := p.dbU.Transaction(EDIT_PERMISSION_STMT, permission.Name, id)
	if err != nil {
		return errors.Join(e.ErrPermissionDomain, e.ErrRepoEdit, err)
	}

	return nil
}

// GetAll
func (p *permissionRepo) GetAll() (*[]Permission, error) {
	p.rw.RLock()
	defer p.rw.RUnlock()

	permission := Permission{}
	permissions := []Permission{}

	rows, err := p.db.Query(GET_ALL_PERMISSION_STMT)
	if err != nil {
		return nil, errors.Join(e.ErrPermissionDomain, e.ErrRepoGetAll, e.ErrRepoPreparingStmt, err)
	}
	defer rows.Close()

	for rows.Next() {
		err = rows.Scan(&permission.ID, &permission.Name)
		if err != nil {
			return nil, errors.Join(e.ErrPermissionDomain, e.ErrRepoGetAll, e.ErrRepoExecutingStmt, err)
		}

		permissions = append(permissions, permission)
	}
	err = rows.Err()
	if err != nil {
		return nil, errors.Join(e.ErrPermissionDomain, e.ErrRepoGetAll, e.ErrRepoExecutingStmt, err)
	}

	return &permissions, nil
}

// GetOne
func (p *permissionRepo) GetOne(id int) (*Permission, error) {
	p.rw.RLock()
	defer p.rw.RUnlock()

	permission := Permission{}

	row := p.db.QueryRow(GET_ONE_PERMISSION_STMT, id)
	err := row.Scan(&permission.ID, &permission.Name)
	if err == sql.ErrNoRows {
		return nil, errors.Join(e.ErrPermissionDomain, e.ErrRepoExecutingStmt, e.NewErrRepoNotFound(strconv.Itoa(id)))
	}

	if err != nil {
		return nil, errors.Join(e.ErrPermissionDomain, e.ErrRepoGetOne, e.ErrRepoExecutingStmt, err)
	}

	return &permission, nil
}

// Remove
func (p *permissionRepo) Remove(id int) error {
	_, err := p.GetOne(id)
	if err != nil {
		return err
	}
	_, err = p.dbU.Transaction(REMOVE_PERMISSION_STMT, id)
	if err != nil {
		return errors.Join(e.ErrPermissionDomain, e.ErrRepoRemove, err)
	}

	return nil
}
