package api

import (
	"strconv"

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
	var phoneDTO company.Phone

	errDTO1 := c.Bind(&registerDTO)
	if errDTO1 != nil {
		response := service.BuildErrorResponse(errDTO1.Error(), shortcut.ValidationError()+"1")
		return c.JSON(400, response)
	}

	errDTO2 := c.Bind(&phoneDTO)
	if errDTO2 != nil {
		response := service.BuildErrorResponse(errDTO2.Error(), shortcut.ValidationError()+"12")
		return c.JSON(400, response)
	}

	// err1 := validator.New().Struct(registerDTO)
	// if err1 != nil {
	// 	errs := service.ErrorHandler(err1)
	// 	res := service.BuildValidateError(errs, shortcut.ValidationError())
	// 	return c.JSON(400, res)
	// }

	// err2 := validator.New().Struct(phoneDTO)
	// if err2 != nil {
	// 	errs := service.ErrorHandler(err2)
	// 	res := service.BuildValidateError(errs, shortcut.ValidationError())
	// 	return c.JSON(400, res)
	// }

	// errEmail := db.Where("company_email = ?", registerDTO.CompanyEmail).First(&model.Company{}).Error
	// if errEmail == nil {
	// 	res := service.BuildErrorResponse(shortcut.IsExists("Email"), shortcut.ValidationError())
	// 	return c.JSON(409, res)
	// } 
	
	// errHp := db.Where("company_phone = ?", phoneDTO.CompanyPhone).First(&model.CompanyPhone{}).Error
	// if errHp == nil {
	// 	res := service.BuildErrorResponse(shortcut.IsExists("Phone Number"), shortcut.ValidationError())
	// 	return c.JSON(409, res)
	// }

	companyPhone := []model.CompanyPhone{
		{CompanyPhone: phoneDTO.CompanyPhone},
	}

	// form, _ := c.FormParams()
	
	// fmt.Println(form)

	// var arrayInt []string = []string{"1", "2"}
	
	// integerArray := "company_phone[" + arrayInt[2] + "]"

	// if val, ok := form[integerArray]; ok {
	// 	fmt.Println(val)
	// }

	// companyPhone := []model.CompanyPhone{}

	// if errors.As(err, &ve) {
	// 	out := make([]ApiError, len(ve))
	// 	for i, fe := range ve {

	// 		out[i] = ApiError{fe.Field(), MsgForTag(fe)}
	// 	}
	// 	return out
	// }

	company := model.Company{
		CompanyName: registerDTO.CompanyName,
		CompanyAddress: registerDTO.CompanyAddress,
		CompanyCity: registerDTO.CompanyCity,
		Latitude: registerDTO.Latitude,
		Longitude: registerDTO.Longitude,
		CompanyPhone: companyPhone,
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
		// res := service.BuildErrorResponse(shortcut.NotFound(""), shortcut.ValidationError())
		res := service.BuildResponseOnlyMessage(shortcut.NotFound("Profile"))
		return c.JSON(400, res)
	}

	return c.JSON(200, JsonFormat.Company(
		"Sucessfully Load",
		company,
	))
}