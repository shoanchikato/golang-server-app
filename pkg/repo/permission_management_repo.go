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

type PermissionManagementRepo interface {
	AddPermissionToRole(permissionId, roleId int) error
	AddPermissionsToRole(permissionIds []int, roleId int) error
	AddRoleToUser(roleId, userId int) error
	GetPermissionsByRoleId(roleId int) (*[]m.Permission, error)
	GetPermissonsByUserId(userId int) (*[]m.Permission, error)
	GetRoleByUserId(userId int) (*m.Role, error)
	RemovePermissionFromRole(roleId, permissionId int) error
	RemovePermissionsFromRole(roleId int, permissionIds []int) error
	RemoveRoleFromUser(roleId int, permissionId int) error
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
func (p *pMRepo) AddPermissionToRole(permissionId, roleId int) error {
	permissions, _ := p.GetPermissionsByRoleId(roleId)
	if permissions != nil {
		pp := *permissions
		for i := 0; i < len(pp); i++ {
			if permissionId == pp[i].Id {
				return nil
			}
		}
	}

	_, err := p.rr.GetOne(roleId)
	if err != nil {
		return err
	}

	_, err = p.pr.GetOne(permissionId)
	if err != nil {
		return err
	}

	_, err = p.dbU.Transaction(st.ADD_PERMISSION_TO_ROLE_STMT, permissionId, roleId)
	if err != nil {
		return errors.Join(e.ErrPermissionManagementDomain, e.ErrOnAdd, err)
	}

	return nil
}

// AddPermissionsToRole
func (p *pMRepo) AddPermissionsToRole(permissionIds []int, roleId int) error {
	for i := 0; i < len(permissionIds); i++ {
		permissionId := permissionIds[i]
		err := p.AddPermissionToRole(permissionId, roleId)
		if err != nil {
			return err
		}
	}

	return nil
}

// AddRoleToUser
func (p *pMRepo) AddRoleToUser(roleId, userId int) error {
	role, _ := p.GetRoleByUserId(userId)
	if role != nil {
		return nil
	}

	_, err := p.ur.GetOne(userId)
	if err != nil {
		return err
	}

	_, err = p.rr.GetOne(roleId)
	if err != nil {
		return err
	}

	_, err = p.dbU.Transaction(st.ADD_ROLE_TO_USER_STMT, roleId, userId)
	if err != nil {
		return errors.Join(e.ErrPermissionManagementDomain, e.ErrOnAdd, err)
	}

	return nil
}

// GetPermissionsByRoleId
func (p *pMRepo) GetPermissionsByRoleId(roleId int) (*[]m.Permission, error) {
	p.rw.RLock()
	defer p.rw.RUnlock()

	_, err := p.rr.GetOne(roleId)
	if err != nil {
		return nil, err
	}

	permission := m.Permission{}
	permissions := []m.Permission{}

	rows, err := p.db.Query(st.GET_PERMISSIONS_BY_ROLE_Id_STMT, roleId)
	if err != nil {
		return nil, errors.Join(e.ErrPermissionManagementDomain, e.ErrOnGetAll, e.ErrRepoPreparingStmt, err)
	}
	defer rows.Close()

	for rows.Next() {
		err = rows.Scan(&permission.Id, &permission.Name)
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

// GetPermissonsByUserId
func (p *pMRepo) GetPermissonsByUserId(userId int) (*[]m.Permission, error) {
	p.rw.RLock()
	defer p.rw.RUnlock()

	_, err := p.ur.GetOne(userId)
	if err != nil {
		return nil, err
	}

	permission := m.Permission{}
	permissions := []m.Permission{}

	rows, err := p.db.Query(st.GET_PERMISSIONS_BY_USER_Id, userId)
	if err != nil {
		return nil, errors.Join(e.ErrPermissionManagementDomain, e.ErrOnGetAll, e.ErrRepoPreparingStmt, err)
	}
	defer rows.Close()

	for rows.Next() {
		err = rows.Scan(&permission.Id, &permission.Name)
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

// GetRoleByUserId
func (p *pMRepo) GetRoleByUserId(userId int) (*m.Role, error) {
	p.rw.RLock()
	defer p.rw.RUnlock()

	role := m.Role{}

	row := p.db.QueryRow(st.GET_ROLE_BY_USER_Id_STMT, userId)
	err := row.Scan(&role.Id, &role.Name)
	if err == sql.ErrNoRows {
		return nil, errors.Join(
			e.ErrPermissionManagementDomain,
			e.ErrRepoExecutingStmt,
			e.NewErrRepoNotFound(strconv.Itoa(userId)),
		)
	}

	if err != nil {
		return nil, errors.Join(e.ErrPermissionManagementDomain, e.ErrOnGetOne, e.ErrRepoExecutingStmt, err)
	}

	return &role, nil
}

// RemovePermissionFromRole
func (p *pMRepo) RemovePermissionFromRole(roleId int, permissionId int) error {
	_, err := p.pr.GetOne(permissionId)
	if err != nil {
		return err
	}

	_, err = p.rr.GetOne(roleId)
	if err != nil {
		return err
	}

	_, err = p.GetPermissionsByRoleId(roleId)
	if err != nil {
		return err
	}

	_, err = p.dbU.Transaction(st.REMOVE_PERMISSION_FROM_ROLE_STMT, roleId, permissionId)
	if err != nil {
		return errors.Join(e.ErrPermissionManagementDomain, e.ErrOnRemove, err)
	}

	return nil
}

// RemovePermissionsFromRole
func (p *pMRepo) RemovePermissionsFromRole(roleId int, permissions []int) error {
	for i := 0; i < len(permissions); i++ {
		permissionId := permissions[i]
		err := p.RemovePermissionFromRole(roleId, permissionId)
		if err != nil {
			return err
		}
	}

	return nil
}

// RemoveRoleFromUser
func (p *pMRepo) RemoveRoleFromUser(roleId int, userId int) error {
	_, err := p.ur.GetOne(userId)
	if err != nil {
		return err
	}

	_, err = p.rr.GetOne(roleId)
	if err != nil {
		return err
	}

	_, err = p.GetRoleByUserId(userId)
	if err != nil {
		return err
	}

	_, err = p.dbU.Transaction(st.REMOVE_ROLE_FROM_USER_STMT, roleId, userId)
	if err != nil {
		return errors.Join(e.ErrPermissionManagementDomain, e.ErrOnRemove, err)
	}

	return nil
}
