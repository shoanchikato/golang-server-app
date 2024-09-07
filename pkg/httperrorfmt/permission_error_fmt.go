package httperrorfmt

import (
	a "app/pkg/authorization"
	m "app/pkg/model"
	s "app/pkg/service"
)

type PermissionHttpErrorFmt interface {
	Add(userId int, permission *m.Permission) error
	AddAll(userId int, permissions *[]*m.Permission) error
	Edit(userId int, id int, newPermission *m.Permission) error
	GetAll(userId, lastId, limit int) (*[]m.Permission, error)
	GetByEntity(userId int, entity string) (*[]m.Permission, error)
	GetOne(userId int, id int) (*m.Permission, error)
	Remove(userId int, id int) error
}

type permissionHttpErrorFmt struct {
	authorization a.PermissionAuthorization
	service       s.HttpErrorFmt
}

func NewPermissionHttpErrorFmt(authorization a.PermissionAuthorization, service s.HttpErrorFmt) PermissionHttpErrorFmt {
	return &permissionHttpErrorFmt{authorization, service}
}

// Add
func (p *permissionHttpErrorFmt) Add(userId int, permission *m.Permission) error {
	err := p.authorization.Add(userId, permission)
	if err != nil {
		return p.service.GetError(err)
	}

	return nil
}

// AddAll
func (p *permissionHttpErrorFmt) AddAll(userId int, permissions *[]*m.Permission) error {
	err := p.authorization.AddAll(userId, permissions)
	if err != nil {
		return p.service.GetError(err)
	}

	return nil
}

// Edit
func (p *permissionHttpErrorFmt) Edit(userId int, id int, newPermission *m.Permission) error {
	err := p.authorization.Edit(userId, id, newPermission)
	if err != nil {
		return p.service.GetError(err)
	}

	return nil
}

// GetAll
func (p *permissionHttpErrorFmt) GetAll(userId int, lastId int, limit int) (*[]m.Permission, error) {
	permissions, err := p.authorization.GetAll(userId, lastId, limit)
	if err != nil {
		return nil, p.service.GetError(err)
	}

	return permissions, nil
}

// GetByEntity
func (p *permissionHttpErrorFmt) GetByEntity(userId int, entity string) (*[]m.Permission, error) {
	permissions, err := p.authorization.GetByEntity(userId, entity)
	if err != nil {
		return nil, p.service.GetError(err)
	}

	return permissions, nil
}

// GetOne
func (p *permissionHttpErrorFmt) GetOne(userId int, id int) (*m.Permission, error) {
	permission, err := p.authorization.GetOne(userId, id)
	if err != nil {
		return nil, p.service.GetError(err)
	}

	return permission, nil
}

// Remove
func (p *permissionHttpErrorFmt) Remove(userId int, id int) error {
	err := p.authorization.Remove(userId, id)
	if err != nil {
		return p.service.GetError(err)
	}

	return nil
}
