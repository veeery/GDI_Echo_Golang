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


	apiUrl := "api/"
	baseUrl := ""
	ver := "v1"

	serverUrl := apiUrl + baseUrl + ver + "/"
	
	firstRoute := e.Group(serverUrl+"auth")
	{
		firstRoute.POST("/register", api.Register)
		firstRoute.POST("/login", api.Login)
		firstRoute.PATCH("/reset-password/:id", api.ResetPassword)
	}

	authRoute := e.Group(serverUrl+"auth")
	authRoute.Use(system.AuthMiddleware()) 
	{
		authRoute.DELETE("/logout", api.Logout)
		authRoute.GET("/profile", api.Profile)
		authRoute.POST("/refresh", api.RefreshToken)
		authRoute.PATCH("/change-password", api.ChangePassword)
	}	

	companyRoute := e.Group(serverUrl+"company")
	companyRoute.Use(system.AuthMiddleware())
	{
		companyRoute.POST("/register", api.RegisterCompany)
		companyRoute.GET("/:id", api.ProfileCompany)
	}

	e.Logger.Fatal(e.Start(":8000"))
}