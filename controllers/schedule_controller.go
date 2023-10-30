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

type ScheduleController interface {
	CreateScheduleController(ctx echo.Context) error
	UpdateScheduleController(ctx echo.Context) error
	GetScheduleController(ctx echo.Context) error
	GetScheduleControllerByDoctor(ctx echo.Context) error
	GetAllScheduleController(ctx echo.Context) error
	DeleteScheduleController(ctx echo.Context) error
}

type ScheduleControllerImpl struct {
	ScheduleService services.ScheduleService
}

func NewScheduleController(scheduleService services.ScheduleService) ScheduleController {
	return &ScheduleControllerImpl{ScheduleService: scheduleService}
}

func (c *ScheduleControllerImpl) CreateScheduleController(ctx echo.Context)  error {
	scheduleCreateRequest := web.ScheduleCreateRequest{}
	err := ctx.Bind(&scheduleCreateRequest)
	if err != nil {
		return ctx.JSON(http.StatusBadRequest, helpers.ErrorResponse("Invalid Client Inputed Schedule"))
	}
	result, err := c.ScheduleService.CreateSchedule(ctx, scheduleCreateRequest)
	if err != nil {
		if strings.Contains(err.Error(), "validation failed") {
			return ctx.JSON(http.StatusBadRequest, helpers.ErrorResponse("Invalid Validation"))
		}
		return ctx.JSON(http.StatusInternalServerError, helpers.ErrorResponse("Create Schedule Error"))
	}
	response := res.CreateScheduleDomainToScheduleResponse(result)

	return ctx.JSON(http.StatusCreated, helpers.SuccessResponse("Succesfully create Schedule", response))
}

func (c *ScheduleControllerImpl) UpdateScheduleController(ctx echo.Context) error {
	scheduleId := ctx.Param("id")
	scheduleIdInt, err := strconv.Atoi(scheduleId)
	if err != nil {
		return ctx.JSON(http.StatusInternalServerError, helpers.ErrorResponse("Invalid Param Id"))
	}

	scheduleUpdateRequest := web.ScheduleUpdateRequest{}
	err = ctx.Bind(&scheduleUpdateRequest)
	if err != nil {
		return ctx.JSON(http.StatusBadRequest, helpers.ErrorResponse("Invalid Client Input"))
	}
	result, err := c.ScheduleService.UpdateSchedule(ctx, scheduleUpdateRequest, scheduleIdInt)
	if err != nil {
		if strings.Contains(err.Error(), "validation failed") {
			return ctx.JSON(http.StatusBadRequest, helpers.ErrorResponse("Invalid validation"))
		}

		if strings.Contains(err.Error(), "Schedule not found"){
			return ctx.JSON(http.StatusNotFound, helpers.ErrorResponse("Schedule not found"))
		}
		return ctx.JSON(http.StatusInternalServerError, helpers.ErrorResponse("Update schedule error"))
	}
	response := res.UpdateScheduleDomainToScheduleResponse(result)
	return ctx.JSON(http.StatusOK, helpers.SuccessResponse("Succesfully Updated Schedule Data", response))
}

func (c *ScheduleControllerImpl) GetScheduleController(ctx echo.Context) error{
	scheduleId := ctx.Param("id")
	scheduleIdInt, err := strconv.Atoi(scheduleId)
	if err != nil {
		return ctx.JSON(http.StatusInternalServerError, helpers.ErrorResponse("Invalid Param ID"))
	}
	result, err := c.ScheduleService.FindById(ctx, scheduleIdInt)
	if err != nil {
		if strings.Contains(err.Error(), "schedule not found"){
			return ctx.JSON(http.StatusNotFound, helpers.ErrorResponse("Schedule not found"))
		}
		return ctx.JSON(http.StatusInternalServerError, helpers.ErrorResponse("Get Schedule Data Error"))
	}
	response := res.ScheduleDomainToScheduleResponse(result)
	return ctx.JSON(http.StatusOK, helpers.SuccessResponse("Successfully Get Schedule Data", response))
}

func (c *ScheduleControllerImpl) GetAllScheduleController(ctx echo.Context) error {
	result, err := c.ScheduleService.FindAll(ctx)
	if err != nil {
		if strings.Contains(err.Error(), "schedule not found"){
			return ctx.JSON(http.StatusNotFound, helpers.ErrorResponse("Schedule not found"))
		}
		return ctx.JSON(http.StatusInternalServerError, helpers.ErrorResponse("Get All Schedule Data Error"))
	}
	response := res.ConvertScheduleResponse(result)
	return ctx.JSON(http.StatusOK, helpers.SuccessResponse("Successfully Get All Schedule Data", response))
}

func (c *ScheduleControllerImpl) DeleteScheduleController(ctx echo.Context) error{
	scheduleId := ctx.Param("id")
	scheduleIdInt, err := strconv.Atoi(scheduleId)
	if err != nil {
		return ctx.JSON(http.StatusInternalServerError, helpers.ErrorResponse("Invalid Param ID"))
	}
	err = c.ScheduleService.DeleteSchedule(ctx, scheduleIdInt)
	if err != nil {
		if strings.Contains(err.Error(), "schedule not found"){
			return ctx.JSON(http.StatusNotFound, helpers.ErrorResponse("Schedule not found"))
		}
		return ctx.JSON(http.StatusInternalServerError, helpers.ErrorResponse("Delete Schedule Data Error"))
	}
	return ctx.JSON(http.StatusOK, helpers.SuccessResponse("Successfully Delete Schedule Data", nil))
}


func (c *ScheduleControllerImpl) GetScheduleControllerByDoctor(ctx echo.Context) error{
	doctorID := middleware.ExtractTokenDoctorId(ctx)

	result, err := c.ScheduleService.FindByDoctor(int(doctorID))
	if err != nil {
		if strings.Contains(err.Error(), "schedule not found"){
			return ctx.JSON(http.StatusNotFound, helpers.ErrorResponse("schedule not found"))
		}
		return ctx.JSON(http.StatusInternalServerError, helpers.ErrorResponse("Get schedule Data Error"))
	}
	response := res.ConvertScheduleResponse(result)
	return ctx.JSON(http.StatusOK, helpers.SuccessResponse("Successfully Get schedule Data", response))
}