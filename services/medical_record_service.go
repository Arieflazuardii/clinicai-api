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

type MedicalRecordService interface {
	CreateMedicalRecord(ctx echo.Context, request web.MedicalRecordRequest) (*domain.MedicalRecord, error)
	UpdateMedicalRecord(ctx echo.Context, request web.MedicalRecordUpdateRequest, id int) (*domain.MedicalRecord, error)
	FindById(ctx echo.Context, id int) (*domain.MedicalRecord, error)
	FindAll() ([]domain.MedicalRecord, error)
	FindByPatient(id int) ([]domain.MedicalRecord, error)
	DeleteMedicalRecord(ctx echo.Context, id int) error
}

type MedicalRecordServiceImpl struct {
	MedicalRecordRepository repository.MedicalRecordRepository
	Validate           *validator.Validate
}

func NewMedicalRecordService(MedicalRecordRepository repository.MedicalRecordRepository, validate *validator.Validate) *MedicalRecordServiceImpl {
	return &MedicalRecordServiceImpl{
		MedicalRecordRepository: MedicalRecordRepository,
		Validate:           validate,
	}
}

func (service *MedicalRecordServiceImpl) CreateMedicalRecord(ctx echo.Context, request web.MedicalRecordRequest) (*domain.MedicalRecord, error) {
	err := service.Validate.Struct(request)
	if err != nil {
		return nil, helpers.ValidationError(ctx, err)
	}

	medicalRecord := req.MedicalRecordCreateRequestToMedicalRecordDomain(request)

	result, err := service.MedicalRecordRepository.Create(medicalRecord)
	if err != nil {
		return nil, fmt.Errorf("error when creating medicalRecord: %s", err.Error())
	}

	return result, nil
}

func (service *MedicalRecordServiceImpl) UpdateMedicalRecord(ctx echo.Context, request web.MedicalRecordUpdateRequest, id int) (*domain.MedicalRecord, error) {
	err := service.Validate.Struct(request)
	if err != nil {
		return nil, helpers.ValidationError(ctx, err)
	}

	existingMedicalRecord, _ := service.MedicalRecordRepository.FindById(id)
	if existingMedicalRecord == nil {
		return nil, fmt.Errorf("medical record not found")
	}
	medicalRecord := req.MedicalRecordUpdateRequestToMedicalRecordDomain(request)

	result, err := service.MedicalRecordRepository.Update(medicalRecord, id)
	if err != nil {
		return nil, fmt.Errorf("error when update medicalRecord: %s", err.Error())
	}
	return result, nil
}

func (service *MedicalRecordServiceImpl) FindById(ctx echo.Context, id int) (*domain.MedicalRecord, error) {
	medicalRecord, _ := service.MedicalRecordRepository.FindById(id)
	if medicalRecord == nil {
		return nil, fmt.Errorf("medical record not found")
	}

	return medicalRecord, nil
}

func (service *MedicalRecordServiceImpl) FindByPatient(id int) ([]domain.MedicalRecord, error) {
	medicalRecord, _ := service.MedicalRecordRepository.FindByPatient(id)
	if medicalRecord == nil {
		return nil, fmt.Errorf("medical record not found")
	}

	return medicalRecord, nil
}


func (service *MedicalRecordServiceImpl) FindAll() ([]domain.MedicalRecord, error) {
	medicalRecord, _ := service.MedicalRecordRepository.FindAll()
	if medicalRecord == nil {
		return nil, fmt.Errorf("medical record not found")
	}

	return medicalRecord, nil
}

func (service *MedicalRecordServiceImpl) DeleteMedicalRecord(ctx echo.Context, id int) error {
	medicalRecord, _ := service.MedicalRecordRepository.FindById(id)
	if medicalRecord == nil {
		return fmt.Errorf("medical record not found")
	}

	err := service.MedicalRecordRepository.Delete(id)
	if err != nil {
		return fmt.Errorf("error when deleting medical record: %s", err)
	}

	return nil
}
