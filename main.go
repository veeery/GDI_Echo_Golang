package main

import (
	"github.com/labstack/echo/v4"
	"gitlab.com/veeery/gdi_echo_golang.git/api"
	"gitlab.com/veeery/gdi_echo_golang.git/db"
	"gitlab.com/veeery/gdi_echo_golang.git/system"
)

func main() {
	db.Init()

	e := echo.New()
	e.Use(system.ServerHeader)
	
	baseUrl := "api/"
	ver := "v1"

	serverUrl := baseUrl + ver + "/"

	firstRoute := e.Group(serverUrl+"auth")
	{
		firstRoute.POST("/register", api.Register)
		firstRoute.POST("/login", api.Login)
	}

	authRoute := e.Group(serverUrl+"auth")
	authRoute.Use(system.AuthMiddleware()) 
	{
		authRoute.GET("/users", api.GetUsers)
		authRoute.POST("/logout", api.Logout)
	}	

	
	
	e.Logger.Fatal(e.Start(":8000"))
}