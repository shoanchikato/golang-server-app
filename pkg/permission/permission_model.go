package permission

type Permission struct {
	ID   int
	Name string
	Entity string
	Operation string
}

func NewPermission(name, entity, operation string) *Permission {
	return &Permission{0, name, entity, operation}
}
