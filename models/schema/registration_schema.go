package schema

import (
	"time"

	"gorm.io/gorm"
)

type Registration struct {
	ID        uint         `gorm:"primaryKey"`
	CreatedAt time.Time      `gorm:"autoCreateTime"`
	UpdatedAt time.Time      `gorm:"autoUpdateTime"`
	DeletedAt gorm.DeletedAt `gorm:"index"`
	
	PatientID  uint   `gorm:"index"`
	Patient    Patient   `gorm:"foreignKey:PatientID"`
	DoctorID   uint   `gorm:"index"`
	Doctor     Doctor   `gorm:"foreignKey:DoctorID"`
	ScheduleID uint   `gorm:"index"`
	Schedule   Schedule   `gorm:"foreignKey:ScheduleID"`
	Complaint  string `json:"complaint"`
}