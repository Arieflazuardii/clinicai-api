package repository

import (
	"clinicai-api/models/domain"
	"clinicai-api/models/schema"
	req "clinicai-api/utils/request"
	res "clinicai-api/utils/response"
	"fmt"

	"gorm.io/gorm"
)

type DoctorRepository interface {
	Create(doctor *domain.Doctor) (*domain.Doctor, error)
	Update(doctor *domain.Doctor, id int) (*domain.Doctor, error)
	FindById(id int) (*domain.Doctor, error)
	FindByEmail(email string) (*domain.Doctor, error)
	FindAll() ([]domain.Doctor, error)
	FindByName(name string) (*domain.Doctor, error)
	Delete(id int) error
}

type DoctorRepositoryImpl struct {
	DB *gorm.DB
}

func NewDoctorRepository(DB *gorm.DB) DoctorRepository {
	return &DoctorRepositoryImpl{DB: DB}
}

func (repository *DoctorRepositoryImpl) Create(doctor *domain.Doctor) (*domain.Doctor, error) {
	doctorDB := req.DoctorDomaintoDoctorSchema(*doctor)
	result := repository.DB.Create(&doctorDB)
	if result.Error != nil {
		return nil, result.Error
	}
	results := res.DoctorSchemaToDoctorDomain(doctorDB)

	return results, nil
}

func (repository *DoctorRepositoryImpl) Update(doctor *domain.Doctor, id int) (*domain.Doctor, error) {
	result := repository.DB.Table("doctors").Where("id = ?", id).Updates(domain.Doctor{		
		Name: 	      doctor.Name,
		Email:        doctor.Email,
		Password: 	  doctor.Password,
		Nik:          doctor.Nik,
		Specialization: doctor.Specialization,
		Gender:       doctor.Gender,
		Phone_number: doctor.Phone_number,})
	if result.Error != nil {
		return nil, result.Error
	}

	return doctor, nil
}

func (repository *DoctorRepositoryImpl) FindById(id int) (*domain.Doctor, error) {
	doctor := domain.Doctor{}

	result := repository.DB.First(&doctor, id)
	if result.Error != nil {
		return nil, result.Error
	}
	return &doctor, nil
}

func (repository *DoctorRepositoryImpl) FindByEmail(email string) (*domain.Doctor, error) {
	doctor := domain.Doctor{}

	result := repository.DB.Where("email = ?", email).First(&doctor)
	if result.Error != nil {
		return nil, result.Error
	}

	return &doctor, nil
}

func (repository *DoctorRepositoryImpl) FindAll() ([]domain.Doctor, error) {
	doctor := []domain.Doctor{}

	result := repository.DB.Find(&doctor)
	if result.Error != nil {
		return nil, result.Error
	}
	return doctor, nil
}

func (repository *DoctorRepositoryImpl) FindByName(name string) (*domain.Doctor, error) {
	doctor := domain.Doctor{}

	result :=repository.DB.Where("LOWER(name) LIKE LOWER(?)", "%"+name+"%").First(&doctor)

	if result.Error != nil {
		return nil, result.Error
	}
	return &doctor, nil
}


func (repository *DoctorRepositoryImpl) Delete(id int) (error) {
	result :=repository.DB.Delete(&schema.Doctor{},id)
	fmt.Println(result)
	if result.Error != nil {
		return result.Error
	}
	return nil
}


