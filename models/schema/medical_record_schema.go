package schema

import (
	"time"

	"gorm.io/gorm"
)

type MedicalRecord struct {
	ID        uint         `gorm:"primaryKey"`
	CreatedAt time.Time      `gorm:"autoCreateTime"`
	UpdatedAt time.Time      `gorm:"autoUpdateTime"`
	DeletedAt gorm.DeletedAt `gorm:"index"`

	RegistrationID   uint 		   `gorm:"index"`
	Registration 	 Registration 	`gorm:"foreignKey:RegistrationID"`
	Symptomps 		 string 		`json:"symptomps"`
	Diagnoses 		 string 		`json:"diagnoses"`
	Solutions 		 string 		`json:"solution"`
}