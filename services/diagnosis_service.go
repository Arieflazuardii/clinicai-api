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

type DiagnosisService interface {
	CreateDiagnosis(ctx echo.Context, request web.DiagnosisCreateRequest) (*domain.Diagnosis, error)
	FindDiagnosesById(ctx echo.Context, id int) (*domain.Diagnosis, error)
	FindAll(ctx echo.Context) ([]domain.Diagnosis, error)
}

type DiagnosisServiceImpl struct {
	DiagnosisRepository repository.DiagnosisRepository
	Validate *validator.Validate
}

func NewDiagnosisService(DiagnosisRepository repository.DiagnosisRepository, validate *validator.Validate) *DiagnosisServiceImpl {
	return &DiagnosisServiceImpl{
		DiagnosisRepository: DiagnosisRepository,
		Validate: validate,
	}
}


func (service *DiagnosisServiceImpl) CreateDiagnosis(ctx echo.Context, request web.DiagnosisCreateRequest) (*domain.Diagnosis, error) {
	err := service.Validate.Struct(request)
	if err != nil {
		return nil, helpers.ValidationError(ctx, err)
	}
	diagnosis := req.DiagnosisCreateRequestToDiagnosisDomain(request)

	result, err := service.DiagnosisRepository.Save(diagnosis)
	if err != nil {
		return nil, fmt.Errorf("error when creating diagnosis: %s", err.Error())
	}

	return result, nil
}


func (service *DiagnosisServiceImpl) FindDiagnosesById(ctx echo.Context, id int) (*domain.Diagnosis, error) {
	diagnosis, _ := service.DiagnosisRepository.FindById(id)
	if diagnosis == nil {
		return nil, fmt.Errorf("diagnosis not found")
	}

	return diagnosis, nil
}

func (service *DiagnosisServiceImpl) FindAll(ctx echo.Context) ([]domain.Diagnosis, error) {
	diagnosis, _ := service.DiagnosisRepository.FindAll()
	if diagnosis == nil {
		return nil, fmt.Errorf("diagnosis not found")
	}

	return diagnosis, nil
}