package api

import (
	"fmt"
	"strconv"

	"github.com/go-playground/validator/v10"
	"github.com/labstack/echo/v4"
	"gitlab.com/veeery/gdi_echo_golang.git/db"
	"gitlab.com/veeery/gdi_echo_golang.git/model"
	"gitlab.com/veeery/gdi_echo_golang.git/service"
)

func RegisterCompany(c echo.Context) error {

	db := db.DbManager()
	var registerDTO model.Company
	var phoneDTO model.CompanyPhone

	errDTO := c.Bind(&registerDTO)
	if errDTO != nil {
		response := service.BuildErrorResponse(errDTO.Error(), "Validation Error")
		return c.JSON(400, response)
	}

	//validasi registerDTO sesuai dengan validate di modal.Company
	err := validator.New().Struct(&registerDTO)
	if err != nil {
		errs := service.ErrorHandler(err)
		res := service.BuildValidateError(errs, "Validation Error")
		return c.JSON(400, res)
	}

	//pindahkan data dari registerDTO(modal.Company) ke phoneDTO(modal.CompanyPhone)
	//untuk memvalidasi sesuai dengan validate di modal.CompanyPhone
	for i:=0; i< len(registerDTO.CompanyPhone); i++ {
		phoneDTO = registerDTO.CompanyPhone[i]
	}

	//validate phoneDTO(modal.CompanyPhone)
	errPhone := validator.New().Struct(&phoneDTO)
	fmt.Println(phoneDTO)
	if errPhone != nil {
		errs := service.ErrorHandler(errPhone)
		res := service.BuildValidateError(errs, "Validation Errossr")
		return c.JSON(400, res)
	}
	
    errEmail := db.Where("company_email = ?", registerDTO.CompanyEmail).First(&model.Company{}).Error
	if errEmail == nil {
		res := service.BuildErrorResponse("Email is Already Exist", "Validation Error")
		return c.JSON(409, res)
	} 
	
	errHp := db.Where("company_phone = ?", registerDTO.CompanyPhone).First(&model.Company{}).Error
	if errHp == nil {
		res := service.BuildErrorResponse("PhoneNumber is Already Exist", "Validation Error")
		return c.JSON(409, res)
	}

	db.Create(&registerDTO)

	return c.JSON(
		201,
		echo.Map{
			"message": "Succesfully Created",
			"data": echo.Map{
				"company": registerDTO,
			},
		},
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
		res := service.BuildResponseOnlyMessage("Profile Not Found")
		return c.JSON(400, res)
	}

	return c.JSON(
		200,
		echo.Map{
			"message": "Succesfully Load",
			"data": echo.Map{
				"company": company,
			},
		},
	)
}