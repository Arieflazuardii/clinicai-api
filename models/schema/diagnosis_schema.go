package schema

import (
	"time"

	"gorm.io/gorm"
)

type Diagnosis struct {
	ID        uint         `gorm:"primaryKey"`
	CreatedAt time.Time      `gorm:"autoCreateTime"`
	UpdatedAt time.Time      `gorm:"autoUpdateTime"`
	DeletedAt gorm.DeletedAt `gorm:"index"`
	
	RegistrationID uint   `gorm:"index"`
	Registration   Registration `gorm:"foreignKey:RegistrationID"`
	OpenAIResult   string `gorm:"open_ai_result"`
}