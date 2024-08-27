package author

import (
	e "app/pkg/errors"
	"fmt"
	"strings"

	valid "github.com/asaskevich/govalidator"
)

type AuthorValidator interface {
	Add(author *Author) error
	AddAll(authors *[]*Author) error
	Edit(id int, newAuthor *Author) error
	GetAll() (*[]Author, error)
	GetOne(id int) (*Author, error)
	Remove(id int) error
}

type authorValidator struct {
	Repo AuthorRepo
}

func NewAuthorValidator(repo AuthorRepo) AuthorValidator {
	return &authorValidator{repo}
}

// Add
func (v *authorValidator) Add(author *Author) error {
	_, err := valid.ValidateStruct(author)
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
func (v *authorValidator) AddAll(authors *[]*Author) error {
	newAuthors := *authors
	errs := make([]string, len(newAuthors))
	for i := 0; i < len(newAuthors); i++ {
		_, err := valid.ValidateStruct(newAuthors[i])
		if err != nil {
			errStr := fmt.Sprintf("\n[%d] %s", i, err.Error())
			errs[i] = errStr
		}
	}

	for i := len(newAuthors); i > 0; i-- {
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
func (v *authorValidator) Edit(id int, newAuthor *Author) error {
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
func (v *authorValidator) GetAll() (*[]Author, error) {
	return v.Repo.GetAll()
}

// GetOne
func (v *authorValidator) GetOne(id int) (*Author, error) {
	return v.Repo.GetOne(id)
}

// Remove
func (v *authorValidator) Remove(id int) error {
	return v.Repo.Remove(id)
}
