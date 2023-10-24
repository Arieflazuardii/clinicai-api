package controllers

import (
	"clinicai-api/models/web"
	"clinicai-api/services"
	"clinicai-api/utils/helpers"
	"clinicai-api/utils/helpers/middleware"
	res "clinicai-api/utils/response"
	"net/http"
	"strconv"
	"strings"

	"github.com/labstack/echo/v4"
)


type PatientController interface {
	RegisterPatientController(ctx echo.Context) error
	LoginPatientController(ctx echo.Context) error
	UpdatePatientController(ctx echo.Context) error
	GetPatientController(ctx echo.Context) error
	GetPatientsController(ctx echo.Context) error
	GetPatientByNameController(ctx echo.Context) error
	DeletePatientController(ctx echo.Context) error
}

type PatientControllerImpl struct {
	PatientService services.PatientService
}


func NewPatientController(patientService services.PatientService) PatientController {
	return &PatientControllerImpl{PatientService: patientService}
}

func (c *PatientControllerImpl) RegisterPatientController(ctx echo.Context) error {
	patientCreateRequest := web.PatientCreateRequest{}
	err := ctx.Bind(&patientCreateRequest)
	if err != nil {
		return ctx.JSON(http.StatusBadRequest, helpers.ErrorResponse("invalid client input"))
	}

	result, err := c.PatientService.CreatePatient(ctx, patientCreateRequest)
	if err != nil {
		if strings.Contains(err.Error(), "validation error") {
			return ctx.JSON(http.StatusBadRequest, helpers.ErrorResponse("invalid validation"))

		}

		if strings.Contains(err.Error(), "email already exist") {
			return ctx.JSON(http.StatusConflict, helpers.ErrorResponse("email already exist"))

		}

		return ctx.JSON(http.StatusInternalServerError, helpers.ErrorResponse("sign up error"))
	}

	response := res.PatientDomainToPatientResponse(result)

	return ctx.JSON(http.StatusCreated, helpers.SuccessResponse("succesfully created account patient", response))
}

func (c *PatientControllerImpl) LoginPatientController(ctx echo.Context) error {
	patientLoginRequest := web.PatientLoginRequest{}

	err := ctx.Bind(&patientLoginRequest)
	if err != nil {
		return ctx.JSON(http.StatusBadRequest, helpers.ErrorResponse("invalid client input"))
	}

	response, err := c.PatientService.LoginPatient(ctx, patientLoginRequest)
	if err != nil {
		if strings.Contains(err.Error(), "validation failed") {
			return ctx.JSON(http.StatusBadRequest, helpers.ErrorResponse("invalid validation"))
		}

		if strings.Contains(err.Error(), "invalid email or password") {
			return ctx.JSON(http.StatusBadRequest, helpers.ErrorResponse("invalid email or password"))
		}
		
		return ctx.JSON(http.StatusInternalServerError, helpers.ErrorResponse("sign in error"))
	}

	patientLoginResponse := res.PatientDomainToPatientLoginResponse(response)

	token, err := middleware.GenerateTokenPatient(&patientLoginResponse, int(response.ID))
	if err != nil {
		return ctx.JSON(http.StatusInternalServerError, helpers.ErrorResponse("generate jwt token error"))
	}

	patientLoginResponse.Token = token
	
	return ctx.JSON(http.StatusCreated, helpers.SuccessResponse("Succesfully Sign In", patientLoginResponse))
}

func (c *PatientControllerImpl) GetPatientController(ctx echo.Context) error {
	patientId := ctx.Param("id")
	patientIdInt, err := strconv.Atoi(patientId)
	if err != nil {
		return ctx.JSON(http.StatusInternalServerError, helpers.ErrorResponse("invalid param id"))
	}

	result, err := c.PatientService.FindById(ctx, patientIdInt)
	if err != nil {
		if strings.Contains(err.Error(), "patient not found"){
			return ctx.JSON(http.StatusNotFound, helpers.ErrorResponse("patient not found"))
		}
		return ctx.JSON(http.StatusInternalServerError, helpers.ErrorResponse("Get patient data error"))
	}
	response :=res.PatientDomainToPatientResponse(result)

	return ctx.JSON(http.StatusOK, helpers.SuccessResponse("Successfully Get Data Patient", response))
}

func(c PatientControllerImpl) GetPatientsController(ctx echo.Context) error {
	result, err := c.PatientService.FindAll(ctx)
	if err != nil {
		if strings.Contains(err.Error(), "patients not found"){
			return ctx.JSON(http.StatusNotFound, helpers.ErrorResponse("patients not found"))
		}

		return ctx.JSON(http.StatusInternalServerError, helpers.ErrorResponse("Get patients data error"))
	}

	response := res.ConvertPatientResponse(result)

	return ctx.JSON(http.StatusOK, helpers.SuccessResponse("Succesfully Get All data Patients", response))
}

func (c PatientControllerImpl) GetPatientByNameController(ctx echo.Context) error {
	patientName := ctx.Param("name")
	
	result, err := c.PatientService.FindByName(ctx, patientName)
	if err != nil {
		if strings.Contains(err.Error(), "patient not found"){
			return ctx.JSON(http.StatusNotFound, helpers.ErrorResponse("patient not found"))
		}
		return ctx.JSON(http.StatusInternalServerError, helpers.ErrorResponse("Get patient data by name error"))
	}
	response := res.PatientDomainToPatientResponse(result)
	return ctx.JSON(http.StatusOK, helpers.SuccessResponse("Succesfully get patient data by name", response))
}

func (c PatientControllerImpl) UpdatePatientController(ctx echo.Context) error {
	patientId := ctx.Param("id")
	patientIdInt, err := strconv.Atoi(patientId)
	if err != nil {
		return ctx.JSON(http.StatusInternalServerError, helpers.ErrorResponse("invalid param id"))
	}

	patientUpdateRequest := web.PatientUpdateRequest{}
	err = ctx.Bind(&patientUpdateRequest)
	if err != nil {
		return ctx.JSON(http.StatusBadRequest, helpers.ErrorResponse("invalid client input"))
	}

	result, err := c.PatientService.UpdatePatient(ctx, patientUpdateRequest, patientIdInt)
	if err != nil {
		if strings.Contains(err.Error(), "validation failed"){
			return ctx.JSON(http.StatusBadRequest, helpers.ErrorResponse("invalid validation"))
		}
		if strings.Contains(err.Error(), "patient not found"){
			return ctx.JSON(http.StatusNotFound, helpers.ErrorResponse("patient not found"))
		}
		return ctx.JSON(http.StatusInternalServerError, helpers.ErrorResponse("update patient error"))
	}
	response := res.PatientDomainToPatientResponse(result)
	return ctx.JSON(http.StatusOK, helpers.SuccessResponse("Successfully updated data patient", response))
}

func (c PatientControllerImpl) DeletePatientController(ctx echo.Context) error {
	patientId := ctx.Param("id")
	patientIdInt, err := strconv.Atoi(patientId)
	if err != nil {
		return ctx.JSON(http.StatusInternalServerError, helpers.ErrorResponse("invalid param id"))
	}
	
	err = c.PatientService.DeletePatient(ctx,patientIdInt)
	if err != nil {
		if strings.Contains(err.Error(), "patient not found"){
			return ctx.JSON(http.StatusNotFound, helpers.ErrorResponse("patient not found"))
		}

		return ctx.JSON(http.StatusInternalServerError, helpers.ErrorResponse("delete data patient error"))
	}

	return ctx.JSON(http.StatusOK, helpers.SuccessResponse("succesfully delete data patient", nil))
}