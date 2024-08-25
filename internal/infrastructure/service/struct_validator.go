package service

import (
	"fmt"

	"github.com/go-playground/validator/v10"
)

type StructValidator struct {
	entity    interface{}
	validator *validator.Validate
}

func NewStructValidator(e interface{}) *StructValidator {
	return &StructValidator{entity: e, validator: validator.New()}
}

func (s *StructValidator) Validate() error {
	if err := s.validator.Struct(s.entity); err != nil {
		errors := err.(validator.ValidationErrors)
		validationErrors := make(map[string]string)
		for _, err := range errors {
			validationErrors[err.Field()] = err.Tag()
		}

		return fmt.Errorf("validation error: %s", validationErrors)
	}

	return nil
}
