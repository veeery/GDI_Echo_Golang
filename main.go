package main

import (
	"github.com/labstack/echo/v4"
	"gitlab.com/veeery/gdi_echo_golang.git/api"
	"gitlab.com/veeery/gdi_echo_golang.git/db"
	"gitlab.com/veeery/gdi_echo_golang.git/system"
)

func ServerHeader(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		c.Response().Header().Set(echo.HeaderServer, "GDI Internal Golang")
		c.Response().Header().Set(echo.HeaderContentType, "application/json")
		// c.Response().Header().Set("Created By", "Very")
		return next(c)
	}
}

func main() {
	db.Init()

	e := echo.New()
	e.Use(ServerHeader)
	
	baseUrl := "api/"
	ver := "v1"

	serverUrl := baseUrl + ver + "/"

	authRoute := e.Group(serverUrl+"auth") 
	{
		authRoute.POST("/register", api.Register)
		authRoute.POST("/login", api.Login)
		authRoute.GET("/users", api.GetUsers, system.AuthMiddleware())
		authRoute.POST("/logout", api.Logout, system.AuthMiddleware())
		// authRoute.DELETE("/remove", api.DeleteUsers)
		// authRoute.POST("/login", api.Login)
	}	

	
	
	e.Logger.Fatal(e.Start(":8000"))
}