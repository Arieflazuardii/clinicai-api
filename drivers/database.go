package drivers

import (
	"clinicai-api/models/schema"
	"fmt"
	"log"
	"os"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB

func ConnectDB() {

	dsn := fmt.Sprintf("%v:%v@tcp(%v:%v)/%v?charset=utf8mb4&parseTime=True&loc=Local",
	os.Getenv("DB_USER"),
	os.Getenv("DB_PASSWORD"),
	os.Getenv("DB_HOST"),
	os.Getenv("DB_PORT"),
	os.Getenv("DB_NAME"))

	var errDB error
	DB, errDB = gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if errDB != nil {
		panic("failed connect to database")
	}

	Migrate()

	fmt.Println("Connected to database")

}

func Migrate() {
	err := DB.AutoMigrate(&schema.Patient{}, 
		&schema.Doctor{}, 
		&schema.Schedule{}, 
		&schema.Registration{}, 
		&schema.Diagnosis{}, 
		&schema.MedicalRecord{})
	if err != nil {
		log.Fatal("Failed to Migrate Database")
	}
	fmt.Println("Success Migrate Database")
}