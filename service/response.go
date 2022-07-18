package service

import "strings"

type ResponseError struct {
	Message string `json:"message"`
	Data  interface{} `json:"data"`
}

type ResponseSuccess struct {
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
}

type ResponseOnlyMessage struct {
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

func BuildErrorResponse(err, titleMessage string) ResponseError {
	split := strings.ReplaceAll(err, ";", "\n") 
	splittedError := strings.Split(split, "\n")
	
	res := ResponseError{
		Message: titleMessage,
		Data: splittedError,
	}
	return res
}

func BuildResponseOnlyMessage(m string) ResponseOnlyMessage {
	res := ResponseOnlyMessage{
		Message: m,
	}
	return res
}