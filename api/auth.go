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

func Login(c echo.Context) error {

	var user model.User
	var userLogin auth.LoginUser
	db := db.DbManager()

	errBind := c.Bind(&userLogin)
	if errBind != nil {
		res := service.BuildErrorResponse(errBind.Error(), "Error Bind Login")
		return c.JSON(400, res)
	}
	
	if errExists := db.Where("email = ?", userLogin.Email).First(&user).Error; errExists != nil {
		res := service.BuildErrorResponse(errExists.Error(), utils.ShorcutValidationError())
		return c.JSON(400, res)
	}

	errPassword := user.CheckPassword(userLogin.Password)
	if errPassword != nil {
		res := service.BuildErrorResponse(utils.ShorcutInvalidPassword(), utils.ShorcutValidationError())
		return c.JSON(400, res)
	}

	token, err := user.GenerateToken()
	if err != nil {
		return err
	}

	// refreshToken, errRefresh := user.GenerateRefreshToken()
	// if errRefresh != nil {
	// 	return err
	// }

	c.SetCookie(&http.Cookie{
		Name: "token",
		Value: token,
		Expires: time.Now().Add(time.Second * time.Duration(utils.ExpiredTokenTime())),
	})

	return c.JSON(200, echo.Map{
		"message": "Successfully login",
		"data": echo.Map{
			"token": echo.Map{
				"access_token": token,
				"type": "jwt",
				"expired": utils.ExpiredTokenTime(),
			},
			"data": user,
		},
	})

}

func Register(c echo.Context) error {

	var register auth.RegisterUser
	db := db.DbManager()

	errDTO := c.Bind(&register)

	if errDTO != nil {
		response := service.BuildErrorResponse(errDTO.Error(), utils.ShorcutValidationError())
		return c.JSON(400, response)
	}

	_, v := govalidator.ValidateStruct(&register)
	if v != nil {
		res := service.BuildErrorResponse(v.Error(), utils.ShorcutValidationError())
		return c.JSON(409, res)
	}

    err := db.Where("email = ?", register.Email).First(&model.User{}).Error
	if err == nil {
		res := service.BuildResponseOnlyMessage(utils.ShorcutIsExists("Email"))
		return c.JSON(409, res)
	} 
	
	errHp := db.Where("hp = ?", register.Hp).First(&model.User{}).Error
	if errHp == nil {
		res := service.BuildResponseOnlyMessage(utils.ShorcutIsExists("Phone number"))
		return c.JSON(409, res)
	}

	if (register.Password != register.ConfirmPassword) {
		res := service.BuildErrorResponse("Password are not same", utils.ShorcutValidationError())
		return c.JSON(400, res)
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

	return c.JSON(201, echo.Map{
		"message": utils.ShorcutSuccessfulyCreated("Users"),
		"user": user,
	})
}

func Logout(c echo.Context) error {
	
	// see https://golang.org/pkg/net/http/#Cookie
 	// Setting MaxAge<0 means delete cookie now.
	cookie := http.Cookie{
		Name: "token",
		MaxAge: -1,
	} 
	
	

	c.SetCookie(&cookie)
	
	res := service.BuildResponseOnlyMessage("Successfully Log out")
	return c.JSON(200, res)
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

