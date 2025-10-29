package utils

import (
	"fmt"

	"github.com/go-playground/validator/v10"
)

var validate = validator.New()

func ValidateStruct(data interface{}) map[string]string {
	err := validate.Struct(data)
	if err == nil {
		return nil
	}

	errors := make(map[string]string)
	for _, err := range err.(validator.ValidationErrors) {
		field := err.Field()
		switch err.Tag() {
		case "required":
			errors[field] = fmt.Sprintf("%s is required", field)
		case "email":
			errors[field] = "invalid email format"
		case "min":
			errors[field] = fmt.Sprintf("%s must be at least %s characters", field, err.Param())
		case "number":
			errors[field] = fmt.Sprintf("%s must be a number", field)
		default:
			errors[field] = fmt.Sprintf("%s is invalid", field)
		}
	}
	return errors
}
