package services

import (
	"clinicai-api/models/domain"
	"clinicai-api/models/web"
	"clinicai-api/repository"
	"clinicai-api/utils/helpers"
	req "clinicai-api/utils/request"
	"fmt"

	"github.com/go-playground/validator"
	"github.com/labstack/echo/v4"
)

type RegistrationService interface {
	CreateRegistration(ctx echo.Context, request web.RegistrationCreateRequest) (*domain.Registration, error)
	UpdateRegistration(ctx echo.Context, request web.RegistrationUpdateRequest, id int) (*domain.Registration, error)
	FindById(ctx echo.Context, id int) (*domain.Registration, error)
	FindByPatient(id int) ([]domain.Registration, error)
	FindAll(ctx echo.Context) ([]domain.Registration, error)
	DeleteRegistration(ctx echo.Context, id int) error
}

type RegistrationServiceImpl struct {
	RegistrationRepository repository.RegistrationRepository
	Validate * validator.Validate
}

func NewRegistrationService(RegistrationRepository repository.RegistrationRepository, validate *validator.Validate) *RegistrationServiceImpl { 
	return &RegistrationServiceImpl{
		RegistrationRepository: RegistrationRepository,
		Validate: validate,
	}
}

func (service *RegistrationServiceImpl) CreateRegistration(ctx echo.Context, request web.RegistrationCreateRequest) (*domain.Registration, error) {
	err := service.Validate.Struct(request)
	if err != nil {
		return nil, helpers.ValidationError(ctx, err)
	}

	registration := req.RegistrationCreateRequestToRegistrationDomain(request)
	scheduleQuota, err := service.RegistrationRepository.GetScheduleQuota(int(registration.ScheduleID))

	if err != nil {
		return nil, fmt.Errorf("error when checking schedule availability: %s", err.Error())
	}
	
	if scheduleQuota <= 0 {
		fmt.Println(scheduleQuota)
		return nil, fmt.Errorf("quota is not available")
	}

	result, err := service.RegistrationRepository.Create(registration)
	if err != nil {
		return nil, fmt.Errorf("error when creating registration: %s", err.Error())
	}

	return result, nil
}


func (service *RegistrationServiceImpl) UpdateRegistration(ctx echo.Context, request web.RegistrationUpdateRequest, id int) (*domain.Registration, error) {
	err := service.Validate.Struct(request)
	if err != nil {
		return nil, helpers.ValidationError(ctx, err)
	}

	existingRegistration, _ := service.RegistrationRepository.FindById(id)
	if existingRegistration == nil {
		return nil, fmt.Errorf("registration not found")
	}
	registration := req.RegistrationUpdateRequestToRegistrationDomain(request)

	result, err := service.RegistrationRepository.Update(registration, id)
	if err != nil {
		return nil, fmt.Errorf("error when update registration: %s", err.Error())
	}
	return result, nil
}

func (service *RegistrationServiceImpl) FindById(ctx echo.Context, id int) (*domain.Registration, error) {
	registration, _ := service.RegistrationRepository.FindById(id)
	if registration == nil {
		return nil, fmt.Errorf("registration not found")
	}

	return registration, nil
}

func (service *RegistrationServiceImpl) FindAll(ctx echo.Context) ([]domain.Registration, error) {
	registration, _ := service.RegistrationRepository.FindAll()
	if registration == nil {
		return nil, fmt.Errorf("registration not found")
	}

	return registration, nil
}

func (service *RegistrationServiceImpl) DeleteRegistration(ctx echo.Context, id int) error {
	registration, _ := service.RegistrationRepository.FindById(id)
	if registration == nil {
		return fmt.Errorf("registration not found")
	}

	err := service.RegistrationRepository.Delete(id)
	if err != nil {
		return fmt.Errorf("error when deleting registration: %s", err)
	}

	return nil
}

func (service *RegistrationServiceImpl) FindByPatient(id int) ([]domain.Registration, error) {
	schedule, _ := service.RegistrationRepository.FindByPatient(id)
	if schedule == nil {
		return nil, fmt.Errorf("registration not found")
	}

	return schedule, nil
}