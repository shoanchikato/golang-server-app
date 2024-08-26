package permission

type Permission struct {
	ID   int
	Name string
}

func NewPermission(name string) *Permission {
	return &Permission{0, name}
}
