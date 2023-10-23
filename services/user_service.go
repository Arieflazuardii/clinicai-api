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



type PatientService interface {
	CreatePatient(ctx echo.Context, request web.PatientCreateRequest) (*domain.Patient, error)
	LoginPatient(ctx echo.Context, request web.PatientLoginRequest) (*domain.Patient, error)
	UpdatePatient(ctx echo.Context, request web.PatientUpdateRequest, id int) (*domain.Patient, error)
	FindById(ctx echo.Context, id int) (*domain.Patient, error)
	FindAll(ctx echo.Context) ([]domain.Patient, error)
	FindByName(ctx echo.Context, name string) (*domain.Patient, error)
	DeletePatient(ctx echo.Context, id int) error
}

type PatientServiceImpl struct {
	PatientRepository repository.PatientRepository
	Validate *validator.Validate
}

func NewPatientService(patientRepository repository.PatientRepository, validate *validator.Validate) *PatientServiceImpl {
	return &PatientServiceImpl{
		PatientRepository: patientRepository,
		Validate: validate,
	}
}

func (service *PatientServiceImpl) CreatePatient(ctx echo.Context, request web.PatientCreateRequest) (*domain.Patient, error) {
	err := service.Validate.Struct(request)
	if err != nil {
		return nil, helpers.ValidationError(ctx, err)
	}
	
	existingPatient, _ := service.PatientRepository.FindByEmail(request.Email)
	if existingPatient != nil {
		return nil, fmt.Errorf("email already exists")
	}
	patient := req.PatientCreateRequestToPatientDomain(request)

	patient.Password = helpers.HashPassword(patient.Password)

	result, err := service.PatientRepository.Create(patient)

	if err != nil {
		return nil, fmt.Errorf("error creating patient %s", err.Error())
	}

	return result, nil
}

func (service *PatientServiceImpl) LoginPatient(ctx echo.Context, request web.PatientLoginRequest) (*domain.Patient, error) {
	err := service.Validate.Struct(request)
	if err != nil {
		return nil, helpers.ValidationError(ctx, err)
	}

	existingPatient, err := service.PatientRepository.FindByEmail(request.Email)
	if err != nil {
		return nil, fmt.Errorf("invalid email or password")
	}

	patient := req.PatientLoginRequestToPatientDomain(request)

	err = helpers.ComparePassword(existingPatient.Password, patient.Password)
	if err != nil {
		return nil, fmt.Errorf("invalid email or password")
	}

	return existingPatient, nil

}

func (service * PatientServiceImpl) UpdatePatient(ctx echo.Context, request web.PatientUpdateRequest, id int) (*domain.Patient, error) {
	err := service.Validate.Struct(request)
	if err != nil {
		return nil, helpers.ValidationError(ctx, err)
	}

	existingPatient, _ := service.PatientRepository.FindById(id)
	if existingPatient == nil {
		return nil, fmt.Errorf("patient not found")
	}

	patient := req.PatientUpdateRequestToPatientDomain(request)
	patient.Password = helpers.HashPassword(patient.Password)
	result, err := service.PatientRepository.Update(patient, id)
	if err != nil {
		return nil, fmt.Errorf("error when updating data patient: %s", err.Error())
	}

	return result, nil
}

func (service *PatientServiceImpl) FindById(ctx echo.Context, id int) (*domain.Patient, error) {
	existingPatient, _ := service.PatientRepository.FindById(id)
	if existingPatient == nil {
		return nil, fmt.Errorf("patient not found")
	}

	return existingPatient, nil
}

func (service *PatientServiceImpl) FindAll(ctx echo.Context) ([]domain.Patient, error) {
	patients, err := service.PatientRepository.FindAll()
	if err != nil {
		return nil, fmt.Errorf("patients not found")
	}

	return patients, nil
}

func (service *PatientServiceImpl) FindByName(ctx echo.Context, name string) (*domain.Patient, error) {
	patient, _ := service.PatientRepository.FindByName(name)
	if patient == nil {
		return nil, fmt.Errorf("patient not found")
	}
	
	return patient, nil
}

func (service *PatientServiceImpl) DeletePatient(ctx echo.Context, id int) error {
	existingPatient, _ := service.PatientRepository.FindById(id)
	if existingPatient == nil {
		return fmt.Errorf("patient not found")
	}
	
	err := service.PatientRepository.Delete(id)
	if err != nil {
		return fmt.Errorf("error deleting patient: %s", err)
	}

	return nil
}
