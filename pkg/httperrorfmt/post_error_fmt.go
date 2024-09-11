package httperrorfmt

import (
	a "app/pkg/authorization"
	m "app/pkg/model"
	s "app/pkg/service"
)

type PostHttpErrorFmt interface {
	Add(userId int, post *m.Post) error
	AddAll(userId int, posts *[]*m.Post) error
	Edit(userId int, id int, newAuthor *m.Post) error
	GetAll(userId, lastId, limit int) (*[]m.Post, error)
	GetOne(userId int, id int) (*m.Post, error)
	Remove(userId int, id int) error
}

type postHttpErrorFmt struct {
	authorization a.PostAuthorization
	service       s.HttpErrorFmt
}

func NewPostHttpErrorFmt(authorization a.PostAuthorization, service s.HttpErrorFmt) PostHttpErrorFmt {
	return &postHttpErrorFmt{authorization, service}
}

// Add
func (r *postHttpErrorFmt) Add(userId int, post *m.Post) error {
	err := r.authorization.Add(userId, post)
	if err != nil {
		return r.service.GetError(err)
	}

	return nil
}

// AddAll
func (r *postHttpErrorFmt) AddAll(userId int, posts *[]*m.Post) error {
	err := r.authorization.AddAll(userId, posts)
	if err != nil {
		return r.service.GetError(err)
	}

	return nil
}

// Edit
func (r *postHttpErrorFmt) Edit(userId int, id int, newAuthor *m.Post) error {
	err := r.authorization.Edit(userId, id, newAuthor)
	if err != nil {
		return r.service.GetError(err)
	}

	return nil
}

// GetAll
func (r *postHttpErrorFmt) GetAll(userId int, lastId int, limit int) (*[]m.Post, error) {
	posts, err := r.authorization.GetAll(userId, lastId, limit)
	if err != nil {
		return nil, r.service.GetError(err)
	}

	return posts, nil
}

// GetOne
func (r *postHttpErrorFmt) GetOne(userId int, id int) (*m.Post, error) {
	posts, err := r.authorization.GetOne(userId, id)
	if err != nil {
		return nil, r.service.GetError(err)
	}

	return posts, nil
}

// Remove
func (r *postHttpErrorFmt) Remove(userId int, id int) error {
	err := r.authorization.Remove(userId, id)
	if err != nil {
		return r.service.GetError(err)
	}

	return nil
}
