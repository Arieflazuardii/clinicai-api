package schema

import (
	"time"

	"gorm.io/gorm"
)

type Schedule struct {
	gorm.Model

	DoctorID    uint 		   `gorm:"index"`
	Doctor 	    uint 	 	   `gorm:"foreignKey:DoctorID"`
	Date   		time.Time 	   `json:"date"`
	Quota       int   		   `json:"quota"`
}