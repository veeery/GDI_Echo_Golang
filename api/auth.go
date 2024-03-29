package api

import (
	"net/http"
	"strconv"
	"time"

	"github.com/go-playground/validator/v10"
	"github.com/labstack/echo/v4"

	"gitlab.com/veeery/gdi_echo_golang.git/db"
	"gitlab.com/veeery/gdi_echo_golang.git/db/table"
	"gitlab.com/veeery/gdi_echo_golang.git/dto/auth"
	"gitlab.com/veeery/gdi_echo_golang.git/model"
	"gitlab.com/veeery/gdi_echo_golang.git/service"
	"gitlab.com/veeery/gdi_echo_golang.git/shortcut"
	"gitlab.com/veeery/gdi_echo_golang.git/system"
)

func GetUsers(c echo.Context) error {
	db := db.DbManager()
	users := []model.User{}

    db.Find(&users)
	return c.JSON(http.StatusOK, users)
}

func Login(c echo.Context) error {

	var user model.User
	var userLoginDTO auth.LoginUser
	db := db.DbManager()

	errBind := c.Bind(&userLoginDTO)
	if errBind != nil {
		res := service.BuildErrorResponse(errBind.Error(), "Error Bind Login")
		return c.JSON(400, res)
	}

	err := validator.New().Struct(userLoginDTO)

	if err != nil {
		errs := service.ErrorHandler(err)
		res := service.BuildValidateError(errs, "Validation Error")
		return c.JSON(400, res)
	}

	if errExists := db.Where("email = ?", userLoginDTO.Email).First(&user).Error; errExists != nil {
		res := service.BuildErrorResponse("Invalid Email", "Validation Error")
		return c.JSON(400, res)
	}

	errPassword := user.CheckPassword(userLoginDTO.Password)
	if errPassword != nil {
		res := service.BuildErrorResponse("Invalid Password", "Validation Error")
		return c.JSON(400, res)
	}

	tokenString, err := user.GenerateToken()
	if err != nil {
		return err
	}

	c.SetCookie(&http.Cookie{
		Name: "token",
		Value: tokenString,
		Expires: time.Now().Add(time.Second * time.Duration(shortcut.ExpiredTokenTime())),
	})

	token := model.Token{
		Token: tokenString,
		Type: "jwt",
		Expired: shortcut.ExpiredTokenTime(),
	}

	return c.JSON(
		200,
		echo.Map{
			"message": "Successfully Login",
			"data": echo.Map{
				"user": user,
				"token": token,
			},
		},
	)
}

func Register(c echo.Context) error {

	var registerDTO auth.RegisterUser
	db := db.DbManager()

	errDTO := c.Bind(&registerDTO)
	if errDTO != nil {
		response := service.BuildErrorResponse(errDTO.Error(), "Validation Error")
		return c.JSON(400, response)
	}

	err := validator.New().Struct(registerDTO)
	if err != nil {
		errs := service.ErrorHandler(err)
		res := service.BuildValidateError(errs, "Validation Error")
		return c.JSON(400, res)
	}
	
    errEmail := db.Where("email = ?", registerDTO.Email).First(&model.User{}).Error
	if errEmail == nil {
		res := service.BuildErrorResponse("Email is Already Exist", "Validation Error")
		return c.JSON(409, res)
	} 
	
	errHp := db.Where("hp = ?", registerDTO.Hp).First(&model.User{}).Error
	if errHp == nil {
		res := service.BuildErrorResponse("Phone Number is Already Exist", "Validation Error")
		return c.JSON(409, res)
	}

	if (registerDTO.Password != registerDTO.ConfirmPassword) {
		res := service.BuildErrorResponse("Password not match", "Validation Error")
		return c.JSON(400, res)
	}

	user := model.User{
		FirstName: registerDTO.FirstName,
		LastName: registerDTO.LastName,
		Email: registerDTO.Email,
		Hp: registerDTO.Hp,
		Password: registerDTO.Password,
	}

	user.HashPassword()
	db.Create(&user)

	return c.JSON(
		201,
		echo.Map{
			"data": echo.Map{
				"message" : "Successfuly Created",
				"user" : user,
			},
		},
	)	
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
	
	res := service.BuildResponseOnlyMessage("Successfully Log Out")
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
		res := service.BuildErrorResponse(errRefresh.Error(), "Validation Error")
		return c.JSON(400, res)
	}

	tokenString, err := user.GenerateToken()
	if err != nil {
		res := service.BuildErrorResponse(err.Error(), "Validation Error")
		return c.JSON(400, res)
	}

	c.SetCookie(&http.Cookie{
		Name: "token",
		Value: tokenString,
		Expires: time.Now().Add(time.Second * time.Duration(shortcut.ExpiredTokenTime())),
	})

	token := model.Token{
		Token: tokenString,
		Type: "jwt",
		Expired: shortcut.ExpiredTokenTime(),
	}
	
	return c.JSON(
		200,
		echo.Map{
			"message": "Successfully Refresh",
			"data": echo.Map{
				"user": user,
				"token": token,
			},
		},
	)
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

	err := validator.New().Struct(resetPasswordDTO)
	if err != nil {
		errs := service.ErrorHandler(err)
		res := service.BuildValidateError(errs, "Validation Error")
		return c.JSON(400, res)
	}

	if (resetPasswordDTO.ConfirmPassword != resetPasswordDTO.Password) {
		res := service.BuildErrorResponse("Password not match", "Validation Error")
		return c.JSON(400, res)
	}
	
	NewHashPassword := model.HashPasswordUpdate(resetPasswordDTO.Password)

	if err := db.Table(table.User()).Where("id_user", id).Update("password", NewHashPassword).Error; err != nil {
		res := service.BuildErrorResponse("Failed to Change Password", "Validation Error")
		return c.JSON(400, res)
	}

	res := service.BuildResponseOnlyMessage("Successfully Update Password")
	return c.JSON(200, res)
}

func Profile(c echo.Context) error {

	db := db.DbManager()
	var user model.User

	userProfile := system.GetDataCookieToken(c)

	errBind := c.Bind(&userProfile)
	if errBind != nil {
		res := service.BuildErrorResponse(errBind.Error(), "Error Bind Refresh")
		return c.JSON(400, res)
	}

	if errRefresh := db.Table(table.User()).Where("id_user = ?", userProfile.IdUser).Find(&user).Error; errRefresh != nil {
		res := service.BuildErrorResponse(errRefresh.Error(), "Validation Error")
		return c.JSON(400, res)
	}

	return c.JSON(
		200,
		echo.Map{
			"data": echo.Map{
				"message":"Successfully Load",
				"user": user,
			},
		},
	)
}

func ChangePassword(c echo.Context) error {

	db := db.DbManager()
	var changePasswordDTO auth.ResetPassword

	userChangePassword := system.GetDataCookieToken(c)
	
	errBind := c.Bind(&changePasswordDTO)
	if errBind != nil {
		res := service.BuildErrorResponse(errBind.Error(), "Error Bind Change Password")
		return c.JSON(400, res)
	}	

	err := validator.New().Struct(changePasswordDTO)
	if err != nil {
		errs := service.ErrorHandler(err)
		res := service.BuildValidateError(errs, "Validation Error")
		return c.JSON(400, res)
	}

	if (changePasswordDTO.ConfirmPassword != changePasswordDTO.Password) {
		res := service.BuildErrorResponse("Password not match", "Validation Error")
		return c.JSON(400, res)
	}
	
	NewHashPassword := model.HashPasswordUpdate(changePasswordDTO.Password)

	if err := db.Table(table.User()).Where("id_user", userChangePassword.IdUser).Update("password", NewHashPassword).Error; err != nil {
		res := service.BuildErrorResponse("Failed to Change Password", "Validation Error")
		return c.JSON(400, res)
	}

	res := service.BuildResponseOnlyMessage("Successfully Update Password")
	return c.JSON(200, res)
}