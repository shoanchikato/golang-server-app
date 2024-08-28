package authorization

import (
	e "app/pkg/errors"
	m "app/pkg/model"
	pe "app/pkg/permission"
	s "app/pkg/service"
	v "app/pkg/validation"
	"errors"
)

type PostAuthorization interface {
	Add(userID int, post *m.Post) error
	AddAll(userID int, posts *[]*m.Post) error
	Edit(userID int, id int, newPost *m.Post) error
	GetAll(userID int) (*[]m.Post, error)
	GetOne(userID int, id int) (*m.Post, error)
	Remove(userID int, id int) error
}

type postAuthorization struct {
	auth      s.AuthorizationService
	validator v.PostValidator
}

func NewPostAuthorization(auth s.AuthorizationService, validator v.PostValidator) PostAuthorization {
	return &postAuthorization{auth, validator}
}

// Add
func (p *postAuthorization) Add(userID int, post *m.Post) error {
	err := p.auth.CheckForAuthorization(userID, pe.PostAdd.Name)
	if err != nil {
		return errors.Join(e.ErrPostDomain, e.ErrOnAdd, err)
	}

	return p.validator.Add(post)
}

// AddAll
func (p *postAuthorization) AddAll(userID int, posts *[]*m.Post) error {
	err := p.auth.CheckForAuthorization(userID, pe.PostAddAll.Name)
	if err != nil {
		return errors.Join(e.ErrPostDomain, e.ErrOnAddAll, err)
	}

	return p.validator.AddAll(posts)
}

// Edit
func (p *postAuthorization) Edit(userID int, id int, newPost *m.Post) error {
	err := p.auth.CheckForAuthorization(userID, pe.PostEdit.Name)
	if err != nil {
		return errors.Join(e.ErrPostDomain, e.ErrOnEdit, err)
	}

	return p.validator.Edit(id, newPost)
}

// GetAll
func (p *postAuthorization) GetAll(userID int) (*[]m.Post, error) {
	err := p.auth.CheckForAuthorization(userID, pe.PostGetAll.Name)
	if err != nil {
		return nil, errors.Join(e.ErrPostDomain, e.ErrOnGetAll, err)
	}

	return p.validator.GetAll()
}

// GetOne
func (p *postAuthorization) GetOne(userID int, id int) (*m.Post, error) {
	err := p.auth.CheckForAuthorization(userID, pe.PostGetOne.Name)
	if err != nil {
		return nil, errors.Join(e.ErrPostDomain, e.ErrOnGetOne, err)
	}

	return p.validator.GetOne(id)
}

// Remove
func (p *postAuthorization) Remove(userID int, id int) error {
	err := p.auth.CheckForAuthorization(userID, pe.PostRemove.Name)
	if err != nil {
		return errors.Join(e.ErrPostDomain, e.ErrOnRemove, err)
	}

	return p.validator.Remove(id)
}
