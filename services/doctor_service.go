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



type DoctorService interface {
	CreateDoctor(ctx echo.Context, request web.DoctorCreateRequest) (*domain.Doctor, error)
	LoginDoctor(ctx echo.Context, request web.DoctorLoginRequest) (*domain.Doctor, error)
	UpdateDoctor(ctx echo.Context, request web.DoctorUpdateRequest, id int) (*domain.Doctor, error)
	FindById(ctx echo.Context, id int) (*domain.Doctor, error)
	FindAll(ctx echo.Context) ([]domain.Doctor, error)
	FindByName(ctx echo.Context, name string) (*domain.Doctor, error)
	DeleteDoctor(ctx echo.Context, id int) error
}

type DoctorServiceImpl struct {
	DoctorRepository repository.DoctorRepository
	Validate *validator.Validate
}

func NewDoctorService(doctorRepository repository.DoctorRepository, validate *validator.Validate) *DoctorServiceImpl {
	return &DoctorServiceImpl{
		DoctorRepository: doctorRepository,
		Validate: validate,
	}
}

func (service *DoctorServiceImpl) CreateDoctor(ctx echo.Context, request web.DoctorCreateRequest) (*domain.Doctor, error) {
	err := service.Validate.Struct(request)
	if err != nil {
		return nil, helpers.ValidationError(ctx, err)
	}
	
	existingDoctor, _ := service.DoctorRepository.FindByEmail(request.Email)
	if existingDoctor != nil {
		return nil, fmt.Errorf("email already exists")
	}
	doctor := req.DoctorCreateRequestToDoctorDomain(request)

	doctor.Password = helpers.HashPassword(doctor.Password)
	result, err := service.DoctorRepository.Create(doctor)

	if err != nil {
		return nil, fmt.Errorf("error creating doctor %s", err.Error())
	}

	return result, nil
}

func (service *DoctorServiceImpl) LoginDoctor(ctx echo.Context, request web.DoctorLoginRequest) (*domain.Doctor, error) {
	err := service.Validate.Struct(request)
	if err != nil {
		return nil, helpers.ValidationError(ctx, err)
	}
	
	existingDoctor, err := service.DoctorRepository.FindByEmail(request.Email)
	if err != nil {
		return nil, fmt.Errorf("invalid email or password")
	}

	doctor := req.DoctorLoginRequestToDoctorDomain(request)

	err = helpers.ComparePassword(existingDoctor.Password, doctor.Password)
	if err != nil {
		return nil, fmt.Errorf("invalid email or password")
	}

	return existingDoctor, nil

}

func (service * DoctorServiceImpl) UpdateDoctor(ctx echo.Context, request web.DoctorUpdateRequest, id int) (*domain.Doctor, error) {
	err := service.Validate.Struct(request)
	if err != nil {
		return nil, helpers.ValidationError(ctx, err)
	}

	existingDoctor, _ := service.DoctorRepository.FindById(id)
	if existingDoctor == nil {
		return nil, fmt.Errorf("doctor not found")
	}

	doctor := req.DoctorUpdateRequestToDoctorDomain(request)
	doctor.Password = helpers.HashPassword(doctor.Password)
	result, err := service.DoctorRepository.Update(doctor, id)
	if err != nil {
		return nil, fmt.Errorf("error when updating data doctor: %s", err.Error())
	}

	return result, nil
}

func (service *DoctorServiceImpl) FindById(ctx echo.Context, id int) (*domain.Doctor, error) {
	existingDoctor, _ := service.DoctorRepository.FindById(id)
	if existingDoctor == nil {
		return nil, fmt.Errorf("doctor not found")
	}

	return existingDoctor, nil
}

func (service *DoctorServiceImpl) FindAll(ctx echo.Context) ([]domain.Doctor, error) {
	doctors, err := service.DoctorRepository.FindAll()
	if err != nil {
		return nil, fmt.Errorf("doctors not found")
	}

	return doctors, nil
}

func (service *DoctorServiceImpl) FindByName(ctx echo.Context, name string) (*domain.Doctor, error) {
	doctor, _ := service.DoctorRepository.FindByName(name)
	if doctor == nil {
		return nil, fmt.Errorf("doctor not found")
	}
	
	return doctor, nil
}

func (service *DoctorServiceImpl) DeleteDoctor(ctx echo.Context, id int) error {
	existingDoctor, _ := service.DoctorRepository.FindById(id)
	if existingDoctor == nil {
		return fmt.Errorf("doctor not found")
	}
	
	err := service.DoctorRepository.Delete(id)
	if err != nil {
		return fmt.Errorf("error deleting Doctor: %s", err)
	}

	return nil
}
