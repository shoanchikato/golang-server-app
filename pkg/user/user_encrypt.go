package user

import (
	s "app/pkg/service"
)

type UserEncryption interface {
	Add(user *User) error
	AddAll(users *[]*User) error
	Edit(id int, newUser *User) error
	GetAll() (*[]User, error)
	GetOne(id int) (*User, error)
	Remove(id int) error
}

type userEncryption struct {
	repo    UserRepo
	encrypt s.EncryptionService
}

func NewUserEncryption(repo UserRepo, encrypt s.EncryptionService) UserEncryption {
	return &userEncryption{repo, encrypt}
}

// Add
func (u *userEncryption) Add(user *User) error {
	err := u.encrypt.HashPassword(&user.Password)
	if err != nil {
		return err
	}

	return u.repo.Add(user)
}

// AddAll
func (u *userEncryption) AddAll(users *[]*User) error {
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
func (u *userEncryption) Edit(id int, newUser *User) error {
	return u.repo.Edit(id, newUser)
}

// GetAll
func (u *userEncryption) GetAll() (*[]User, error) {
	return u.repo.GetAll()
}

// GetOne
func (u *userEncryption) GetOne(id int) (*User, error) {
	return u.repo.GetOne(id)
}

// Remove
func (u *userEncryption) Remove(id int) error {
	return u.repo.Remove(id)
}
