package model

import (
	"fmt"
)

type Author struct {
	Id        int     `json:"id"`
	FirstName string  `json:"first_name" valid:"required~first_name is required"`
	LastName  string  `json:"last_name" valid:"required~last_name is required"`
	Books     *[]Book `json:"books,omitempty"`
}

func (a Author) String() string {
	return fmt.Sprintf(
		`{%d, "%s", "%s", %v}`,
		a.Id, a.FirstName, a.LastName, a.Books,
	)
}

func NewAuthor(firstName, lastName string) *Author {
	return &Author{0, firstName, lastName, nil}
}
