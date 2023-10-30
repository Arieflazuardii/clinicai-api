package controllers

import (
	"clinicai-api/models/web"
	"clinicai-api/services"
	"clinicai-api/utils/helpers"
	"clinicai-api/utils/helpers/middleware"
	res "clinicai-api/utils/response"
	"fmt"
	"net/http"
	"strconv"
	"strings"

	"github.com/labstack/echo/v4"
)


type DoctorController interface {
	RegisterDoctorController(ctx echo.Context) error
	LoginDoctorController(ctx echo.Context) error
	UpdateDoctorController(ctx echo.Context) error
	GetDoctorController(ctx echo.Context) error
	GetDoctorsController(ctx echo.Context) error
	GetDoctorByNameController(ctx echo.Context) error
	DeleteDoctorController(ctx echo.Context) error
}

type DoctorControllerImpl struct {
	DoctorService services.DoctorService
}


func NewDoctorController(doctorService services.DoctorService) DoctorController {
	return &DoctorControllerImpl{DoctorService: doctorService}
}

func (c *DoctorControllerImpl) RegisterDoctorController(ctx echo.Context) error {
	doctorCreateRequest := web.DoctorCreateRequest{}
	err := ctx.Bind(&doctorCreateRequest)
	if err != nil {
		return ctx.JSON(http.StatusBadRequest, helpers.ErrorResponse("invalid client input"))
	}

	result, err := c.DoctorService.CreateDoctor(ctx, doctorCreateRequest)
	if err != nil {
		if strings.Contains(err.Error(), "validation error") {
			return ctx.JSON(http.StatusBadRequest, helpers.ErrorResponse("invalid validation"))

		}

		if strings.Contains(err.Error(), "email already exist") {
			return ctx.JSON(http.StatusConflict, helpers.ErrorResponse("email already exist"))

		}

		return ctx.JSON(http.StatusInternalServerError, helpers.ErrorResponse("sign up error"))
	}

	response := res.DoctorDomainToDoctorResponse(result)

	return ctx.JSON(http.StatusCreated, helpers.SuccessResponse("succesfully created account doctor", response))
}

func (c *DoctorControllerImpl) LoginDoctorController(ctx echo.Context) error {
	doctorLoginRequest := web.DoctorLoginRequest{}

	err := ctx.Bind(&doctorLoginRequest)
	if err != nil {
		return ctx.JSON(http.StatusBadRequest, helpers.ErrorResponse("invalid client input"))
	}

	response, err := c.DoctorService.LoginDoctor(ctx, doctorLoginRequest)
	if err != nil {
		if strings.Contains(err.Error(), "validation failed") {
			return ctx.JSON(http.StatusBadRequest, helpers.ErrorResponse("invalid validation"))
		}

		if strings.Contains(err.Error(), "invalid email or password") {
			return ctx.JSON(http.StatusBadRequest, helpers.ErrorResponse("invalid email or password"))
		}
		
		return ctx.JSON(http.StatusInternalServerError, helpers.ErrorResponse("sign in error"))
	}

	doctorLoginResponse := res.DoctorDomainToDoctorLoginResponse(response)

	token, err := middleware.GenerateTokenDoctor(response.ID)
	if err != nil {
		return ctx.JSON(http.StatusInternalServerError, helpers.ErrorResponse("generate jwt token error"))
	}

	doctorLoginResponse.Token = token
	
	return ctx.JSON(http.StatusCreated, helpers.SuccessResponse("Succesfully Sign In", doctorLoginResponse))
}

func (c *DoctorControllerImpl) GetDoctorController(ctx echo.Context) error {
	doctortId := ctx.Param("id")
	doctorIdInt, err := strconv.Atoi(doctortId)
	if err != nil {
		return ctx.JSON(http.StatusInternalServerError, helpers.ErrorResponse("invalid param id"))
	}

	result, err := c.DoctorService.FindById(ctx, doctorIdInt)
	if err != nil {
		if strings.Contains(err.Error(), "doctor not found"){
			return ctx.JSON(http.StatusNotFound, helpers.ErrorResponse("doctor not found"))
		}
		return ctx.JSON(http.StatusInternalServerError, helpers.ErrorResponse("Get doctor data error"))
	}
	response :=res.DoctorDomainToDoctorResponse(result)

	return ctx.JSON(http.StatusOK, helpers.SuccessResponse("Successfully Get Data Doctor", response))
}

func(c DoctorControllerImpl) GetDoctorsController(ctx echo.Context) error {
	result, err := c.DoctorService.FindAll(ctx)
	if err != nil {
		if strings.Contains(err.Error(), "doctors not found"){
			return ctx.JSON(http.StatusNotFound, helpers.ErrorResponse("doctors not found"))
		}

		return ctx.JSON(http.StatusInternalServerError, helpers.ErrorResponse("Get doctors data error"))
	}

	response := res.ConvertDoctorResponse(result)

	return ctx.JSON(http.StatusOK, helpers.SuccessResponse("Succesfully Get All data Doctors", response))
}

func (c DoctorControllerImpl) GetDoctorByNameController(ctx echo.Context) error {
	doctorName := ctx.Param("name")
	fmt.Println(doctorName)
	
	result, err := c.DoctorService.FindByName(ctx, doctorName)
	if err != nil {
		if strings.Contains(err.Error(), "doctor not found"){
			return ctx.JSON(http.StatusNotFound, helpers.ErrorResponse("doctor not found"))
		}
		return ctx.JSON(http.StatusInternalServerError, helpers.ErrorResponse("Get doctor data by name error"))
	}
	response := res.DoctorDomainToDoctorResponse(result)
	return ctx.JSON(http.StatusOK, helpers.SuccessResponse("Succesfully get doctor data by name", response))
}

func (c DoctorControllerImpl) UpdateDoctorController(ctx echo.Context) error {
	doctorId := ctx.Param("id")
	doctorIdInt, err := strconv.Atoi(doctorId)
	if err != nil {
		return ctx.JSON(http.StatusInternalServerError, helpers.ErrorResponse("invalid param id"))
	}

	doctorUpdateRequest := web.DoctorUpdateRequest{}
	err = ctx.Bind(&doctorUpdateRequest)
	if err != nil {
		return ctx.JSON(http.StatusBadRequest, helpers.ErrorResponse("invalid client input"))
	}

	result, err := c.DoctorService.UpdateDoctor(ctx, doctorUpdateRequest, doctorIdInt)
	if err != nil {
		if strings.Contains(err.Error(), "validation failed"){
			return ctx.JSON(http.StatusBadRequest, helpers.ErrorResponse("invalid validation"))
		}
		if strings.Contains(err.Error(), "doctor not found"){
			return ctx.JSON(http.StatusNotFound, helpers.ErrorResponse("doctor not found"))
		}
		return ctx.JSON(http.StatusInternalServerError, helpers.ErrorResponse("update doctor error"))
	}
	response := res.DoctorDomainToDoctorResponse(result)
	return ctx.JSON(http.StatusOK, helpers.SuccessResponse("Successfully updated data doctor", response))
}

func (c DoctorControllerImpl) DeleteDoctorController(ctx echo.Context) error {
	doctorId := ctx.Param("id")
	doctorIdInt, err := strconv.Atoi(doctorId)
	if err != nil {
		return ctx.JSON(http.StatusInternalServerError, helpers.ErrorResponse("invalid param id"))
	}
	
	err = c.DoctorService.DeleteDoctor(ctx,doctorIdInt)
	if err != nil {
		if strings.Contains(err.Error(), "doctor not found"){
			return ctx.JSON(http.StatusNotFound, helpers.ErrorResponse("doctor not found"))
		}

		return ctx.JSON(http.StatusInternalServerError, helpers.ErrorResponse("delete data doctor error"))
	}

	return ctx.JSON(http.StatusOK, helpers.SuccessResponse("succesfully delete data doctor", nil))
}