package repository

import (
	"clinicai-api/models/domain"
	"clinicai-api/models/schema"
	req "clinicai-api/utils/request"
	res "clinicai-api/utils/response"
	"fmt"

	"gorm.io/gorm"
)

type PatientRepository interface {
	Create(patient *domain.Patient) (*domain.Patient, error)
	Update(patient *domain.Patient, id int) (*domain.Patient, error)
	FindById(id int) (*domain.Patient, error)
	FindByEmail(email string) (*domain.Patient, error)
	FindAll() ([]domain.Patient, error)
	FindByName(name string) (*domain.Patient, error)
	Delete(id int) error
}

type PatientRepositoryImpl struct {
	DB *gorm.DB
}

func NewPatientRepository(DB *gorm.DB) PatientRepository {
	return &PatientRepositoryImpl{DB: DB}
}

func (repository *PatientRepositoryImpl) Create(patient *domain.Patient) (*domain.Patient, error) {
	patientDB := req.PatientDomaintoPatientSchema(*patient)
	result := repository.DB.Create(&patientDB)
	if result.Error != nil {
		return nil, result.Error
	}
	results := res.PatientSchemaToPatientDomain(patientDB)
	
	return results, nil
}

func (repository *PatientRepositoryImpl) Update(patient *domain.Patient, id int) (*domain.Patient, error) {
	result := repository.DB.Table("patients").Where("id = ?", id).Updates(domain.Patient{		
		Name: 	      patient.Name,
		Email:        patient.Email,
		Password: 	  patient.Password,
		Nik:          patient.Nik,
		Birthday:     patient.Birthday,
		Age:          patient.Age,
		Address:      patient.Address,
		Gender:       patient.Gender,
		Phone_number: patient.Phone_number,})
	if result.Error != nil {
		return nil, result.Error
	}

	return patient, nil
}

func (repository *PatientRepositoryImpl) FindById(id int) (*domain.Patient, error) {
	patient := domain.Patient{}

	result := repository.DB.First(&patient, id)
	if result.Error != nil {
		return nil, result.Error
	}
	return &patient, nil
}

func (repository *PatientRepositoryImpl) FindByEmail(email string) (*domain.Patient, error) {
	patient := domain.Patient{}

	result := repository.DB.Where("email = ?", email).First(&patient)
	if result.Error != nil {
		return nil, result.Error
	}

	return &patient, nil
}

func (repository *PatientRepositoryImpl) FindAll() ([]domain.Patient, error) {
	patient := []domain.Patient{}

	result := repository.DB.Find(&patient)
	if result.Error != nil {
		return nil, result.Error
	}
	return patient, nil
}

func (repository *PatientRepositoryImpl) FindByName(name string) (*domain.Patient, error) {
	patient := domain.Patient{}

	result :=repository.DB.Where("LOWER(name) LIKE LOWER(?)", "%"+name+"%").First(&patient)

	if result.Error != nil {
		return nil, result.Error
	}
	return &patient, nil
}


func (repository *PatientRepositoryImpl) Delete(id int) (error) {
	result :=repository.DB.Delete(&schema.Patient{},id)
	fmt.Println(result)
	if result.Error != nil {
		return result.Error
	}
	return nil
}


