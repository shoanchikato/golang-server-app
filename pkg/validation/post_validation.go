package validation

import (
	e "app/pkg/errors"
	m "app/pkg/model"
	r "app/pkg/repo"
	s "app/pkg/service"
)

type PostValidator interface {
	Add(post *m.Post) error
	AddAll(posts *[]*m.Post) error
	Edit(id int, newPost *m.Post) error
	GetAll(lastId, limit int) (*[]m.Post, error)
	GetOne(id int) (*m.Post, error)
	Remove(id int) error
}

type postValidator struct {
	Repo    r.PostRepo
	service s.ValidationService
}

func NewPostValidator(repo r.PostRepo, service s.ValidationService) PostValidator {
	return &postValidator{repo, service}
}

// Add
func (v *postValidator) Add(post *m.Post) error {
	err := v.service.Validate(post)
	if err != nil {
		return e.NewValidationError(e.ErrAddValidation, err)
	}

	err = v.Repo.Add(post)
	if err != nil {
		return err
	}

	return nil
}

// AddAll
func (v *postValidator) AddAll(posts *[]*m.Post) error {
	newPosts := *posts
	errs := make([]error, 0, len(newPosts))
	for i := range newPosts {
		err := v.service.Validate(newPosts[i])
		if err != nil {
			errs = append(errs, err)
		}
	}

	if len(errs) > 0 {
		return e.NewValidationError(e.ErrAddAllValidation, errs...)
	}

	err := v.Repo.AddAll(posts)
	if err != nil {
		return err
	}

	return nil
}

// Edit
func (v *postValidator) Edit(id int, newPost *m.Post) error {
	err := v.service.Validate(newPost)
	if err != nil {
		return e.NewValidationError(e.ErrEditValidation, err)
	}

	err = v.Repo.Edit(id, newPost)
	if err != nil {
		return err
	}

	return nil
}

// GetAll
func (v *postValidator) GetAll(lastId, limit int) (*[]m.Post, error) {
	return v.Repo.GetAll(lastId, limit)
}

// GetOne
func (v *postValidator) GetOne(id int) (*m.Post, error) {
	return v.Repo.GetOne(id)
}

// Remove
func (v *postValidator) Remove(id int) error {
	return v.Repo.Remove(id)
}
