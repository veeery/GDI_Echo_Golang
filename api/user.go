package api

import (
	"net/http"
	"time"

	"github.com/asaskevich/govalidator"
	"github.com/labstack/echo/v4"

	"gitlab.com/veeery/gdi_echo_golang.git/controller/auth"
	"gitlab.com/veeery/gdi_echo_golang.git/db"
	"gitlab.com/veeery/gdi_echo_golang.git/model"
	"gitlab.com/veeery/gdi_echo_golang.git/service"
	"gitlab.com/veeery/gdi_echo_golang.git/utils"
)

func Register(c echo.Context) error {
	
	db := db.DbManager()

	var register auth.RegisterUser
	
	errDTO := c.Bind(&register)

	if errDTO != nil {
		response := service.BuildErrorResponse(errDTO.Error())
		return c.JSON(400, response)
	}

	_, v := govalidator.ValidateStruct(&register)
	if v != nil {
		res := service.BuildErrorResponse(v.Error())
		return c.JSON(409, res)
	}
	
    err := db.Where("email = ?", register.Email).First(&model.User{}).Error
	if err == nil {
		res := service.BuildCustomErrorResponse(utils.ShorcutIsTaken(register.Email))
		return c.JSON(409, res)
	} 

	er := db.Where("hp = ?", register.Hp).First(&model.User{}).Error
	if er == nil {
		res := service.BuildCustomErrorResponse(utils.ShorcutIsTaken(register.Hp))
		return c.JSON(409, res)
	}

	user := model.User{
		FirstName: register.FirstName,
		LastName: register.LastName,
		Email: register.Email,
		Hp: register.Hp,
		Password: register.Password,
	}

	user.HashPassword()
	db.Create(&user)

	token, _:= user.GenerateToken()

	var cookie http.Cookie

	cookie.Name = "token"
	cookie.Value = token
	cookie.Expires = time.Now().Add(7 * 24 * time.Hour)

	c.SetCookie(&cookie)

	return c.JSON(201, echo.Map{
		"token": token,
		"user": user,
	})

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

