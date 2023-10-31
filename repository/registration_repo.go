package repository

import (
	"clinicai-api/models/domain"
	"clinicai-api/models/schema"
	req "clinicai-api/utils/request"
	res "clinicai-api/utils/response"
	"errors"

	"gorm.io/gorm"
)

type RegistrationRepository interface {
	Create(registration *domain.Registration) (*domain.Registration, error)
	CreateDiagnosis(registration *domain.Diagnosis) (*domain.Diagnosis, error)
	Update(registration *domain.Registration, id int) (*domain.Registration, error)
	GetScheduleQuota(sheduleID int) (int,error)
	FindById(id int) (*domain.Registration, error)
	FindByPatient(id int) ([]domain.Registration, error)
	FindAll() ([]domain.Registration, error)
	Delete(id int) error
}

type RegistrationRepositoryImpl struct {
	DB *gorm.DB
}


func NewRegistrationRepository(db *gorm.DB) RegistrationRepository{
	return &RegistrationRepositoryImpl{DB: db}
}

func (repository *RegistrationRepositoryImpl) Create(registration *domain.Registration) (*domain.Registration, error) {
	registrationDb := req.RegistrationDomainToRegistrationSchema(*registration)

	tx := repository.DB.Begin()
	defer func() {
		if r := recover(); r != nil {
			tx.Rollback()
		}
	}()
	
	//Checking the Availability Schedule 
	var schedule schema.Schedule
	result := tx.First(&schedule, registration.ScheduleID)
	if schedule.ID == 0 {
		tx.Rollback()
		return nil, errors.New("invalid ScheduleID")
	}
	if result.Error != nil {
		tx.Rollback()
		return nil, result.Error
	}

	quotaToDecrease := 1
	
	result = tx.Model(&schema.Schedule{}).Where("ID = ? ", registration.ScheduleID).Update("quota", gorm.Expr("quota - ?", quotaToDecrease))
	if result.Error != nil {
		tx.Rollback()
		return nil, result.Error
	}

	if err := tx.Create(&registrationDb).Error; err != nil {
		tx.Rollback()
		return nil, result.Error
	}

	
	if err := tx.Commit().Error; err != nil {
		return nil, err
	}

	results := res.RegistrationSchemaToRegistrationDomain(registrationDb)
	return results, nil
}

func (repository *RegistrationRepositoryImpl) CreateDiagnosis(diagnosis *domain.Diagnosis) (*domain.Diagnosis, error) {
	
	diagnosisDB := req.DiagnosisDomainToDiagnosisSchema(*diagnosis)

	results := res.DiagnosisSchemaToDiagnosisDomain(diagnosisDB)

	return results, nil
}

func (repository *RegistrationRepositoryImpl) Update(registration *domain.Registration, id int) (*domain.Registration, error){
<<<<<<< Updated upstream
<<<<<<< Updated upstream
	result := repository.DB.Table("registrations").Where("id = ?", id).Updates(domain.Registration{PatientID: registration.PatientID, DoctorID: registration.DoctorID, ScheduleID: registration.ScheduleID, Complaint: registration.Complaint})
=======
=======
>>>>>>> Stashed changes
	result := repository.DB.Table("registrations").Where("id = ?", id).Updates(domain.Registration{
		PatientID: registration.PatientID, 
		DoctorID: registration.DoctorID, 
		ScheduleID: registration.ScheduleID, 
		Complaint: registration.Complaint})
<<<<<<< Updated upstream
>>>>>>> Stashed changes
=======
>>>>>>> Stashed changes
	if result.Error != nil {
		return nil, result.Error
	}
	return registration, nil
}

func (repository *RegistrationRepositoryImpl) GetScheduleQuota(scheduleID int) (int, error) {
	var scheduleQuota int

	query := "SELECT quota FROM schedules WHERE id = ?"

	result := repository.DB.Raw(query, scheduleID).Scan(&scheduleQuota)

	if result.Error != nil {
		return 0, result.Error
	}

	return scheduleQuota, nil
}

func (repository *RegistrationRepositoryImpl) FindById(id int) (*domain.Registration, error) {
	var registration domain.Registration

	if err := repository.DB.First(&registration, id).Error; err != nil {
		return nil, err
	}
	query := `SELECT registrations.*, 
	patients.name AS patient_name, 
	doctors.name AS doctor_name, 
	schedules.date AS schedule_date,
	registrations.complaint AS complaint
	FROM registrations 
	LEFT JOIN patients ON registrations.patient_id = patients.id
	LEFT JOIN doctors ON registrations.doctor_id = doctors.id
	LEFT JOIN schedules ON registrations.schedule_id = schedules.id
<<<<<<< Updated upstream
<<<<<<< Updated upstream
	WHERE registrations.id = ?`
=======
	WHERE registrations.id = ? AND registrations.deleted_at IS NULL`
>>>>>>> Stashed changes
=======
	WHERE registrations.id = ? AND registrations.deleted_at IS NULL`
>>>>>>> Stashed changes

	result := repository.DB.Raw(query, id).Scan(&registration)
	if result.Error != nil {
		return nil, result.Error
	}
	return &registration, nil
}

func (repository *RegistrationRepositoryImpl) FindAll() ([]domain.Registration, error) {
	registration := []domain.Registration{}
	query:= `SELECT registrations.*, 
	patients.name AS patient_name, 
	doctors.name AS doctor_name, 
	schedules.date AS schedule_date 
	FROM registrations 
	LEFT JOIN patients ON registrations.patient_id = patients.id
	LEFT JOIN doctors ON registrations.doctor_id = doctors.id
<<<<<<< Updated upstream
<<<<<<< Updated upstream
	LEFT JOIN schedules ON registrations.schedule_id = schedules.id`
=======
	LEFT JOIN schedules ON registrations.schedule_id = schedules.id
	WHERE registrations.deleted_at IS NULL`
>>>>>>> Stashed changes
=======
	LEFT JOIN schedules ON registrations.schedule_id = schedules.id
	WHERE registrations.deleted_at IS NULL`
>>>>>>> Stashed changes
	result := repository.DB.Raw(query).Scan(&registration)
	if result.Error != nil {
		return nil, result.Error
	}
	return registration, nil
	
}

func (repository *RegistrationRepositoryImpl) Delete(id int) error {
	result:= repository.DB.Delete(&schema.Registration{}, id)
	if result.Error != nil {
		return result.Error
	}
	return nil
}


func (repository *RegistrationRepositoryImpl) FindByPatient(id int) ([]domain.Registration, error) {
	var registration []domain.Registration

	if err := repository.DB.First(&registration, id).Error; err != nil {
		return nil, err
	}
	query := `SELECT registrations.*, 
	patients.name AS patient_name, 
	doctors.name AS doctor_name, 
	schedules.date AS schedule_date 
	FROM registrations 
	LEFT JOIN patients ON registrations.patient_id = patients.id
	LEFT JOIN doctors ON registrations.doctor_id = doctors.id
	LEFT JOIN schedules ON registrations.schedule_id = schedules.id
<<<<<<< Updated upstream
<<<<<<< Updated upstream
	where patients.id = (?)
=======
	where patients.id = (?) AND registrations.deleted_at IS NULL
>>>>>>> Stashed changes
=======
	where patients.id = (?) AND registrations.deleted_at IS NULL
>>>>>>> Stashed changes
	`
	result := repository.DB.Raw(query, id).Scan(&registration)

	if result.Error != nil {
		return nil, result.Error
	}

	return registration, nil
}