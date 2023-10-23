package schema

import (
	"time"

	"gorm.io/gorm"
)


type Patient struct {
	ID        uint64      `gorm:"primaryKey"`
	CreatedAt time.Time `gorm:"autoCreateTime"`
	UpdatedAt time.Time `gorm:"autoUpdateTime"`
	DeletedAt gorm.DeletedAt `gorm:"index"`

	Name 	  	 string    `json:"name"`
	Email 		 string    `json:"email"`
	Password 	 string    `json:"password"`
	Nik 		 string    `json:"nik"`
	Birthday 	 time.Time `json:"birthday"`
	Age 		 uint64    `json:"age"`
	Address 	 string    `json:"address"`
	Gender 		 string    `gorm:"type:ENUM('MALE', 'FEMALE', 'UNKNOWN');not null;default:'UNKNOWN'"`
	Phone_number string    	`json:"phone_number"`
}