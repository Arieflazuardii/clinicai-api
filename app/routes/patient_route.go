package routes

import (
	"clinicai-api/controllers"
	"clinicai-api/repository"
	"clinicai-api/services"
	"os"

	"github.com/go-playground/validator"
	echoJwt "github.com/labstack/echo-jwt/v4"
	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
)

func PatientRoutes(e *echo.Echo, db *gorm.DB, validate *validator.Validate) {
	patientRepository :=repository.NewPatientRepository(db)
	patientService := services.NewPatientService(patientRepository, validate)
	PatientController := controllers.NewPatientController(patientService)

	patientGroup := e.Group("api/v1/patients")

	patientGroup.POST("/register", PatientController.RegisterPatientController)
	patientGroup.POST("/login", PatientController.LoginPatientController)

	patientGroup.Use(echoJwt.JWT([]byte(os.Getenv("JWT_SECRET_KEY"))))

	patientGroup.GET("/:id", PatientController.GetPatientController)
	patientGroup.GET("", PatientController.GetPatientsController)
	patientGroup.GET("/name/:name", PatientController.GetPatientByNameController)
	patientGroup.PUT("/:id", PatientController.UpdatePatientController)
	patientGroup.DELETE("/:id", PatientController.DeletePatientController)
}
