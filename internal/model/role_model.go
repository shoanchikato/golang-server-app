package model

import valid "github.com/go-ozzo/ozzo-validation/v4"

type Role struct {
	Id   int    `json:"id,omitempty"`
	Name string `json:"name"`
}

func NewRole(name string) *Role {
	return &Role{0, name}
}

func (r *Role) Validate() error {
	return valid.ValidateStruct(r,
		valid.Field(&r.Name, valid.Required.Error("name is required")),
	)
}
