package service

import "strings"

type ResponseError struct {
	Message interface{} `json:"message"`
	// Errors  interface{} `json:"errors_hint"`
}

type ResponseSuccess struct {
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
}

type ResponseCustomMessage struct {
	Message string `json:"message"`
}

type EmptyResponse struct{}

func BuildResponse(message string, data interface{}) ResponseSuccess {
	res := ResponseSuccess{
		Message: message,
		Data:    data,
	}
	return res
}

func BuildErrorResponse(err string) ResponseError {
	split := strings.ReplaceAll(err, ";", "\n") 
	splittedError := strings.Split(split, "\n")
	
	res := ResponseError{
		Message: splittedError,
	}
	return res
}

func BuildCustomErrorResponse(messageError string) ResponseCustomMessage {
	res := ResponseCustomMessage{
		Message: messageError,
	}
	return res
}