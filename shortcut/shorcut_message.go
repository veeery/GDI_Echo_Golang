package shortcut

//Message body
func IsExists(data string) string {
	s := data + " is already exists"
	return s
}

func Invalid(data string) string {
	s :=  data + " Invalid"
	return s
}

func UnAuthorization() string {
	s := "Request Unauthorized"
	return s
}

func NotFound(data string) string {
	s := data + " Not Found"
	return s
}

// Message Title
func ValidationError() string {
	data := "Validation Error"
	return data
}

func Error() string {
	data := "Error"
	return data
}

//Only Message
func SuccessfulyCreated(data string) string {
	s := data + " has been successfully created"
	return s
}

func SuccessfulyWithParam(typeSucessfully string) string {
	s := "Successfully " + typeSucessfully
	return s
}


