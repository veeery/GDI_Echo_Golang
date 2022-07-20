package api

import (
	"net/http"
	"strconv"
	"time"

	"github.com/asaskevich/govalidator"
	"github.com/labstack/echo/v4"

	"gitlab.com/veeery/gdi_echo_golang.git/controller/auth"
	"gitlab.com/veeery/gdi_echo_golang.git/db"
	"gitlab.com/veeery/gdi_echo_golang.git/model"
	"gitlab.com/veeery/gdi_echo_golang.git/service"
	"gitlab.com/veeery/gdi_echo_golang.git/system"
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
		res := service.BuildErrorResponse("Password Must Be 6 Character", utils.ShorcutValidationError())
		return c.JSON(409, res)
	}

    err := db.Where("email = ?", register.Email).First(&model.User{}).Error
	if err == nil {
		res := service.BuildErrorResponse(utils.ShorcutIsExists("Email"), utils.ShorcutValidationError())
		return c.JSON(409, res)
	} 
	
	errHp := db.Where("hp = ?", register.Hp).First(&model.User{}).Error
	if errHp == nil {
		res := service.BuildErrorResponse(utils.ShorcutIsExists("Phone Number"), utils.ShorcutValidationError())
		return c.JSON(409, res)
	}

	if (register.Password != register.ConfirmPassword) {
		res := service.BuildErrorResponse("Password not match", utils.ShorcutValidationError())
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
	
 	// Setting MaxAge<0 means delete cookie now.
	cookie := http.Cookie{
		Name: "token",
		Value: "",
		Expires: time.Now().Add(-time.Hour),
		MaxAge: -1,
	} 

	c.SetCookie(&cookie)
	
	res := service.BuildResponseOnlyMessage(utils.ShorcutSuccessfulyWithParam("Log Out"))
	return c.JSON(200, res)
}

func RefreshToken(c echo.Context) error {

	var user model.User
	db := db.DbManager()
	dataEmail := system.GetDataCookieToken(c)

	errBind := c.Bind(&dataEmail)
	if errBind != nil {
		res := service.BuildErrorResponse(errBind.Error(), "Error Bind Refresh")
		return c.JSON(400, res)
	}

	if errRefresh := db.Where("email = ?", dataEmail).First(&user).Error; errRefresh != nil {
		res := service.BuildErrorResponse(errRefresh.Error(), utils.ShorcutValidationError())
		return c.JSON(400, res)
	}

	token, err := user.GenerateToken()
	if err != nil {
		return err
	}

	c.SetCookie(&http.Cookie{
		Name: "token",
		Value: token,
		Expires: time.Now().Add(time.Second * time.Duration(utils.ExpiredTokenTime())),
	})

	return c.JSON(200, echo.Map{
		"message": "Successfully Refresh",
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

func ChangePassword(c echo.Context) error {

	db := db.DbManager()
	
	errBind := c.Bind(&auth.UpdatePassword{})
	if errBind != nil {
		res := service.BuildErrorResponse(errBind.Error(), "Error Bind Change Password")
		return c.JSON(400, res)
	}	
	
	id, _ := strconv.Atoi(c.Param("id"))

	password := c.FormValue("password")
	confirmPassword := c.FormValue("confirm_password")

	_, v := govalidator.ValidateStruct(&auth.UpdatePassword{Password: password, ConfirmPassword: confirmPassword})
	if v!= nil {
		res := service.BuildErrorResponse("Password Must Be 6 Character", utils.ShorcutValidationError())
		return c.JSON(409, res)
	}


	if (password != confirmPassword) {
		res := service.BuildErrorResponse("Password not match", utils.ShorcutValidationError())
		return c.JSON(400, res)
	}
	
	NewHashPassword := model.HashPasswordUpdate(password)

	if err := db.Table("users").Where("id_user", id).Update("password", NewHashPassword).Error; err != nil {
		res := service.BuildErrorResponse("Failed to Change Password", utils.ShorcutValidationError())
		return c.JSON(400, res)
	}


	res := service.BuildResponseOnlyMessage(utils.ShorcutSuccessfulyWithParam("Update Password"))
	return c.JSON(200, res)

}