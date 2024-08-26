package permission

import (
	e "app/pkg/errors"
	r "app/pkg/repo"
	rr "app/pkg/role"
	u "app/pkg/user"
	"database/sql"
	"errors"
	"sync"
)

type PermissionManagementRepo interface {
	AddPermissionToRole(permission *Permission, roleID int) error
	AddPermissionsToRole(permissions *[]*Permission, roleID int) error
	AddRoleToUser(roleID, userID int) error
	GetPermissionsByRoleID(roleID int) (*[]Permission, error)
	GetPermissonsByUserID(userID int) (*[]Permission, error)
	RemovePermissionFromRole(roleID, permissionID int) error
	RemovePermissionsFromRole(roleID int, permissionIDs []int) error
}

type pMRepo struct {
	db  *sql.DB
	rw  *sync.RWMutex
	dbU r.DBUtil
	ur  u.UserRepo
	rr  rr.RoleRepo
	pr  PermissionRepo
}

func NewPermissonManagementRepo(
	db *sql.DB,
	rw *sync.RWMutex,
	dbU r.DBUtil,
	ur u.UserRepo,
	rr rr.RoleRepo,
	pr PermissionRepo,
) PermissionManagementRepo {
	return &pMRepo{db, rw, dbU, ur, rr, pr}
}

// AddPermissionToRole
func (p *pMRepo) AddPermissionToRole(permission *Permission, roleID int) error {
	_, err := p.rr.GetOne(roleID)
	if err != nil {
		return err
	}

	_, err = p.pr.GetOne(permission.ID)
	if err != nil {
		return err
	}

	id, err := p.dbU.Transaction(ADD_PERMISSION_TO_ROLE_STMT, permission.ID, roleID)
	if err != nil {
		return errors.Join(e.ErrPermissionManagement, e.ErrRepoAdd, err)
	}

	permission.ID = int(id)

	return nil
}

// AddPermissionsToRole
func (p *pMRepo) AddPermissionsToRole(permissions *[]*Permission, roleID int) error {
	newPermissions := *permissions
	for i := 0; i < len(newPermissions); i++ {
		permission := newPermissions[i]
		err := p.AddPermissionToRole(permission, roleID)
		if err != nil {
			return err
		}
	}

	return nil
}

// AddRoleToUser
func (p *pMRepo) AddRoleToUser(roleID, userID int) error {
	_, err := p.ur.GetOne(userID)
	if err != nil {
		return err
	}

	_, err = p.rr.GetOne(roleID)
	if err != nil {
		return err
	}

	_, err = p.dbU.Transaction(ADD_ROLE_TO_USER_STMT, roleID, userID)
	if err != nil {
		return errors.Join(e.ErrPermissionManagement, e.ErrRepoAdd, err)
	}

	return nil
}

// GetPermissionsByRoleID
func (p *pMRepo) GetPermissionsByRoleID(roleID int) (*[]Permission, error) {
	p.rw.RLock()
	defer p.rw.RUnlock()

	_, err := p.rr.GetOne(roleID)
	if err != nil {
		return nil, err
	}

	permission := Permission{}
	permissions := []Permission{}

	rows, err := p.db.Query(GET_PERMISSIONS_BY_ROLE_ID_STMT, roleID)
	if err != nil {
		return nil, errors.Join(e.ErrPermissionManagement, e.ErrRepoGetAll, e.ErrRepoPreparingStmt, err)
	}
	defer rows.Close()

	for rows.Next() {
		err = rows.Scan(&permission.ID, &permission.Name)
		if err != nil {
			return nil, errors.Join(e.ErrPermissionManagement, e.ErrRepoGetAll, e.ErrRepoExecutingStmt, err)
		}

		permissions = append(permissions, permission)
	}
	err = rows.Err()
	if err != nil {
		return nil, errors.Join(e.ErrPermissionManagement, e.ErrRepoGetAll, e.ErrRepoExecutingStmt, err)
	}

	return &permissions, nil
}

// GetPermissonsByUserID
func (p *pMRepo) GetPermissonsByUserID(userID int) (*[]Permission, error) {
	p.rw.RLock()
	defer p.rw.RUnlock()

	_, err := p.ur.GetOne(userID)
	if err != nil {
		return nil, err
	}

	permission := Permission{}
	permissions := []Permission{}

	rows, err := p.db.Query(GET_PERMISSIONS_BY_USER_ID, userID)
	if err != nil {
		return nil, errors.Join(e.ErrPermissionManagement, e.ErrRepoGetAll, e.ErrRepoPreparingStmt, err)
	}
	defer rows.Close()

	for rows.Next() {
		err = rows.Scan(&permission.ID, &permission.Name)
		if err != nil {
			return nil, errors.Join(e.ErrPermissionManagement, e.ErrRepoGetAll, e.ErrRepoExecutingStmt, err)
		}

		permissions = append(permissions, permission)
	}
	err = rows.Err()
	if err != nil {
		return nil, errors.Join(e.ErrPermissionManagement, e.ErrRepoGetAll, e.ErrRepoExecutingStmt, err)
	}

	return &permissions, nil
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

	_, err = p.dbU.Transaction(REMOVE_PERMISSION_FROM_ROLE_STMT, roleID, permissionID)
	if err != nil {
		return errors.Join(e.ErrPermissionManagement, e.ErrRepoRemove, err)
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
