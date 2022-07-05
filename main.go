package main

import (
	"github.com/labstack/echo/v4"
	"gitlab.com/veeery/gdi_echo_golang.git/api"
	"gitlab.com/veeery/gdi_echo_golang.git/db"
)

// func ServerHeader(next echo.HandlerFunc) echo.HandlerFunc {
// 	return func(c echo.Context) error {
// 		c.Response().Header().Set(echo.HeaderServer, "GDI Internal Golang")
// 		// c.Response().Header().Set("Created By", "Very")
// 		return next(c)
// 	}
// }

// func AuthMiddleware() echo.MiddlewareFunc {
// 	return func(next echo.HandlerFunc) echo.HandlerFunc {
// 		return func(c echo.Context) error {

// 			reqToken := c.Request().Header.Get(echo.HeaderAuthorization)
// 			splitToken := strings.Split(reqToken, "Bearer: ")

// 			fmt.Println("reqToken :", reqToken)
// 			fmt.Println("SplitToken :", splitToken)

// 			if reqToken == "" {
// 				res := service.BuildErrorResponse("Unauthorized",utils.ShorcutUnAuthorization())
// 				return c.JSON(401, res)
// 			}

// 			err := model.ValidateToken(reqToken)
// 			if err != nil {
// 				return err
// 			}

// 			return next(c)

// 		}
// 	}
// }

func main() {
	db.Init()

	e := echo.New()
	// e.Use(ServerHeader)

	baseUrl := "api/"
	ver := "v1"

	serverUrl := baseUrl + ver + "/"

	authRoute := e.Group(serverUrl+"auth") 
	{
		authRoute.GET("/users", api.GetUsers)
		authRoute.POST("/register", api.Register)
		authRoute.POST("/login", api.Login)
		// authRoute.DELETE("/remove", api.DeleteUsers)
		// authRoute.POST("/login", api.Login)
	}	

	
	e.Logger.Fatal(e.Start(":8000"))
}