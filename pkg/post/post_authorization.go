package post

import (
	a "app/pkg/authorization"
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
	v  PostValidator
}

func NewPostAuthorization(v PostValidator, auth a.AuthorizationService) PostAuthorization {
	return &postAuthorization{auth, v}
}

// Add
func (p *postAuthorization) Add(userID int, post *Post) error {
	err := p.auth.CheckForAuthorization(userID, "post: add")
	if err != nil {
		return err
	}

	return p.v.Add(post)
}

// AddAll
func (p *postAuthorization) AddAll(userID int, posts *[]*Post) error {
	err := p.auth.CheckForAuthorization(userID, "post: add all")
	if err != nil {
		return err
	}

	return p.v.AddAll(posts)
}

// Edit
func (p *postAuthorization) Edit(userID int, id int, newPost *Post) error {
	err := p.auth.CheckForAuthorization(userID, "post: edit")
	if err != nil {
		return err
	}

	return p.v.Edit(id, newPost)
}

// GetAll
func (p *postAuthorization) GetAll(userID int) (*[]Post, error) {
	err := p.auth.CheckForAuthorization(userID, "post: get all")
	if err != nil {
		return nil, err
	}

	return p.v.GetAll()
}

// GetOne
func (p *postAuthorization) GetOne(userID int, id int) (*Post, error) {
	err := p.auth.CheckForAuthorization(userID, "post: get one")
	if err != nil {
		return nil, err
	}

	return p.v.GetOne(id)
}

// Remove
func (p *postAuthorization) Remove(userID int, id int) error {
	err := p.auth.CheckForAuthorization(userID, "post: remove")
	if err != nil {
		return err
	}

	return p.v.Remove(id)
}
