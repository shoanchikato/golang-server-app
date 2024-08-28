package encrypt

import (
	s "app/pkg/service"
	m "app/pkg/model"
	r "app/pkg/repo"
)

type UserEncryption interface {
	Add(user *m.User) error
	AddAll(users *[]*m.User) error
	Edit(id int, newUser *m.User) error
	GetAll() (*[]m.User, error)
	GetOne(id int) (*m.User, error)
	Remove(id int) error
}

type userEncryption struct {
	repo    r.UserRepo
	encrypt s.EncryptionService
}

func NewUserEncryption(repo r.UserRepo, encrypt s.EncryptionService) UserEncryption {
	return &userEncryption{repo, encrypt}
}

// Add
func (u *userEncryption) Add(user *m.User) error {
	err := u.encrypt.HashPassword(&user.Password)
	if err != nil {
		return err
	}

	return u.repo.Add(user)
}

// AddAll
func (u *userEncryption) AddAll(users *[]*m.User) error {
	newUsers := *users
	for i := 0; i < len(newUsers); i++ {
		user := newUsers[i]
		err := u.encrypt.HashPassword(&user.Password)
		if err != nil {
			return err
		}
	}

	return u.repo.AddAll(users)
}

// Edit
func (u *userEncryption) Edit(id int, newUser *m.User) error {
	return u.repo.Edit(id, newUser)
}

// GetAll
func (u *userEncryption) GetAll() (*[]m.User, error) {
	return u.repo.GetAll()
}

// GetOne
func (u *userEncryption) GetOne(id int) (*m.User, error) {
	return u.repo.GetOne(id)
}

// Remove
func (u *userEncryption) Remove(id int) error {
	return u.repo.Remove(id)
}
