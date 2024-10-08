package helpers

import (
	"fmt"
	"strings"

	"github.com/go-playground/validator/v10"
)

var validate *validator.Validate

type ValidationOutput map[string]string

func msgForTag(fe validator.FieldError) string {
	switch fe.Tag() {
	case "required":
		return "Kolom ini harus diisi"
	case "email":
		return "Invalid email"
	case "min":
		return "Minimal " + fe.Param() + " karakter"
	case "number":
		return "Harus berupa angka."
	case "alpha":
		return "Must not contain numbers"
	}

	return fe.Error()
}

func Validation(payload interface{}) ValidationOutput {
	validate = validator.New()
	errs := validate.Struct(payload)
	if errs != nil {
		fmt.Println(errs.Error())

		var apiErrors = make(ValidationOutput)
		for _, err := range errs.(validator.ValidationErrors) {
			fmt.Println()
			apiErrors[strings.ToLower(err.Field())] = msgForTag(err)
		}
		return apiErrors
	}

	return nil
}
