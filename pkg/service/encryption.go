package service

import (
	e "app/pkg/errors"
	"errors"

	"golang.org/x/crypto/bcrypt"
)

type EncryptionService interface {
	HashPassword(password *string) error
	CheckPassword(hash, password *string) (bool, error)
}

type encryptionService struct{}

func NewEncryptionService() EncryptionService {
	return &encryptionService{}
}

func (en *encryptionService) HashPassword(password *string) error {
	hashBytes, err := bcrypt.GenerateFromPassword([]byte(*password), 12)
	if err != nil {
		return errors.Join(e.ErrEncryptPassword, err)
	}

	*password = string(hashBytes)

	return nil
}

func (en *encryptionService) CheckPassword(hash, password *string) (bool, error) {
	err := bcrypt.CompareHashAndPassword([]byte(*hash), []byte(*password))
	if err != nil {
		return false, errors.Join(e.ErrEncryptionComparePassword, err)
	}

	return true, nil
}
