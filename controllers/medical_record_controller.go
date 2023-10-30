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

type MedicalRecordController interface {
	CreateMedicalRecordController(ctx echo.Context) error
	UpdateMedicalRecordController(ctx echo.Context) error
	GetMedicalRecordController(ctx echo.Context) error
	GetAllMedicalRecordController(ctx echo.Context) error
	GetMedicalRecordControllerByPatient(ctx echo.Context) error
	DeleteMedicalRecordController(ctx echo.Context) error
}

type MedicalRecordControllerImpl struct {
	MedicalRecordService services.MedicalRecordService
	RegistrationService services.RegistrationService
}

func NewMedicalRecordController(medicalRecordService services.MedicalRecordService, registrationService services.RegistrationService) MedicalRecordController {
	return &MedicalRecordControllerImpl{MedicalRecordService: medicalRecordService, RegistrationService: registrationService}
}

func (c *MedicalRecordControllerImpl) CreateMedicalRecordController(ctx echo.Context)  error {
	medicalRecordCreateRequest := web.MedicalRecordRequest{}
	err := ctx.Bind(&medicalRecordCreateRequest)
	if err != nil {
		return ctx.JSON(http.StatusBadRequest, helpers.ErrorResponse("Invalid Client Inputed MedicalRecord"))
	}
	result, err := c.MedicalRecordService.CreateMedicalRecord(ctx, medicalRecordCreateRequest)
	if err != nil {
		if strings.Contains(err.Error(), "validation failed") {
			return ctx.JSON(http.StatusBadRequest, helpers.ErrorResponse("Invalid Validation"))
		}
		return ctx.JSON(http.StatusInternalServerError, helpers.ErrorResponse("Create MedicalRecord Error"))
	}
	
	resultRegistration, err := c.RegistrationService.FindById(ctx, int(result.RegistrationID))
	if err != nil {
		if strings.Contains(err.Error(), "medical record not found"){
			return ctx.JSON(http.StatusNotFound, helpers.ErrorResponse("medical record not found"))
		}
		return ctx.JSON(http.StatusInternalServerError, helpers.ErrorResponse("Get medical record data error"))
	}

	response := res.MedicalRecordDomainToMedicalRecordResponse(result, resultRegistration)

	return ctx.JSON(http.StatusCreated, helpers.SuccessResponse("Succesfully create MedicalRecord", response))
}

func (c *MedicalRecordControllerImpl) UpdateMedicalRecordController(ctx echo.Context) error {
	medicalRecordId := ctx.Param("id")
	medicalRecordIdInt, err := strconv.Atoi(medicalRecordId)
	if err != nil {
		return ctx.JSON(http.StatusInternalServerError, helpers.ErrorResponse("Invalid Param Id"))
	}

	medicalRecordUpdateRequest := web.MedicalRecordUpdateRequest{}
	err = ctx.Bind(&medicalRecordUpdateRequest)
	if err != nil {
		return ctx.JSON(http.StatusBadRequest, helpers.ErrorResponse("Invalid Client Input"))
	}
	result, err := c.MedicalRecordService.UpdateMedicalRecord(ctx, medicalRecordUpdateRequest, medicalRecordIdInt)
	if err != nil {
		if strings.Contains(err.Error(), "validation failed") {
			return ctx.JSON(http.StatusBadRequest, helpers.ErrorResponse("Invalid validation"))
		}

		if strings.Contains(err.Error(), "MedicalRecord not found"){
			return ctx.JSON(http.StatusNotFound, helpers.ErrorResponse("MedicalRecord not found"))
		}
		return ctx.JSON(http.StatusInternalServerError, helpers.ErrorResponse("Update medicalRecord error"))
	}
	response := res.UpdateMedicalRecordDomainToMedicalRecordResponse(result)
	return ctx.JSON(http.StatusOK, helpers.SuccessResponse("Succesfully Updated MedicalRecord Data", response))
}

func (c *MedicalRecordControllerImpl) GetMedicalRecordController(ctx echo.Context) error{
	medicalRecordId := ctx.Param("id")
	medicalRecordIdInt, err := strconv.Atoi(medicalRecordId)
	if err != nil {
		return ctx.JSON(http.StatusInternalServerError, helpers.ErrorResponse("Invalid Param ID oo"))
	}
	result, err := c.MedicalRecordService.FindById(ctx, medicalRecordIdInt)
	if err != nil {
		if strings.Contains(err.Error(), "medical record not found"){
			return ctx.JSON(http.StatusNotFound, helpers.ErrorResponse("medical record not found"))
		}
		return ctx.JSON(http.StatusInternalServerError, helpers.ErrorResponse("Get medical record Data Error"))
	}


	resultRegistration, err := c.RegistrationService.FindById(ctx, int(result.RegistrationID))
	if err != nil {
		if strings.Contains(err.Error(), "registration not found"){
			return ctx.JSON(http.StatusNotFound, helpers.ErrorResponse("registration not found"))
		}
		return ctx.JSON(http.StatusInternalServerError, helpers.ErrorResponse("Get registration data error"))
	}

	response := res.MedicalRecordDomainToMedicalRecordResponse(result, resultRegistration)

	return ctx.JSON(http.StatusOK, helpers.SuccessResponse("Successfully Get MedicalRecord Data", response))
}

func (c *MedicalRecordControllerImpl) GetAllMedicalRecordController(ctx echo.Context) error {
	result, err := c.MedicalRecordService.FindAll()
	if err != nil {
		if strings.Contains(err.Error(), "medicalRecord not found"){
			return ctx.JSON(http.StatusNotFound, helpers.ErrorResponse("MedicalRecord not found"))
		}
		return ctx.JSON(http.StatusInternalServerError, helpers.ErrorResponse("Get All MedicalRecord Data Error"))
	}
	response := res.ConvertMedicalRecordResponse(result)
	return ctx.JSON(http.StatusOK, helpers.SuccessResponse("Successfully Get MedicalRecord Data", response))
}

func (c *MedicalRecordControllerImpl) DeleteMedicalRecordController(ctx echo.Context) error{
	medicalRecordId := ctx.Param("id")
	medicalRecordIdInt, err := strconv.Atoi(medicalRecordId)
	if err != nil {
		return ctx.JSON(http.StatusInternalServerError, helpers.ErrorResponse("Invalid Param ID"))
	}
	err = c.MedicalRecordService.DeleteMedicalRecord(ctx, medicalRecordIdInt)
	if err != nil {
		if strings.Contains(err.Error(), "medicalRecord not found"){
			return ctx.JSON(http.StatusNotFound, helpers.ErrorResponse("MedicalRecord not found"))
		}
		return ctx.JSON(http.StatusInternalServerError, helpers.ErrorResponse("Delete MedicalRecord Data Error"))
	}
	return ctx.JSON(http.StatusOK, helpers.SuccessResponse("Successfully Get MedicalRecord Data", nil))
}


func (c *MedicalRecordControllerImpl) GetMedicalRecordControllerByPatient(ctx echo.Context) error{
	patientID := middleware.ExtractTokenPatientId(ctx)

	result, err := c.MedicalRecordService.FindByPatient(int(patientID))
	if err != nil {
		if strings.Contains(err.Error(), "medicalRecord not found"){
			return ctx.JSON(http.StatusNotFound, helpers.ErrorResponse("MedicalRecord not found"))
		}
		return ctx.JSON(http.StatusInternalServerError, helpers.ErrorResponse("Get MedicalRecord Data Error"))
	}
	response := res.ConvertMedicalRecordResponse(result)
	return ctx.JSON(http.StatusOK, helpers.SuccessResponse("Successfully Get MedicalRecord Data", response))
}