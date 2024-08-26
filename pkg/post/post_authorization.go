package post

import (
	a "app/pkg/authorization"
	e "app/pkg/errors"
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
	auth a.AuthorizationService
	v    PostValidator
}

func NewPostAuthorization(v PostValidator, auth a.AuthorizationService) PostAuthorization {
	return &postAuthorization{auth, v}
}

// Add
func (p *postAuthorization) Add(userID int, post *Post) error {
	err := p.auth.CheckForAuthorization(userID, PostAdd)
	if err != nil {
		return errors.Join(e.ErrPostDomain, e.ErrOnAdd, err)
	}

	return p.v.Add(post)
}

// AddAll
func (p *postAuthorization) AddAll(userID int, posts *[]*Post) error {
	err := p.auth.CheckForAuthorization(userID, PostAddAll)
	if err != nil {
		return errors.Join(e.ErrPostDomain, e.ErrOnAddAll, err)
	}

	return p.v.AddAll(posts)
}

// Edit
func (p *postAuthorization) Edit(userID int, id int, newPost *Post) error {
	err := p.auth.CheckForAuthorization(userID, PostEdit)
	if err != nil {
		return errors.Join(e.ErrPostDomain, e.ErrOnEdit, err)
	}

	return p.v.Edit(id, newPost)
}

// GetAll
func (p *postAuthorization) GetAll(userID int) (*[]Post, error) {
	err := p.auth.CheckForAuthorization(userID, PostGetAll)
	if err != nil {
		return nil, errors.Join(e.ErrPostDomain, e.ErrOnGetAll, err)
	}

	return p.v.GetAll()
}

// GetOne
func (p *postAuthorization) GetOne(userID int, id int) (*Post, error) {
	err := p.auth.CheckForAuthorization(userID, PostGetOne)
	if err != nil {
		return nil, errors.Join(e.ErrPostDomain, e.ErrOnGetOne, err)
	}

	return p.v.GetOne(id)
}

// Remove
func (p *postAuthorization) Remove(userID int, id int) error {
	err := p.auth.CheckForAuthorization(userID, PostRemove)
	if err != nil {
		return errors.Join(e.ErrPostDomain, e.ErrOnRemove, err)
	}

	return p.v.Remove(id)
}
