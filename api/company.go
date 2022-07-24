package api

import (
	"github.com/go-playground/validator/v10"
	"github.com/labstack/echo/v4"
	"gitlab.com/veeery/gdi_echo_golang.git/db"
	"gitlab.com/veeery/gdi_echo_golang.git/dto/company"
	"gitlab.com/veeery/gdi_echo_golang.git/model"
	"gitlab.com/veeery/gdi_echo_golang.git/service"
	JsonFormat "gitlab.com/veeery/gdi_echo_golang.git/service/json_format"
	"gitlab.com/veeery/gdi_echo_golang.git/shortcut"
)

func RegisterCompany(c echo.Context) error {

	db := db.DbManager()
	var registerDTO company.RegisterCompany

	errDTO := c.Bind(&registerDTO)
	if errDTO != nil {
		response := service.BuildErrorResponse(errDTO.Error(), shortcut.ValidationError())
		return c.JSON(400, response)
	}

	err := validator.New().Struct(registerDTO)
	if err != nil {
		errs := service.ErrorHandler(err)
		res := service.BuildValidateError(errs, shortcut.ValidationError())
		return c.JSON(400, res)
	}

	errEmail := db.Where("company_email = ?", registerDTO.CompanyEmail).First(&model.Company{}).Error
	if errEmail == nil {
		res := service.BuildErrorResponse(shortcut.IsExists("Email"), shortcut.ValidationError())
		return c.JSON(409, res)
	} 
	
	errHp := db.Where("company-phone = ?", registerDTO.CompanyPhone).First(&model.Company{}).Error
	if errHp == nil {
		res := service.BuildErrorResponse(shortcut.IsExists("Phone Number"), shortcut.ValidationError())
		return c.JSON(409, res)
	}

	company := model.Company{
		CompanyName: registerDTO.CompanyName,
		CompanyAddress: registerDTO.CompanyAddress,
		CompanyCity: registerDTO.CompanyCity,
		Latitude: registerDTO.Latitude,
		Longitude: registerDTO.Longitude,
		CompanyPhone: registerDTO.CompanyPhone,
		CompanyEmail: registerDTO.CompanyEmail,
		CompanyLeader: registerDTO.CompanyLeader,
		CompanyCategory: registerDTO.CompanyCategory,
	}

	db.Create(&company)

	return c.JSON(
		201,
		JsonFormat.Company(
			shortcut.SuccessfulyCreated("Company"),
			company,
		),
	)

}