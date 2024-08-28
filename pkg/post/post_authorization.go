package post

import (
	e "app/pkg/errors"
	p "app/pkg/permission"
	"errors"
)

type PostAuthorization interface {
	Add(userID int, post *Post) error
	AddAll(userID int, posts *[]*Post) error
	Edit(userID int, id int, newPost *Post) error
	GetAll(userID int) (*[]Post, error)
	GetOne(userID int, id int) (*Post, error)
	Remove(userID int, id int) error
}

type postAuthorization struct {
	auth p.AuthorizationService
	v    PostValidator
}

func NewPostAuthorization(auth p.AuthorizationService, v PostValidator) PostAuthorization {
	return &postAuthorization{auth, v}
}

// Add
func (p *postAuthorization) Add(userID int, post *Post) error {
	err := p.auth.CheckForAuthorization(userID, PostAdd.Name)
	if err != nil {
		return errors.Join(e.ErrPostDomain, e.ErrOnAdd, err)
	}

	return p.v.Add(post)
}

// AddAll
func (p *postAuthorization) AddAll(userID int, posts *[]*Post) error {
	err := p.auth.CheckForAuthorization(userID, PostAddAll.Name)
	if err != nil {
		return errors.Join(e.ErrPostDomain, e.ErrOnAddAll, err)
	}

	return p.v.AddAll(posts)
}

// Edit
func (p *postAuthorization) Edit(userID int, id int, newPost *Post) error {
	err := p.auth.CheckForAuthorization(userID, PostEdit.Name)
	if err != nil {
		return errors.Join(e.ErrPostDomain, e.ErrOnEdit, err)
	}

	return p.v.Edit(id, newPost)
}

// GetAll
func (p *postAuthorization) GetAll(userID int) (*[]Post, error) {
	err := p.auth.CheckForAuthorization(userID, PostGetAll.Name)
	if err != nil {
		return nil, errors.Join(e.ErrPostDomain, e.ErrOnGetAll, err)
	}

	return p.v.GetAll()
}

// GetOne
func (p *postAuthorization) GetOne(userID int, id int) (*Post, error) {
	err := p.auth.CheckForAuthorization(userID, PostGetOne.Name)
	if err != nil {
		return nil, errors.Join(e.ErrPostDomain, e.ErrOnGetOne, err)
	}

	return p.v.GetOne(id)
}

// Remove
func (p *postAuthorization) Remove(userID int, id int) error {
	err := p.auth.CheckForAuthorization(userID, PostRemove.Name)
	if err != nil {
		return errors.Join(e.ErrPostDomain, e.ErrOnRemove, err)
	}

	return p.v.Remove(id)
}
