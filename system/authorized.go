package system

import (
	"strings"

	"github.com/labstack/echo/v4"
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

			_, err := model.ValidateToken(splitToken)

			if err != nil {
				res := service.BuildErrorResponse(err.Error(), err.Error())
				return c.JSON(401, res)
			}
			
			c.Set("user", model.User{})

			return next(c)

		}
	}
}