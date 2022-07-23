package service

import (
	"strings"
)

type Response struct {
	Message string `json:"message"`
	Data  interface{} `json:"data"`
}

type ResponseOnlyMessage struct {
	Message string `json:"message"`
}

type EmptyResponse struct{}

//Response 200, 201
func BuildResponse(message string, data interface{}) Response {
	res := Response{
		Message: message,
		Data:    data,
	}
	return res
}

//Response 400, 405, 409, bisa Customize Message Error
func BuildErrorResponse(err, titleMessage string) Response {
	split := strings.ReplaceAll(err, ";", "\n") 
	splittedError := strings.Split(split, "\n")

	res := Response{
		Message: titleMessage,
		Data: splittedError,
	}
	return res
}

//Response 400, hanya bisa di pakai ketika return data ialah 'validator'
func BuildValidateError(err []ApiError, titleMessage string) Response {
	
	res := Response{
		Message: titleMessage,
		Data: err,
	}

	return res
}

//Semua Response, bisa Customize dan hanya Messsage (tanpa title)
func BuildResponseOnlyMessage(m string) ResponseOnlyMessage {
	res := ResponseOnlyMessage{
		Message: m,
	}
	return res
}