package api

import (
	"net/http"
	"strconv"
	"time"

	"github.com/asaskevich/govalidator"
	"github.com/go-playground/validator/v10"
	"github.com/labstack/echo/v4"

	"gitlab.com/veeery/gdi_echo_golang.git/db"
	"gitlab.com/veeery/gdi_echo_golang.git/db/table"
	"gitlab.com/veeery/gdi_echo_golang.git/dto/auth"
	"gitlab.com/veeery/gdi_echo_golang.git/model"
	"gitlab.com/veeery/gdi_echo_golang.git/service"
	"gitlab.com/veeery/gdi_echo_golang.git/system"
	"gitlab.com/veeery/gdi_echo_golang.git/utils"
)

func GetUsers(c echo.Context) error {
	db := db.DbManager()
	users := []model.User{}

    db.Find(&users)
	return c.JSON(http.StatusOK, users)
}

func Login(c echo.Context) error {

	var user model.User
	var userLogin auth.LoginUser
	db := db.DbManager()

	errBind := c.Bind(&userLogin)
	if errBind != nil {
		res := service.BuildErrorResponse(errBind.Error(), "Error Bind Login")
		return c.JSON(400, res)
	}

	err := validator.New().Struct(userLogin)

	if err != nil {
		errs := service.ErrorHandler(err)
		res := service.BuildValidateError(errs, utils.ShorcutValidationError())
		return c.JSON(405, res)
	}

	if errExists := db.Where("email = ?", userLogin.Email).First(&user).Error; errExists != nil {
		res := service.BuildErrorResponse(utils.ShorcutInvalid("Email"), utils.ShorcutValidationError())
		return c.JSON(400, res)
	}

	errPassword := user.CheckPassword(userLogin.Password)
	if errPassword != nil {
		res := service.BuildErrorResponse(utils.ShorcutInvalid("Password"), utils.ShorcutValidationError())
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
		res := service.BuildErrorResponse(v.Error(), utils.ShorcutValidationError())
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

func RegisterV2(c echo.Context) error {

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

	userRefresh := system.GetDataCookieToken(c)
	
	errBind := c.Bind(&userRefresh)
	if errBind != nil {
		res := service.BuildErrorResponse(errBind.Error(), "Error Bind Refresh")
		return c.JSON(400, res)
	}

	if errRefresh := db.Table(table.User()).Where("id_user = ?", userRefresh.IdUser).Find(&user).Error; errRefresh != nil {
		res := service.BuildErrorResponse(errRefresh.Error(), utils.ShorcutValidationError())
		return c.JSON(400, res)
	}

	token, err := user.GenerateToken()
	if err != nil {
		res := service.BuildErrorResponse(err.Error(), utils.ShorcutValidationError())
		return c.JSON(400, res)
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

func ResetPassword(c echo.Context) error {

	db := db.DbManager()
	var resetPasswordDTO auth.ResetPassword
	
	errBind := c.Bind(&resetPasswordDTO)
	if errBind != nil {
		res := service.BuildErrorResponse(errBind.Error(), "Error Bind Change Password")
		return c.JSON(400, res)
	}	
	
	//strconv, string convert
	id, _ := strconv.Atoi(c.Param("id"))

	_, v := govalidator.ValidateStruct(resetPasswordDTO)

	if v != nil {
		res := service.BuildErrorResponse("Password Must Be 6 Character", utils.ShorcutValidationError())
		return c.JSON(409, res)
	}

	if (resetPasswordDTO.ConfirmPassword != resetPasswordDTO.Password) {
		res := service.BuildErrorResponse("Password not match", utils.ShorcutValidationError())
		return c.JSON(400, res)
	}
	
	NewHashPassword := model.HashPasswordUpdate(resetPasswordDTO.Password)

	if err := db.Table(table.User()).Where("id_user", id).Update("password", NewHashPassword).Error; err != nil {
		res := service.BuildErrorResponse("Failed to Change Password", utils.ShorcutValidationError())
		return c.JSON(400, res)
	}

	res := service.BuildResponseOnlyMessage(utils.ShorcutSuccessfulyWithParam("Update Password"))
	return c.JSON(200, res)
}

func Profile(c echo.Context) error {

	return nil
}