package helper

import "github.com/go-playground/validator/v10"

func ValidationErrorsFormatter(err error) map[string][]string {
	var errors []string
	for _, fieldError := range err.(validator.ValidationErrors) {
		errors = append(errors, fieldError.Error())
	}

	return map[string][]string{
		"errors": errors,
	}
}

func PanicIfError(err error) {
	if err != nil {
		panic(err)
	}
}
