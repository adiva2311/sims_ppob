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
			errors[field] = fmt.Sprintf("%s tidak boleh kosong, harus diisi", field)
		case "email":
			errors[field] = "Parameter email tidak sesuai format"
		case "min":
			errors[field] = fmt.Sprintf("%s minimal harus %s karakter", field, err.Param())
		case "number":
			errors[field] = fmt.Sprintf("%s harus diisi dengan angka", field)
		case "gte":
			errors[field] = fmt.Sprintf("%s harus lebih besar dari %s", field, err.Param())
		default:
			errors[field] = fmt.Sprintf("%s tidak valid", field)
		}
	}
	return errors
}
