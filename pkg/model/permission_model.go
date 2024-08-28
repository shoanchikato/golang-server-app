package model

type Permission struct {
	Id        int
	Name      string `json:"name" valid:"required~name is required"`
	Entity    string `json:"entity" valid:"required~entity is required"`
	Operation string `json:"operation" valid:"required~operation is required"`
}

func NewPermission(name, entity, operation string) *Permission {
	return &Permission{0, name, entity, operation}
}
