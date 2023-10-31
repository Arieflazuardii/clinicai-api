package repository

import (
	"clinicai-api/models/domain"
	"clinicai-api/models/schema"
	req "clinicai-api/utils/request"
	res "clinicai-api/utils/response"

	"gorm.io/gorm"
)

type ScheduleRepository interface {
	Create(schedule *domain.Schedule) (*domain.Schedule, error)
	Update(schedule *domain.Schedule, id int) (*domain.Schedule, error)
	FindById(id int) (*domain.Schedule, error)
	FindByDoctor(id int) ([]domain.Schedule, error)
	GetScheduleQuota(scheduleID int) (int, error)
	FindAll() ([]domain.Schedule, error)
	Delete(id int) error
}

type ScheduleRepositoryImpl struct {
	DB *gorm.DB
}

func NewScheduleRepository(db *gorm.DB) ScheduleRepository{
	return &ScheduleRepositoryImpl{DB: db}
}

func (repository *ScheduleRepositoryImpl) Create(schedule *domain.Schedule) (*domain.Schedule, error){
	//Transaction Database
	tx := repository.DB.Begin()
	if tx.Error != nil {
		return nil, tx.Error
	}

	scheduleDb := req.ScheduleDomainToScheduleSchema(*schedule)
	result := tx.Create(&scheduleDb)
	if result.Error != nil {
		tx.Rollback()
		return nil, result.Error
	}

	if err := tx.Commit().Error; err != nil {
		return nil, err
	}

	results := res.ScheduleSchemaToScheduleDomain(scheduleDb)
	return results, nil
}

func (repository *ScheduleRepositoryImpl) Update(schedule *domain.Schedule, id int) (*domain.Schedule, error) {
	result := repository.DB.Table("schedules").Where("id = ?", id).Updates(domain.Schedule{DoctorID: schedule.DoctorID, Date: schedule.Date, Quota: schedule.Quota})
	if result.Error != nil {
		return nil, result.Error
	}

	return schedule, nil

}

func (repository *ScheduleRepositoryImpl) FindById(id int) (*domain.Schedule, error) {
	var schedule domain.Schedule

	if err := repository.DB.First(&schedule, id).Error; err != nil {
		return nil, err
	}
<<<<<<< Updated upstream
	query := `SELECT schedules.*, doctors.name AS doctor_name
	FROM schedules 
	LEFT JOIN doctors ON schedules.doctor_id = doctors.id
	WHERE schedules.id = ?`

=======
<<<<<<< Updated upstream
	query := "SELECT * FROM schedules WHERE schedules.id = ?"
=======
	query := `SELECT schedules.*, doctors.name AS doctor_name
	FROM schedules 
	LEFT JOIN doctors ON schedules.doctor_id = doctors.id
	WHERE schedules.id = ? AND schedules.deleted_at IS NULL`

>>>>>>> Stashed changes
>>>>>>> Stashed changes
	result := repository.DB.Raw(query, id).Scan(&schedule)

	if result.Error != nil {
		return nil, result.Error
	}

	return &schedule, nil
}

<<<<<<< Updated upstream
func (repository *ScheduleRepositoryImpl) FindByDoctor(id int) ([]domain.Schedule, error) {
	var schedule []domain.Schedule

	if err := repository.DB.First(&schedule, id).Error; err != nil {
		return nil, err
	}
	query := `SELECT schedules.*, doctors.name AS doctor_name
	FROM schedules 
	LEFT JOIN doctors ON schedules.doctor_id = doctors.id
	where doctors.id = (?)
	`
	result := repository.DB.Raw(query, id).Scan(&schedule)

	if result.Error != nil {
		return nil, result.Error
	}

	return schedule, nil
}

func (repository *ScheduleRepositoryImpl) GetScheduleQuota(scheduleID int) (int, error) {
	var scheduleQuota int
	query := "SELECT quota FROM schedules WHERE id = ?"
=======
<<<<<<< Updated upstream
func (repository *ScheduleRepositoryImpl) GetScheduleQuota(scheduleID int) (int, error) {
	var scheduleQuota int
	query := "SELECT quota FROM schedule WHERE id = ?"
=======
func (repository *ScheduleRepositoryImpl) FindByDoctor(id int) ([]domain.Schedule, error) {
	var schedule []domain.Schedule

	if err := repository.DB.First(&schedule, id).Error; err != nil {
		return nil, err
	}
	query := `SELECT schedules.*, doctors.name AS doctor_name
	FROM schedules 
	LEFT JOIN doctors ON schedules.doctor_id = doctors.id
	where doctors.id = (?) AND schedules.deleted_at IS NULL
	`
	result := repository.DB.Raw(query, id).Scan(&schedule)

	if result.Error != nil {
		return nil, result.Error
	}

	return schedule, nil
}

func (repository *ScheduleRepositoryImpl) GetScheduleQuota(scheduleID int) (int, error) {
	var scheduleQuota int
	query := "SELECT quota FROM schedules WHERE id = (?) AND schedules.deleted_at IS NULL"
>>>>>>> Stashed changes
>>>>>>> Stashed changes
	result := repository.DB.Raw(query, scheduleID).Scan(&scheduleQuota)

	if result.Error != nil {
		return 0, result.Error
	}

	return scheduleQuota, nil
}

func (repository *ScheduleRepositoryImpl) FindAll() ([]domain.Schedule, error) {
	schedule := []domain.Schedule{}
<<<<<<< Updated upstream
	query := `SELECT schedules.*, doctors.name AS doctor_name
	FROM schedules 
	LEFT JOIN doctors ON schedules.doctor_id = doctors.id`
=======
<<<<<<< Updated upstream
	query := "SELECT schedules.* FROM schedules"
=======
	query := `SELECT schedules.*, doctors.name AS doctor_name
	FROM schedules 
	LEFT JOIN doctors ON schedules.doctor_id = doctors.id
	WHERE schedules.deleted_at IS NULL`
>>>>>>> Stashed changes
>>>>>>> Stashed changes
	result := repository.DB.Raw(query).Scan(&schedule)
	if result.Error != nil {
		return nil, result.Error
	}

	return schedule, nil
}

func (repository *ScheduleRepositoryImpl) Delete(id int) error {
	result := repository.DB.Delete(&schema.Schedule{},id)
	if result.Error != nil {
		return result.Error
	}
	return nil
}