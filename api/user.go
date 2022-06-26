package api

import (
	"net/http"

	"github.com/labstack/echo/v4"
	"gitlab.com/veeery/gdi_echo_golang.git/db"
	"gitlab.com/veeery/gdi_echo_golang.git/model"
	"gitlab.com/veeery/gdi_echo_golang.git/service"
)

func CreateUsers(c echo.Context) error {
	
	db := db.DbManager()
	var userDTO model.User
	
	errDTO := c.Bind(&userDTO)

	if errDTO != nil {
		response := service.BuildErrorResponse("Failed to Request",errDTO.Error())
		return c.JSON(400, response)
	} else {
		result := db.Create(&model.User{
			Name: userDTO.Name,
			// Email: userDTO.Email,
			// Hp: userDTO.Hp,
		})

		if result.Error != nil {
			response := service.BuildErrorResponse("Error ID User", result.Error.Error())
			return c.JSON(400, response)
		} else {
			response := service.BuildResponse("Success", result)
			return c.JSON(201, response)
		}
	}
}

func GetUsers(c echo.Context) error {
	db := db.DbManager()
	users := []model.User{}

    db.Find(&users)
	return c.JSON(http.StatusOK, users)
}

func DeleteUsers(c echo.Context) error {
	db := db.DbManager()
	users := []model.User{}

	db.Delete(&users)
	return c.JSON(http.StatusOK, users)
}

