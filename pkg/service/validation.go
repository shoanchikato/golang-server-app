package service

import (
	valid "github.com/asaskevich/govalidator"
)

type ValidationService interface {
	Validate(value any) error
}

type validationService struct{}

func NewValidationService() ValidationService {
	return &validationService{}
}

// Validate
func (v *validationService) Validate(value any) error {
	_, err := valid.ValidateStruct(value)
	if err != nil {
		return err
	}

	return nil
}
