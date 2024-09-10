package model

import "fmt"

type Permission struct {
	Id        int    `json:"id,omitempty"`
	Name      string `json:"name" validate:"required~name is required"`
	Entity    string `json:"entity" validate:"required~entity is required"`
	Operation string `json:"operation" validate:"required~operation is required"`
}

func NewPermission(name, entity, operation string) *Permission {
	return &Permission{0, name, entity, operation}
}

func (p Permission) String() string {
	return fmt.Sprintf(
		`{%d, "%s", "%s", "%s"}`,
		p.Id, p.Name, p.Entity, p.Operation)
}
