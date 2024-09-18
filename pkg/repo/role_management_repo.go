package repo

import (
	e "app/pkg/errors"
	m "app/pkg/model"
	st "app/pkg/stmt"
	"database/sql"
	"errors"
	"sync"
)

type RoleManagementRepo interface {
	AddRoleToUser(roleId, userId int) error
	GetRolesByUserId(userId int) (*[]m.Role, error)
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
	_, err := r.ur.GetOne(userId)
	if err != nil {
		return err
	}

	_, err = r.rr.GetOne(roleId)
	if err != nil {
		return err
	}

	roles, err := r.GetRolesByUserId(userId)
	if err != nil {
		return err
	}

	for _, role := range *roles {
		if role.Id == roleId {
			return errors.Join(e.ErrRoleManagementDomain, e.ErrOnAddRoleToUser, e.ErrRepoUserAlreadyHasRole)
		}
	}

	_, err = r.dbU.Transaction(st.ADD_ROLE_TO_USER_STMT, roleId, userId)
	if err != nil {
		return errors.Join(e.ErrRoleManagementDomain, e.ErrOnAddRoleToUser, err)
	}

	return nil
}

// GetRolesByUserId
func (r *rMRepo) GetRolesByUserId(userId int) (*[]m.Role, error) {
	_, err := r.ur.GetOne(userId)
	if err != nil {
		return nil, err
	}

	r.rw.RLock()
	defer r.rw.RUnlock()

	roles := []m.Role{}
	role := m.Role{}

	rows, err := r.db.Query(st.GET_ROLES_BY_USER_ID_STMT, userId)
	if err != nil {
		return nil, errors.Join(e.ErrRoleManagementDomain, e.ErrOnGetRolesByUserId, e.ErrRepoPreparingStmt, err)
	}
	defer rows.Close()

	for rows.Next() {
		err = rows.Scan(&role.Id, &role.Name)
		if err != nil {
			return nil, errors.Join(e.ErrPostDomain, e.ErrOnGetAll, e.ErrRepoExecutingStmt, err)
		}

		roles = append(roles, role)
	}

	if err != nil {
		return nil, errors.Join(e.ErrRoleManagementDomain, e.ErrOnGetRolesByUserId, e.ErrRepoExecutingStmt, err)
	}

	return &roles, nil
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

	_, err = r.GetRolesByUserId(userId)
	if err != nil {
		return err
	}

	_, err = r.dbU.Transaction(st.REMOVE_ROLE_FROM_USER_STMT, roleId, userId)
	if err != nil {
		return errors.Join(e.ErrRoleManagementDomain, e.ErrOnRemoveRoleFromUser, err)
	}

	return nil
}
