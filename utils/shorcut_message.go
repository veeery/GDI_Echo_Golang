package utils

//Message body
func ShorcutIsExists(data string) string {
	s := data + " is already exists"
	return s
}

func ShorcutInvalidPassword() string {
	s := "Password Invalid"
	return s
}

func ShorcutUnAuthorization() string {
	s := "Request Unauthorized"
	return s
}

// Message Title
func ShorcutValidationError() string {
	data := "Validation Error"
	return data
}

func ShorcutError() string {
	data := "Error"
	return data
}

//Only Message
func ShorcutSuccessfulyCreated(data string) string {
	s := data + " has been successfully created"
	return s
}

func ShorcutSuccessfulyWithParam(typeSucessfully string) string {
	s := "Successfully " + typeSucessfully
	return s
}


