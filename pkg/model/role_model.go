package model

type Role struct {
	Id   int    `json:"id,omitempty"`
	Name string `json:"name" validate:"required~name is required"`
}

func NewRole(name string) *Role {
	return &Role{0, name}
}
