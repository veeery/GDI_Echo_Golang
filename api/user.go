package api

import (
	"net/http"
	"time"

	"github.com/labstack/echo/v4"
	"gitlab.com/veeery/gdi_echo_golang.git/db"
	"gitlab.com/veeery/gdi_echo_golang.git/dto"
	"gitlab.com/veeery/gdi_echo_golang.git/model"
	"gitlab.com/veeery/gdi_echo_golang.git/service"
	"gitlab.com/veeery/gdi_echo_golang.git/utils"
)

func Register(c echo.Context) error {
	
	db := db.DbManager()
	var registerDTO dto.RegisterDTO

	errDTO := c.Bind(&registerDTO)

	if errDTO != nil {
		response := service.BuildErrorResponse(errDTO.Error())
		return c.JSON(400, response)
	}

	// if err := db.Where("name = ?", registerDTO.Name).First(&model.User{}).Error; err == nil {
	// 	// res := service.BuildErrorResponse("Failed to Request", err.Error())
	// 	// return c.JSON(http.StatusConflict, res)
	// 	return c.JSON(http.StatusConflict, "S")
	// }

    err := db.Where("name = ?", registerDTO.Name).First(&model.User{}).Error
	if err == nil {
		res := service.BuildCustomErrorResponse(utils.ShorcutMessage("Name"))
		return c.JSON(http.StatusConflict, res)
	} 

	user := model.User{
		Name: registerDTO.Name,
		Email: registerDTO.Email,
		Password: registerDTO.Password,
	}

	user.HashPassword()
	db.Create(&user)

	token, _:= user.GenerateToken()

	var cookie http.Cookie

	cookie.Name = "token"
	cookie.Value = token
	cookie.Expires = time.Now().Add(7 * 24 * time.Hour)

	c.SetCookie(&cookie)

	return c.JSON(200, echo.Map{
		"token": token,
		"user": user,
	})

}


func CreateUsers(c echo.Context) error {
	
	db := db.DbManager()
	var userDTO model.User
	
	errDTO := c.Bind(&userDTO)

	if errDTO != nil {
		response := service.BuildErrorResponse(errDTO.Error())
		return c.JSON(400, response)
	} else {
		result := db.Create(&model.User{
			Name: userDTO.Name,
			Email: userDTO.Email,
			// Hp: userDTO.Hp,
			Password: userDTO.Password,
		})

		if result.Error != nil {
			response := service.BuildErrorResponse(result.Error.Error())
			return c.JSON(400, response)
		} else {
			response := service.BuildResponse("Success", userDTO)
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

