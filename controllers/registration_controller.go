package controllers

import (
	"clinicai-api/models/web"
	"clinicai-api/services"
	"clinicai-api/utils/helpers"
	"clinicai-api/utils/helpers/middleware"
	res "clinicai-api/utils/response"
	"net/http"
	"os"
	"strconv"
	"strings"

	"github.com/labstack/echo/v4"
)


type RegistrationController interface {
	CreateRegistrationController(ctx echo.Context) error
	UpdateRegistrationController(ctx echo.Context) error
	GetRegistrationController(ctx echo.Context) error
	GetRegistrationControllerByPatient(ctx echo.Context) error
	GetAllRegistrationController(ctx echo.Context) error
	DeleteRegistrationController(ctx echo.Context) error
}

type RegistrationControllerImpl struct {
	RegistrationService services.RegistrationService
	DiagnosisService services.DiagnosisService
}


func NewRegistrationController(registrationService services.RegistrationService, diagnosisService services.DiagnosisService) RegistrationController {
	return &RegistrationControllerImpl{RegistrationService: registrationService, DiagnosisService: diagnosisService}
}


func (c *RegistrationControllerImpl) CreateRegistrationController(ctx echo.Context)  error {
	registrationCreateRequest := web.RegistrationCreateRequest{}
	err := ctx.Bind(&registrationCreateRequest)
	if err != nil {
		return ctx.JSON(http.StatusBadRequest, helpers.ErrorResponse("Invalid Client Inputed Registration"))
	}
	
	responseRegistration, err := c.RegistrationService.CreateRegistration(ctx, registrationCreateRequest)
	if err != nil {
		if strings.Contains(err.Error(), "validation failed") {
			return ctx.JSON(http.StatusBadRequest, helpers.ErrorResponse("Invalid Validation"))
		}
		return ctx.JSON(http.StatusInternalServerError, helpers.ErrorResponse("Create Registration Error"))
	}

	//Create Diagnosis
	diagnostic, err := helpers.DiagnosticAI(responseRegistration.Complaint, os.Getenv("OPENAI_KEY"))

	if err != nil {
		return ctx.JSON(http.StatusInternalServerError, helpers.ErrorResponse("error create diagnostic with Open AI"))
	}
	diagnosisSearch := web.DiagnosisCreateRequest{
		RegistrationID: responseRegistration.ID,
		OpenAIResult: diagnostic,
	}
	_, err = c.DiagnosisService.CreateDiagnosis(ctx, diagnosisSearch)
	if err != nil {
		if strings.Contains(err.Error(), "validation failed") {
			return ctx.JSON(http.StatusBadRequest, helpers.ErrorResponse("Invalid Validation"))
		}
		return ctx.JSON(http.StatusInternalServerError, helpers.ErrorResponse("Create Diagnosis Error"))
	}

	result, err := c.DiagnosisService.FindDiagnosesById(ctx, int(responseRegistration.ID))
	if err != nil {
		if strings.Contains(err.Error(), "diagnoses not found"){
			return ctx.JSON(http.StatusNotFound, helpers.ErrorResponse("Diagnoses not found"))
		}
		return ctx.JSON(http.StatusInternalServerError, helpers.ErrorResponse("Get Registration Data Error"))
	}

	registrationResult, err := c.RegistrationService.FindById(ctx, int(responseRegistration.ID))
	if err != nil {
		if strings.Contains(err.Error(), "registration not found"){
			return ctx.JSON(http.StatusNotFound, helpers.ErrorResponse("Registration not found"))
		}
		return ctx.JSON(http.StatusInternalServerError, helpers.ErrorResponse("Get Registration Data Error"))
	}
	results := res.DiagnosisDomainToDiagnosisResponse(result, registrationResult)

	return ctx.JSON(http.StatusCreated, helpers.SuccessResponse("Succesfully create Registration", results))
}

func (c *RegistrationControllerImpl) UpdateRegistrationController(ctx echo.Context) error {
	registrationId := ctx.Param("id")
	registrationIdInt, err := strconv.Atoi(registrationId)
	if err != nil {
		return ctx.JSON(http.StatusInternalServerError, helpers.ErrorResponse("Invalid Param Id"))
	}

	registrationUpdateRequest := web.RegistrationUpdateRequest{}
	err = ctx.Bind(&registrationUpdateRequest)
	if err != nil {
		return ctx.JSON(http.StatusBadRequest, helpers.ErrorResponse("Invalid Client Input"))
	}
	_ , err = c.RegistrationService.UpdateRegistration(ctx, registrationUpdateRequest, registrationIdInt)

	result, err := c.RegistrationService.UpdateRegistration(ctx, registrationUpdateRequest, registrationIdInt)
	if err != nil {
		if strings.Contains(err.Error(), "validation failed") {
			return ctx.JSON(http.StatusBadRequest, helpers.ErrorResponse("Invalid validation"))
		}

		if strings.Contains(err.Error(), "Registration not found"){
			return ctx.JSON(http.StatusNotFound, helpers.ErrorResponse("Registration not found"))
		}
		return ctx.JSON(http.StatusInternalServerError, helpers.ErrorResponse("Update registration error"))
	}

	response := res.UpdateRegistrationDomainToRegistrationResponse(result)
	results, err := c.RegistrationService.FindById(ctx, registrationIdInt)
	if err != nil {
		return ctx.JSON(http.StatusBadRequest, helpers.ErrorResponse("invalid client input"))
	}

	response = res.UpdateRegistrationDomainToRegistrationResponse(results)
	return ctx.JSON(http.StatusOK, helpers.SuccessResponse("Succesfully Updated Registration Data", response))
}

func (c *RegistrationControllerImpl) GetRegistrationController(ctx echo.Context) error{
	registrationId := ctx.Param("id")
	registrationIdInt, err := strconv.Atoi(registrationId)
	if err != nil {
		return ctx.JSON(http.StatusInternalServerError, helpers.ErrorResponse("Invalid Param ID"))
	}
	result, err := c.RegistrationService.FindById(ctx, registrationIdInt)
	if err != nil {
		if strings.Contains(err.Error(), "registration not found"){
			return ctx.JSON(http.StatusNotFound, helpers.ErrorResponse("Registration not found"))
		}
		return ctx.JSON(http.StatusInternalServerError, helpers.ErrorResponse("Get Registration Data Error"))
	}
	response := res.RegistrationDomainToRegistrationResponse(result)
	return ctx.JSON(http.StatusOK, helpers.SuccessResponse("Successfully Get Registration Data", response))
}

func (c *RegistrationControllerImpl) GetAllRegistrationController(ctx echo.Context) error {
	result, err := c.RegistrationService.FindAll(ctx)
	if err != nil {
		if strings.Contains(err.Error(), "registration not found"){
			return ctx.JSON(http.StatusNotFound, helpers.ErrorResponse("Registration not found"))
		}
		return ctx.JSON(http.StatusInternalServerError, helpers.ErrorResponse("Get All Registration Data Error"))
	}
	response := res.ConvertRegistrationResponse(result)
	return ctx.JSON(http.StatusOK, helpers.SuccessResponse("Successfully Get Registration Data", response))
}

func (c *RegistrationControllerImpl) DeleteRegistrationController(ctx echo.Context) error{
	registrationId := ctx.Param("id")
	registrationIdInt, err := strconv.Atoi(registrationId)
	if err != nil {
		return ctx.JSON(http.StatusInternalServerError, helpers.ErrorResponse("Invalid Param ID"))
	}
	err = c.RegistrationService.DeleteRegistration(ctx, registrationIdInt)
	if err != nil {
		if strings.Contains(err.Error(), "registration not found"){
			return ctx.JSON(http.StatusNotFound, helpers.ErrorResponse("Registration not found"))
		}
		return ctx.JSON(http.StatusInternalServerError, helpers.ErrorResponse("Delete Registration Data Error"))
	}
	return ctx.JSON(http.StatusOK, helpers.SuccessResponse("Successfully Get Registration Data", nil))
}

func (c *RegistrationControllerImpl) GetRegistrationControllerByPatient(ctx echo.Context) error{
	patientID := middleware.ExtractTokenPatientId(ctx)

	result, err := c.RegistrationService.FindByPatient(int(patientID))
	if err != nil {
		if strings.Contains(err.Error(), "registration not found"){
			return ctx.JSON(http.StatusNotFound, helpers.ErrorResponse("registration not found"))
		}
		return ctx.JSON(http.StatusInternalServerError, helpers.ErrorResponse("Get registration Data Error"))
	}
	response := res.ConvertRegistrationResponse(result)
	return ctx.JSON(http.StatusOK, helpers.SuccessResponse("Successfully Get registration Data", response))
}