package helper

import "github.com/go-playground/validator/v10"

func Validate(data interface{}) error {
	err := validator.New().Struct(data)
	if err != nil {
		var fieldErrors []validator.FieldError
		for _, err := range err.(validator.ValidationErrors) {
			fieldErrors = append(fieldErrors, err)
		}
		return ErrorArrayToError(fieldErrors)
	}
	return nil
}
