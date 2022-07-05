package utils

func ShorcutIsExists(data string) string {
	s := data + " is already exists"
	return s
}

func ShorcutSuccessfulyCreated(data string) string {
	s := data + " has been successfully created"
	return s
}

func ShorcutValidationError() string {
	data := "Validation Error"
	return data
}

func ShorcutInvalidPassword() string {
	s := "Password Invalid"
	return s
}

func ShorcutUnAuthorization() string {
	s := "Request Unauthorized"
	return s
}