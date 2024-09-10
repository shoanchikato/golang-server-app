package model

import (
	"fmt"

	valid "github.com/go-ozzo/ozzo-validation/v4"
)

type Permission struct {
	Id        int    `json:"id,omitempty"`
	Name      string `json:"name"`
	Entity    string `json:"entity"`
	Operation string `json:"operation"`
}

func NewPermission(name, entity, operation string) *Permission {
	return &Permission{0, name, entity, operation}
}

func (p Permission) String() string {
	return fmt.Sprintf(
		`{%d, "%s", "%s", "%s"}`,
		p.Id, p.Name, p.Entity, p.Operation)
}

func (p *Permission) Validate() error {
	return valid.ValidateStruct(p,
		valid.Field(&p.Name, valid.Required.Error("title is required")),
		valid.Field(&p.Entity, valid.Required.Error("body is required")),
		valid.Field(&p.Operation, valid.Required.Error("user_id is required")),
	)
}
