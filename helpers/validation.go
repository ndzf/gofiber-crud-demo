package helpers

import "github.com/go-playground/validator/v10"

type ValidationErrorResponse struct {
	FailedFiled string
	Tag         string
	Value       string
}

var validate = validator.New()

func Validate(model interface{}) []*ValidationErrorResponse {
	var errors []*ValidationErrorResponse
	err := validate.Struct(model)
	if err != nil {
		for _, err := range err.(validator.ValidationErrors) {
			var element ValidationErrorResponse
			element.FailedFiled = err.StructNamespace()
			element.Tag = err.Tag()
			element.Value = err.Param()
			errors = append(errors, &element)
		}
	}

	return errors
}
