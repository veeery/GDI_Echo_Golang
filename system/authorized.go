package system

import (
	"fmt"
	"strings"

	"github.com/labstack/echo/v4"
	"gitlab.com/veeery/gdi_echo_golang.git/db"
	"gitlab.com/veeery/gdi_echo_golang.git/model"
	"gitlab.com/veeery/gdi_echo_golang.git/service"
	"gitlab.com/veeery/gdi_echo_golang.git/utils"
)

func AuthMiddleware() echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			
			// var user model.User
			reqToken := c.Request().Header.Get(echo.HeaderAuthorization)
		
			leftRemove := strings.Replace(reqToken, "[", "", -1)
			rightRemove := strings.Replace(leftRemove, "]", "", -1)
			removeBearer := strings.Replace(rightRemove, "Bearer", "", -1)

			splitToken := strings.TrimSpace(removeBearer)

			if reqToken == "" {
				res := service.BuildErrorResponse("Unauthorized",utils.ShorcutUnAuthorization())
				return c.JSON(401, res)
			}

			token, err := model.ValidateToken(splitToken)

			if err != nil {
				res := service.BuildErrorResponse(err.Error(), utils.ShorcutValidationError())
				return c.JSON(401, res)
			}
			
			c.Set("user", model.User{Email: token})

			return next(c)

		}
	}
}

func DeleteAuth(c echo.Context, email string) error {
	db := db.DbManager()

	errBind	:= c.Bind(&email)
	if errBind != nil {
		res := service.BuildErrorResponse(errBind.Error(), "Error Bind Login")
		return c.JSON(400, res)
	}

	// fmt.Println(db.Where("email = ?", email).Take(&model.User{}).Delete(&model.User{}).Error)

	if errDeleteAuth := db.Where("email = ?", email).Take(&model.User{}).Delete(&model.User{}).Error; errDeleteAuth != nil {
		res := service.BuildErrorResponse(errDeleteAuth.Error(), utils.ShorcutValidationError())
		return c.JSON(400, res)
	}
	return nil
}

func GetDataCookieToken(c echo.Context) (data string) {

	cookie := c.Get("user")
	dataCookie := fmt.Sprintf("%v", cookie)

	removeBool := strings.Replace(dataCookie, "0", "", -1)
	leftRemove := strings.ReplaceAll(removeBool, "{", "")
	rightRemove := strings.ReplaceAll(leftRemove, "}", "")
	removeWhiteSpace := strings.ReplaceAll(rightRemove, " ", "")

	data = removeWhiteSpace

	return data
}