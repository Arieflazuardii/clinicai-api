package routes

import (
	"clinicai-api/controllers"
	"clinicai-api/repository"
	"clinicai-api/services"
	"clinicai-api/utils/helpers/middleware"
	"os"

	"github.com/go-playground/validator"
	echoJwt "github.com/labstack/echo-jwt/v4"
	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
)

func MedicalRecordRoutes(e *echo.Echo, db *gorm.DB, validate *validator.Validate) {
	registrationRepository := repository.NewRegistrationRepository(db)
	registrationService := services.NewRegistrationService(registrationRepository, validate)

	medicalRecordRepository := repository.NewMedicalRecordRepository(db)
	medicalRecordService := services.NewMedicalRecordService(medicalRecordRepository, validate)
	MedicalRecordController := controllers.NewMedicalRecordController(medicalRecordService, registrationService)

	medicalRecordGroup := e.Group("api/v1/medicalrecord")

	medicalRecordGroup.Use(echoJwt.JWT([]byte(os.Getenv("SECRET_KEY"))))

	medicalRecordGroup.POST("", MedicalRecordController.CreateMedicalRecordController, middleware.AuthMiddleware("Doctor"))
	medicalRecordGroup.GET("/:id", MedicalRecordController.GetMedicalRecordController)
	medicalRecordGroup.GET("/patient", MedicalRecordController.GetMedicalRecordControllerByPatient, middleware.AuthMiddleware("Patient"))
	medicalRecordGroup.GET("", MedicalRecordController.GetAllMedicalRecordController)
	medicalRecordGroup.PUT("/:id", MedicalRecordController.UpdateMedicalRecordController, middleware.AuthMiddleware("Doctor"))
	medicalRecordGroup.DELETE("/:id", MedicalRecordController.DeleteMedicalRecordController, middleware.AuthMiddleware("Doctor"))
}