package model

type Role struct {
	ID   int
	Name string `json:"name" valid:"required~name is required"`
}

func NewRole(name string) *Role {
	return &Role{0, name}
}
