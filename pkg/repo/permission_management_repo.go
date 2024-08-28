package repo

import (
	e "app/pkg/errors"
	st "app/pkg/stmt"
	m "app/pkg/model"
	"database/sql"
	"errors"
	"strconv"
	"sync"
)

type PermissionManagementRepo interface {
	AddPermissionToRole(permissionID, roleID int) error
	AddPermissionsToRole(permissionIDs []int, roleID int) error
	AddRoleToUser(roleID, userID int) error
	GetPermissionsByRoleID(roleID int) (*[]m.Permission, error)
	GetPermissonsByUserID(userID int) (*[]m.Permission, error)
	GetRoleByUserID(userID int) (*m.Role, error)
	RemovePermissionFromRole(roleID, permissionID int) error
	RemovePermissionsFromRole(roleID int, permissionIDs []int) error
	RemoveRoleFromUser(roleID int, permissionID int) error
}

type pMRepo struct {
	db  *sql.DB
	rw  *sync.RWMutex
	dbU DBUtil
	ur  UserRepo
	rr  RoleRepo
	pr  PermissionRepo
}

func NewPermissionManagementRepo(
	db *sql.DB,
	rw *sync.RWMutex,
	dbU DBUtil,
	ur UserRepo,
	rr RoleRepo,
	pr PermissionRepo,
) PermissionManagementRepo {
	return &pMRepo{db, rw, dbU, ur, rr, pr}
}

// AddPermissionToRole
func (p *pMRepo) AddPermissionToRole(permissionID, roleID int) error {
	permissions, _ := p.GetPermissionsByRoleID(roleID)
	if permissions != nil {
		pp := *permissions
		for i := 0; i < len(pp); i++ {
			if permissionID == pp[i].ID {
				return nil
			}
		}
	}

	_, err := p.rr.GetOne(roleID)
	if err != nil {
		return err
	}

	_, err = p.pr.GetOne(permissionID)
	if err != nil {
		return err
	}

	_, err = p.dbU.Transaction(st.ADD_PERMISSION_TO_ROLE_STMT, permissionID, roleID)
	if err != nil {
		return errors.Join(e.ErrPermissionManagementDomain, e.ErrOnAdd, err)
	}

	return nil
}

// AddPermissionsToRole
func (p *pMRepo) AddPermissionsToRole(permissionIDs []int, roleID int) error {
	for i := 0; i < len(permissionIDs); i++ {
		permissionID := permissionIDs[i]
		err := p.AddPermissionToRole(permissionID, roleID)
		if err != nil {
			return err
		}
	}

	return nil
}

// AddRoleToUser
func (p *pMRepo) AddRoleToUser(roleID, userID int) error {
	role, _ := p.GetRoleByUserID(userID)
	if role != nil {
		return nil
	}

	_, err := p.ur.GetOne(userID)
	if err != nil {
		return err
	}

	_, err = p.rr.GetOne(roleID)
	if err != nil {
		return err
	}

	_, err = p.dbU.Transaction(st.ADD_ROLE_TO_USER_STMT, roleID, userID)
	if err != nil {
		return errors.Join(e.ErrPermissionManagementDomain, e.ErrOnAdd, err)
	}

	return nil
}

// GetPermissionsByRoleID
func (p *pMRepo) GetPermissionsByRoleID(roleID int) (*[]m.Permission, error) {
	p.rw.RLock()
	defer p.rw.RUnlock()

	_, err := p.rr.GetOne(roleID)
	if err != nil {
		return nil, err
	}

	permission := m.Permission{}
	permissions := []m.Permission{}

	rows, err := p.db.Query(st.GET_PERMISSIONS_BY_ROLE_ID_STMT, roleID)
	if err != nil {
		return nil, errors.Join(e.ErrPermissionManagementDomain, e.ErrOnGetAll, e.ErrRepoPreparingStmt, err)
	}
	defer rows.Close()

	for rows.Next() {
		err = rows.Scan(&permission.ID, &permission.Name)
		if err != nil {
			return nil, errors.Join(e.ErrPermissionManagementDomain, e.ErrOnGetAll, e.ErrRepoExecutingStmt, err)
		}

		permissions = append(permissions, permission)
	}
	err = rows.Err()
	if err != nil {
		return nil, errors.Join(e.ErrPermissionManagementDomain, e.ErrOnGetAll, e.ErrRepoExecutingStmt, err)
	}

	return &permissions, nil
}

// GetPermissonsByUserID
func (p *pMRepo) GetPermissonsByUserID(userID int) (*[]m.Permission, error) {
	p.rw.RLock()
	defer p.rw.RUnlock()

	_, err := p.ur.GetOne(userID)
	if err != nil {
		return nil, err
	}

	permission := m.Permission{}
	permissions := []m.Permission{}

	rows, err := p.db.Query(st.GET_PERMISSIONS_BY_USER_ID, userID)
	if err != nil {
		return nil, errors.Join(e.ErrPermissionManagementDomain, e.ErrOnGetAll, e.ErrRepoPreparingStmt, err)
	}
	defer rows.Close()

	for rows.Next() {
		err = rows.Scan(&permission.ID, &permission.Name)
		if err != nil {
			return nil, errors.Join(e.ErrPermissionManagementDomain, e.ErrOnGetAll, e.ErrRepoExecutingStmt, err)
		}

		permissions = append(permissions, permission)
	}
	err = rows.Err()
	if err != nil {
		return nil, errors.Join(e.ErrPermissionManagementDomain, e.ErrOnGetAll, e.ErrRepoExecutingStmt, err)
	}

	return &permissions, nil
}

// GetRoleByUserID
func (p *pMRepo) GetRoleByUserID(userID int) (*m.Role, error) {
	p.rw.RLock()
	defer p.rw.RUnlock()

	role := m.Role{}

	row := p.db.QueryRow(st.GET_ROLE_BY_USER_ID_STMT, userID)
	err := row.Scan(&role.ID, &role.Name)
	if err == sql.ErrNoRows {
		return nil, errors.Join(
			e.ErrPermissionManagementDomain,
			e.ErrRepoExecutingStmt,
			e.NewErrRepoNotFound(strconv.Itoa(userID)),
		)
	}

	if err != nil {
		return nil, errors.Join(e.ErrPermissionManagementDomain, e.ErrOnGetOne, e.ErrRepoExecutingStmt, err)
	}

	return &role, nil
}

// RemovePermissionFromRole
func (p *pMRepo) RemovePermissionFromRole(roleID int, permissionID int) error {
	_, err := p.pr.GetOne(permissionID)
	if err != nil {
		return err
	}

	_, err = p.rr.GetOne(roleID)
	if err != nil {
		return err
	}

	_, err = p.GetPermissionsByRoleID(roleID)
	if err != nil {
		return err
	}

	_, err = p.dbU.Transaction(st.REMOVE_PERMISSION_FROM_ROLE_STMT, roleID, permissionID)
	if err != nil {
		return errors.Join(e.ErrPermissionManagementDomain, e.ErrOnRemove, err)
	}

	return nil
}

// RemovePermissionsFromRole
func (p *pMRepo) RemovePermissionsFromRole(roleID int, permissions []int) error {
	for i := 0; i < len(permissions); i++ {
		permissionID := permissions[i]
		err := p.RemovePermissionFromRole(roleID, permissionID)
		if err != nil {
			return err
		}
	}

	return nil
}

// RemoveRoleFromUser
func (p *pMRepo) RemoveRoleFromUser(roleID int, userID int) error {
	_, err := p.ur.GetOne(userID)
	if err != nil {
		return err
	}

	_, err = p.rr.GetOne(roleID)
	if err != nil {
		return err
	}

	_, err = p.GetRoleByUserID(userID)
	if err != nil {
		return err
	}

	_, err = p.dbU.Transaction(st.REMOVE_ROLE_FROM_USER_STMT, roleID, userID)
	if err != nil {
		return errors.Join(e.ErrPermissionManagementDomain, e.ErrOnRemove, err)
	}

	return nil
}
