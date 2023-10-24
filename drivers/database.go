package drivers

import (
	"clinicai-api/models/schema"
	"fmt"
	"os"
	"path/filepath"

	"github.com/joho/godotenv"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB

func ConnectDB() *gorm.DB {
	err := godotenv.Load(filepath.Join(".", ".env"))
	if err != nil {
		fmt.Println("error loading .env file: ")
		os.Exit(1)
	}

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

	return DB
}

func Migrate() {
	DB.AutoMigrate(schema.Patient{}, schema.Doctor{})
}