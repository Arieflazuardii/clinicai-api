package repository

import (
	"clinicai-api/models/domain"
	"clinicai-api/models/schema"
	req "clinicai-api/utils/request"
	res "clinicai-api/utils/response"

	"gorm.io/gorm"
)

type MedicalRecordRepository interface {
	Create(medicalRecord *domain.MedicalRecord) (*domain.MedicalRecord, error)
	Update(medicalRecord *domain.MedicalRecord, id int) (*domain.MedicalRecord, error)
	FindById(id int) (*domain.MedicalRecord, error)
	FindAll() ([]domain.MedicalRecord, error)
	FindByPatient(id int) ([]domain.MedicalRecord, error)
	Delete(id int) error
}

type MedicalRecordRepositoryImpl struct {
	DB *gorm.DB
}

func NewMedicalRecordRepository(db *gorm.DB) MedicalRecordRepository{
	return &MedicalRecordRepositoryImpl{DB: db}
}

func (repository *MedicalRecordRepositoryImpl) Create(medicalRecord *domain.MedicalRecord) (*domain.MedicalRecord, error){
	//Transaction Database
	tx := repository.DB.Begin()
	if tx.Error != nil {
		return nil, tx.Error
	}

	medicalRecordDb := req.MedicalRecordDomainToMedicalRecordSchema(*medicalRecord)
	result := tx.Create(&medicalRecordDb)
	if result.Error != nil {
		tx.Rollback()
		return nil, result.Error
	}

	if err := tx.Commit().Error; err != nil {
		return nil, err
	}

	results := res.MedicalRecordSchemaToMedicalRecordDomain(medicalRecordDb)
	return results, nil
}

func (repository *MedicalRecordRepositoryImpl) Update(medicalRecord *domain.MedicalRecord, id int) (*domain.MedicalRecord, error) {
	result := repository.DB.Table("medical_records").Where("id = ?", id).Updates(domain.MedicalRecord{RegistrationID: medicalRecord.RegistrationID, Symptomps: medicalRecord.Symptomps, Diagnoses: medicalRecord.Diagnoses, Solutions: medicalRecord.Solutions})
	if result.Error != nil {
		return nil, result.Error
	}

	return medicalRecord, nil

}

func (repository *MedicalRecordRepositoryImpl) FindById(id int) (*domain.MedicalRecord, error) {
	var medicalRecord domain.MedicalRecord

	if err := repository.DB.First(&medicalRecord, id).Error; err != nil {
		return nil, err
	}
<<<<<<< Updated upstream
	query := "SELECT * FROM medical_records WHERE medical_records.id = ?"
=======
	query := "SELECT * FROM medical_records WHERE medical_records.id = ? AND medical_records.deleted_at IS NULL"
>>>>>>> Stashed changes
	result := repository.DB.Raw(query, id).Scan(&medicalRecord)

	if result.Error != nil {
		return nil, result.Error
	}

	return &medicalRecord, nil
}


func (repository *MedicalRecordRepositoryImpl) FindAll() ([]domain.MedicalRecord, error) {
	medicalRecord := []domain.MedicalRecord{}
	query := `SELECT medical_records.*, patients.name AS patient_name, doctors.name AS doctor_name
	FROM medical_records
	LEFT JOIN registrations ON medical_records.registration_id = registrations.id
	LEFT JOIN patients ON registrations.patient_id = patients.id
	LEFT JOIN doctors ON registrations.doctor_id = doctors.id
<<<<<<< Updated upstream
=======
	WHERE medical_records.deleted_at IS NULL
>>>>>>> Stashed changes
	`
	result := repository.DB.Raw(query).Scan(&medicalRecord)
	if result.Error != nil {
		return nil, result.Error
	}
	return medicalRecord, nil
}

func (repository *MedicalRecordRepositoryImpl) FindByPatient(id int) ([]domain.MedicalRecord, error) {
	var medicalRecord []domain.MedicalRecord

	query := `SELECT medical_records.*, 
	patients.name AS patient_name,
	doctors.name AS doctor_name
	FROM medical_records
	LEFT JOIN registrations ON medical_records.registration_id = registrations.id
	LEFT JOIN patients ON registrations.patient_id = patients.id
	LEFT JOIN doctors ON registrations.doctor_id = doctors.id
<<<<<<< Updated upstream
	WHERE medical_records.patients.id = (?)`
=======
	WHERE medical_records.patients.id = (?) AND medical_recoeds.deleted_at IS NULL`
>>>>>>> Stashed changes

	result := repository.DB.Raw(query, id).Scan(&medicalRecord)
	if result.Error != nil {
		return nil, result.Error
	}

	return medicalRecord, nil
}

func (repository *MedicalRecordRepositoryImpl) Delete(id int) error {
	result := repository.DB.Delete(&schema.MedicalRecord{},id)
	if result.Error != nil {
		return result.Error
	}
	return nil
}