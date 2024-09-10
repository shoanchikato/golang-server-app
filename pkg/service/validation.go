package service

import (
	_ "github.com/asaskevich/govalidator"
	_ "github.com/go-playground/validator"
)

type ValidateModel interface {
	Validate() error
}

type ValidationService interface {
	Validate(model ValidateModel) error
}

type validationService struct {}

func NewValidationService() ValidationService {
	return &validationService{}
}

// Validate
func (v *validationService) Validate(model ValidateModel) error {
	err := model.Validate()
	if err != nil {
		return err
	}

	return nil
}
