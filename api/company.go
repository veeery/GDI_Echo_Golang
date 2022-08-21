package api

import (
	"fmt"
	"strconv"

	"github.com/labstack/echo/v4"
	"gitlab.com/veeery/gdi_echo_golang.git/db"
	"gitlab.com/veeery/gdi_echo_golang.git/model"
	"gitlab.com/veeery/gdi_echo_golang.git/service"
	JsonFormat "gitlab.com/veeery/gdi_echo_golang.git/service/json_format"
	"gitlab.com/veeery/gdi_echo_golang.git/shortcut"
)

func RegisterCompany(c echo.Context) error {

	db := db.DbManager()
	var registerDTO model.Company

	errDTO1 := c.Bind(&registerDTO)
	if errDTO1 != nil {
		response := service.BuildErrorResponse(errDTO1.Error(), "Validation Error 1")
		return c.JSON(400, response)
	}

	company := model.Company{
		CompanyName: registerDTO.CompanyName,
		CompanyAddress: registerDTO.CompanyAddress,
		CompanyPhone: registerDTO.CompanyPhone,
	}

	db.Create(&company)

	fmt.Println(company)

	return c.JSON(
		201,
		JsonFormat.Company(
			shortcut.SuccessfulyCreated("Company"),
			company,
		),
	)
}

func ProfileCompany(c echo.Context) error {

	db := db.DbManager()
	var company model.Company

	errBind := c.Bind(&company)

	if errBind != nil {
		res := service.BuildErrorResponse(errBind.Error(), "Error Bind Profile Company")
		return c.JSON(400, res)
	}

	id, _ := strconv.Atoi(c.Param("id"))

	err := db.Where("id_company = ? ", id).Preload("CompanyPhone").First(&company).Error
	if err != nil {
		res := service.BuildResponseOnlyMessage(shortcut.NotFound("Profile"))
		return c.JSON(400, res)
	}

	return c.JSON(200, JsonFormat.Company(
		"Sucessfully Load",
		company,
	))
}