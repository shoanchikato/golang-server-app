package validation

import (
	e "app/pkg/errors"
	m "app/pkg/model"
	r "app/pkg/repo"
	s "app/pkg/service"
	"fmt"
	"strings"
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
	Repo r.PostRepo
	service s.ValidationService
}

func NewPostValidator(repo r.PostRepo, service s.ValidationService) PostValidator {
	return &postValidator{repo, service}
}

// Add
func (v *postValidator) Add(post *m.Post) error {
	err := v.service.Validate(post)
	if err != nil {
		return e.NewValidationError(e.ErrAddValidation, err.Error())
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
	errs := make([]string, len(newPosts))
	for i := range newPosts {
		err := v.service.Validate(newPosts[i])
		if err != nil {
			errStr := fmt.Sprintf("\n[%d] %s", i, err.Error())
			errs[i] = errStr
		}
	}

	for i := len(newPosts) - 1; i > 0; i-- {
		if errs[i] != "" {
			newErrors := strings.Join(errs, "")
			return e.NewValidationError(e.ErrAddAllValidation, newErrors)
		}
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
		return e.NewValidationError(e.ErrEditValidation, err.Error())
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
