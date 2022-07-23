package service

import (
	"errors"

	"github.com/go-playground/validator/v10"
)

type ApiError struct {
	Param   string
	Message string
}

func MsgForTag(fe validator.FieldError) string {
	switch fe.Tag() {
	case "required":
		return "This field is required"
	case "email":
		return "Invalid Email"
	case "min":
		return "Character Min 6"
	}
	return fe.Error()
}

func ErrorHandler(err error) ([]ApiError) {
	var ve validator.ValidationErrors

	if errors.As(err, &ve) {
		out := make([]ApiError, len(ve))
		for i, fe := range ve {
			out[i] = ApiError{fe.Field(), MsgForTag(fe)}
		}
		return out
	}
	return nil	
}