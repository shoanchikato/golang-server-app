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

type RoleManagementRepo interface {
	AddRoleToUser(roleId, userId int) error
	GetRoleByUserId(userId int) (*m.Role, error)
	RemoveRoleFromUser(roleId int, permissionId int) error
}

type rMRepo struct {
	db  *sql.DB
	rw  *sync.RWMutex
	dbU DBUtil
	ur  UserRepo
	rr  RoleRepo
	pr  PermissionRepo
}

func NewRoleManagementRepo(
	db *sql.DB,
	rw *sync.RWMutex,
	dbU DBUtil,
	ur UserRepo,
	rr RoleRepo,
	pr PermissionRepo,
) RoleManagementRepo {
	return &rMRepo{db, rw, dbU, ur, rr, pr}
}

// AddRoleToUser
func (r *rMRepo) AddRoleToUser(roleId, userId int) error {
	role, _ := r.GetRoleByUserId(userId)
	if role != nil {
		return nil
	}

	_, err := r.ur.GetOne(userId)
	if err != nil {
		return err
	}

	_, err = r.rr.GetOne(roleId)
	if err != nil {
		return err
	}

	_, err = r.dbU.Transaction(st.ADD_ROLE_TO_USER_STMT, roleId, userId)
	if err != nil {
		return errors.Join(e.ErrRoleManagementDomain, e.ErrOnAddRoleToUser, err)
	}

	return nil
}

// GetRoleByUserId
func (r *rMRepo) GetRoleByUserId(userId int) (*m.Role, error) {
	r.rw.RLock()
	defer r.rw.RUnlock()

	role := m.Role{}

	row := r.db.QueryRow(st.GET_ROLE_BY_USER_ID_STMT, userId)
	err := row.Scan(&role.Id, &role.Name)
	if err == sql.ErrNoRows {
		return nil, errors.Join(
			e.ErrRoleManagementDomain,
			e.ErrRepoExecutingStmt,
			e.NewErrRepoNotFound("user id", strconv.Itoa(userId)),
		)
	}

	if err != nil {
		return nil, errors.Join(e.ErrRoleManagementDomain, e.ErrOnGetRoleByUserId, e.ErrRepoExecutingStmt, err)
	}

	return &role, nil
}

// RemoveRoleFromUser
func (r *rMRepo) RemoveRoleFromUser(roleId int, userId int) error {
	_, err := r.ur.GetOne(userId)
	if err != nil {
		return err
	}

	_, err = r.rr.GetOne(roleId)
	if err != nil {
		return err
	}

	_, err = r.GetRoleByUserId(userId)
	if err != nil {
		return err
	}

	_, err = r.dbU.Transaction(st.REMOVE_ROLE_FROM_USER_STMT, roleId, userId)
	if err != nil {
		return errors.Join(e.ErrRoleManagementDomain, e.ErrOnRemoveRoleFromUser, err)
	}

	return nil
}
