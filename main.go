package main

import (
	"github.com/labstack/echo/v4"
	"gitlab.com/veeery/gdi_echo_golang.git/api"
	"gitlab.com/veeery/gdi_echo_golang.git/db"
)

func main() {
	db.Init()

	e := echo.New()

	baseUrl := "api/"
	ver := "v1"

	serverUrl := baseUrl + ver + "/"

	authRoute := e.Group(serverUrl+"auth") 
	{
		authRoute.GET("/users", api.GetUsers)
		authRoute.POST("/create", api.CreateUsers) 
		// authRoute.DELETE("/remove", api.DeleteUsers)
		// authRoute.POST("/login", api.Login)
	}	

	e.Logger.Fatal(e.Start(":8000"))
}