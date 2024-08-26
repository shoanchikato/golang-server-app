package role

type Role struct {
	ID   int
	Name string
}

func NewRole(name string) *Role {
	return &Role{0, name}
}
