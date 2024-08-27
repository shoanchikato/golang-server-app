package post

import (
	e "app/pkg/errors"
	"fmt"
	"strings"

	valid "github.com/asaskevich/govalidator"
)

type PostValidator interface {
	Add(post *Post) error
	AddAll(posts *[]*Post) error
	Edit(id int, newPost *Post) error
	GetAll() (*[]Post, error)
	GetOne(id int) (*Post, error)
	Remove(id int) error
}

type postValidator struct {
	Repo PostRepo
}

func NewPostValidator(repo PostRepo) PostValidator {
	return &postValidator{repo}
}

// Add
func (v *postValidator) Add(post *Post) error {
	_, err := valid.ValidateStruct(post)
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
func (v *postValidator) AddAll(posts *[]*Post) error {
	newPosts := *posts
	errs := make([]string, len(newPosts))
	for i := 0; i < len(newPosts); i++ {
		_, err := valid.ValidateStruct(newPosts[i])
		if err != nil {
			errStr := fmt.Sprintf("\n[%d] %s", i, err.Error())
			errs[i] = errStr
		}
	}

	if len(errs) > 0 {
		newErrors := strings.Join(errs, "")
		return e.NewValidationError(e.ErrAddAllValidation, newErrors)
	}

	err := v.Repo.AddAll(posts)
	if err != nil {
		return err
	}

	return nil
}

// Edit
func (v *postValidator) Edit(id int, newPost *Post) error {
	_, err := valid.ValidateStruct(newPost)
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
func (v *postValidator) GetAll() (*[]Post, error) {
	return v.Repo.GetAll()
}

// GetOne
func (v *postValidator) GetOne(id int) (*Post, error) {
	return v.Repo.GetOne(id)
}

// Remove
func (v *postValidator) Remove(id int) error {
	return v.Repo.Remove(id)
}
