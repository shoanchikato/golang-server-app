package service

import (
	_ "github.com/asaskevich/govalidator"
	valid "github.com/go-playground/validator"
)

type ValidationService interface {
	Validate(value any) error
}

type validationService struct {
	validator *valid.Validate
}

func NewValidationService() ValidationService {
	validator := valid.New()
	return &validationService{validator}
}

// Validate
func (v *validationService) Validate(value any) error {
	err := v.validator.Struct(value)
	if err != nil {
		return err
	}

	return nil
}
