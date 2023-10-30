package repository

import (
	"clinicai-api/models/domain"
	req "clinicai-api/utils/request"
	res "clinicai-api/utils/response"

	"gorm.io/gorm"
)

type DiagnosisRepository interface {
	Save(diagnosis *domain.Diagnosis) (*domain.Diagnosis, error)
	FindById(id int) (*domain.Diagnosis, error)
	FindAll() ([]domain.Diagnosis, error)
}

type DiagnosisRepositoryImpl struct {
	DB *gorm.DB
}

func NewDiagnosisRepository(db *gorm.DB) DiagnosisRepository{
	return &DiagnosisRepositoryImpl{DB: db}
}

func (repository *DiagnosisRepositoryImpl) Save(diagnosis *domain.Diagnosis) (*domain.Diagnosis, error) {
	diagnosisDB := req.DiagnosisDomainToDiagnosisSchema(*diagnosis)
	
	result := repository.DB.Create(&diagnosisDB)
	if result.Error != nil {
		return nil, result.Error
	}
	results := res.DiagnosisSchemaToDiagnosisDomain(diagnosisDB)

	return results, nil
}


func (repository *DiagnosisRepositoryImpl) FindById(id int) (*domain.Diagnosis, error) {
	var diagnosis domain.Diagnosis

	query := `SELECT *, 
	patients.name AS patient_name, 
	doctors.name AS doctor_name, 
	schedules.date AS schedule_date,
	complaint AS complaint 
	FROM diagnoses 
	LEFT JOIN registrations ON diagnoses.registration_id = registrations.id
	LEFT JOIN patients ON registrations.patient_id = patients.id
	LEFT JOIN doctors ON registrations.doctor_id = doctors.id
	LEFT JOIN schedules ON registrations.schedule_id = schedules.id
	WHERE registrations.id = ?`
	result := repository.DB.Raw(query, id).Scan(&diagnosis)

	if result.Error != nil {
		return nil, result.Error
	}

	return &diagnosis, nil
}

func (repository *DiagnosisRepositoryImpl) FindAll() ([]domain.Diagnosis, error) {
	diagnosis := []domain.Diagnosis{}
	query := "SELECT diagnosis.* FROM diagnosis"
	result := repository.DB.Raw(query).Scan(&diagnosis)
	if result.Error != nil {
		return nil, result.Error
	}

	return diagnosis, nil
}