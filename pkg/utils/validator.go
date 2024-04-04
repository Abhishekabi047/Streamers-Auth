package utils

import (
	"fmt"

	"github.com/go-playground/validator"
)

func Validate(s interface{}) error {
	validate := validator.New()
	if err := validate.Struct(s); err != nil {
		if _, ok := err.(*validator.InvalidValidationError); ok {
			return err
		}
		errors := err.(validator.ValidationErrors)
		errorMsg := "Validation failed: "
		for _, e := range errors {
			switch e.Tag() {
			case "required":
				errorMsg += fmt.Sprintf("%s is required; ", e.Field())
			case "alpha":
				errorMsg += fmt.Sprintf("%s should contain only alphabetic characters; ", e.Field())
			case "email":
				errorMsg += fmt.Sprintf("%s should be a valid email; ", e.Field())
			case "numeric":
				errorMsg += fmt.Sprintf("%s should contain only numeric values; ", e.Field())
			case "min":
				errorMsg += fmt.Sprintf("%s should contain minimum 8 charchters; ", e.Field())
			case "len":
				errorMsg += fmt.Sprintf("%s should contain 10 numbers; ", e.Field())
			default:
				errorMsg += fmt.Sprintf("%s has an invalid value; ", e.Field())
			}
		}
		return fmt.Errorf(errorMsg)
	}
	return nil
}
