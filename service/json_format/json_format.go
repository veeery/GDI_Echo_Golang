package JsonFormat

import (
	"github.com/labstack/echo/v4"
	"gitlab.com/veeery/gdi_echo_golang.git/model"
)

//JSON with Full Data of User
func AuthUser(
		message string, 
		token model.Token,
		user model.User, 
	) (echo.Map) {
	
	return echo.Map{
		"message": message,
		"data": echo.Map{
			"token": token,
			"user":user,
		},
	}
}

//JSON with 
func User(
		message string, 
		user model.User, 
	) (echo.Map) {
	
	return echo.Map{
		"message": message,
		"data": echo.Map{
			"user":user,
		},
	}
}