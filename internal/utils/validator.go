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
	case "max":
		return "This field is required maximum " + fe.Param()
	case "email":
		return "Email is invalid"
	case "oneof":
		return "This field is allow only value in (" + fe.Param() + ")"
	case "alphanum":
		return "This field allow only alpha and numeric"
	}
	return "Unknown error"
}
