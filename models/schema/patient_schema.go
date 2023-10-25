package schema

import (
	"time"

	"gorm.io/gorm"
)


type Patient struct {
	ID        uint      `gorm:"primaryKey"`
	CreatedAt time.Time `gorm:"autoCreateTime"`
	UpdatedAt time.Time `gorm:"autoUpdateTime"`
	DeletedAt gorm.DeletedAt `gorm:"index"`

	Name 	  	 string    `json:"name"`
	Email 		 string    `json:"email"`
	Password 	 string    `json:"password"`
	Nik 		 string    `json:"nik"`
<<<<<<< Updated upstream
	Birthday 	 string 	`json:"birthday"`
=======
<<<<<<< Updated upstream
	Birthday 	 time.Time `json:"birthday"`
>>>>>>> Stashed changes
	Age 		 uint64    `json:"age"`
=======
	Birthday 	 string    `json:"birthday"`
	Age 		 uint      `json:"age"`
>>>>>>> Stashed changes
	Address 	 string    `json:"address"`
	Gender 		 string    `gorm:"type:ENUM('MALE', 'FEMALE', 'UNKNOWN');not null;default:'UNKNOWN'"`
	Phone_number string    	`json:"phone_number"`
}