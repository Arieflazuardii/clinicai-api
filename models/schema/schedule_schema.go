package schema

import (
	"time"

	"gorm.io/gorm"
)

type Schedule struct {
	ID        uint         `gorm:"primaryKey"`
	CreatedAt time.Time      `gorm:"autoCreateTime"`
	UpdatedAt time.Time      `gorm:"autoUpdateTime"`
	DeletedAt gorm.DeletedAt `gorm:"index"`

	DoctorID    uint 		   `gorm:"index"`
	Doctor 	    Doctor 	 	   `gorm:"foreignKey:DoctorID"`
	Date   		time.Time 	   `json:"date"`
	Quota       int   		   `json:"quota"`
}