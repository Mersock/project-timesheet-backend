package utils

import "github.com/go-playground/validator/v10"

func GetValidateTag(fe validator.FieldError) string {
	switch fe.Tag() {
	case "required":
		return "This field is required"
	case "lte":
		return "Should be less than " + fe.Param()
	case "gte":
		return "Should be greater than " + fe.Param()
	case "numeric":
		return "This field is required only numeric"
	case "min":
		return "This field is required minimum " + fe.Param()
	}
	return "Unknown error"
}
