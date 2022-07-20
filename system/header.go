package system

import "github.com/labstack/echo/v4"

func ServerHeader(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		c.Response().Header().Set(echo.HeaderServer, "GDI Internal Golang")
		c.Response().Header().Set(echo.HeaderContentType, "application/json")
		c.Response().Header().Set(echo.HeaderAccept, "application/x-www-form-urlencoded")
		// c.Response().Header().Set("Created By", "Very")
		return next(c)
	}
}