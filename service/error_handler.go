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
		return "This Field is required"
	case "email":
		return "Invalid Email Format"
	case "min":
		return "Character Min is " + fe.Param()
	case "max":
		return "Character Max is " + fe.Param()
	case "numeric":
		if fe.Field() == "Hp" {
			return "Invalid PhoneNumber Format"
		} else {
			return "Please input using numeric"
		}
	case "alphanum":
		return "Please Input Between Text or Number"
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