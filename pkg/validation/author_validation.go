package validation

import (
	e "app/pkg/errors"
	m "app/pkg/model"
	r "app/pkg/repo"
	s "app/pkg/service"
	"fmt"
	"strings"

	valid "github.com/asaskevich/govalidator"
)

type AuthorValidator interface {
	Add(author *m.Author) error
	AddAll(authors *[]*m.Author) error
	Edit(id int, newAuthor *m.Author) error
	GetAll(lastId, limit int) (*[]m.Author, error)
	GetOne(id int) (*m.Author, error)
	Remove(id int) error
}

type authorValidator struct {
	Repo r.AuthorRepo
	service s.ValidationService
}

func NewAuthorValidator(repo r.AuthorRepo, service s.ValidationService) AuthorValidator {
	return &authorValidator{repo, service}
}

// Add
func (v *authorValidator) Add(author *m.Author) error {
	err := v.service.Validate(author)
	if err != nil {
		return e.NewValidationError(e.ErrAddValidation, err.Error())
	}

	err = v.Repo.Add(author)
	if err != nil {
		return err
	}

	return nil
}

// AddAll
func (v *authorValidator) AddAll(authors *[]*m.Author) error {
	newAuthors := *authors
	errs := make([]string, len(newAuthors))
	for i := range newAuthors {
		_, err := valid.ValidateStruct(newAuthors[i])
		if err != nil {
			errStr := fmt.Sprintf("\n[%d] %s", i, err.Error())
			errs[i] = errStr
		}
	}

	for i := len(newAuthors) - 1; i > 0; i-- {
		if errs[i] != "" {
			newErrors := strings.Join(errs, "")
			return e.NewValidationError(e.ErrAddAllValidation, newErrors)
		}
	}

	err := v.Repo.AddAll(authors)
	if err != nil {
		return err
	}

	return nil
}

// Edit
func (v *authorValidator) Edit(id int, newAuthor *m.Author) error {
	_, err := valid.ValidateStruct(newAuthor)
	if err != nil {
		return e.NewValidationError(e.ErrEditValidation, err.Error())
	}

	err = v.Repo.Edit(id, newAuthor)
	if err != nil {
		return err
	}

	return nil
}

// GetAll
func (v *authorValidator) GetAll(lastId, limit int) (*[]m.Author, error) {
	return v.Repo.GetAll(lastId, limit)
}

// GetOne
func (v *authorValidator) GetOne(id int) (*m.Author, error) {
	return v.Repo.GetOne(id)
}

// Remove
func (v *authorValidator) Remove(id int) error {
	return v.Repo.Remove(id)
}
