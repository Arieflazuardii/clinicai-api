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

type ScheduleService interface {
	CreateSchedule(ctx echo.Context, request web.ScheduleCreateRequest) (*domain.Schedule, error)
	UpdateSchedule(ctx echo.Context, request web.ScheduleUpdateRequest, id int) (*domain.Schedule, error)
	FindById(ctx echo.Context, id int) (*domain.Schedule, error)
	FindAll(ctx echo.Context) ([]domain.Schedule, error)
	DeleteSchedule(ctx echo.Context, id int) error
}

type ScheduleServiceImpl struct {
	ScheduleRepository repository.ScheduleRepository
	Validate           *validator.Validate
}

func NewScheduleService(ScheduleRepository repository.ScheduleRepository, validate *validator.Validate) *ScheduleServiceImpl {
	return &ScheduleServiceImpl{
		ScheduleRepository: ScheduleRepository,
		Validate:           validate,
	}
}

func (service *ScheduleServiceImpl) CreateSchedule(ctx echo.Context, request web.ScheduleCreateRequest) (*domain.Schedule, error) {
	err := service.Validate.Struct(request)
	if err != nil {
		return nil, helpers.ValidationError(ctx, err)
	}

	schedule := req.ScheduleCreateRequestToScheduleDomain(request)

	result, err := service.ScheduleRepository.Create(schedule)
	if err != nil {
		return nil, fmt.Errorf("error when creating schedule: %s", err.Error())
	}

	return result, nil
}
func (service *ScheduleServiceImpl) UpdateSchedule(ctx echo.Context, request web.ScheduleUpdateRequest, id int) (*domain.Schedule, error) {
	err := service.Validate.Struct(request)
	if err != nil {
		return nil, helpers.ValidationError(ctx, err)
	}

	existingSchedule, _ := service.ScheduleRepository.FindById(id)
	if existingSchedule == nil {
		return nil, fmt.Errorf("schedule not found")
	}
	schedule := req.ScheduleUpdateRequestToScheduleDomain(request)

	result, err := service.ScheduleRepository.Update(schedule, id)
	if err != nil {
		return nil, fmt.Errorf("error when update schedule: %s", err.Error())
	}
	return result, nil
}

func (service *ScheduleServiceImpl) FindById(ctx echo.Context, id int) (*domain.Schedule, error) {
	schedule, _ := service.ScheduleRepository.FindById(id)
	if schedule == nil {
		return nil, fmt.Errorf("schedule not found")
	}

	return schedule, nil
}

func (service *ScheduleServiceImpl) FindAll(ctx echo.Context) ([]domain.Schedule, error) {
	schedule, _ := service.ScheduleRepository.FindAll()
	if schedule == nil {
		return nil, fmt.Errorf("schedule not found")
	}

	return schedule, nil
}

func (service *ScheduleServiceImpl) DeleteSchedule(ctx echo.Context, id int) error {
	schedule, _ := service.ScheduleRepository.FindById(id)
	if schedule == nil {
		return fmt.Errorf("book not found")
	}

	err := service.ScheduleRepository.Delete(id)
	if err != nil {
		return fmt.Errorf("error when deleting schedule: %s", err)
	}

	return nil
}
