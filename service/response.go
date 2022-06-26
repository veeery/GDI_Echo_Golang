package service

import (
	"strings"
)

type ResponseError struct {
	Message string `json:"message"`
	Errors  interface{} `json:"errors_hint"`
}

type ResponseSuccess struct {
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
}

func BuildResponse(message string, data interface{}) ResponseSuccess {
	res := ResponseSuccess{
		Message: message,
		Data:    data,
	}
	return res
}

func BuildErrorResponse(message, err string) ResponseError {
	splittedError := strings.Split(err, "\n")	

	res := ResponseError{
		Message: message,
		Errors: splittedError,
	}

	return res
}