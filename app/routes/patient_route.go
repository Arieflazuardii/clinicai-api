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

func PatientRoutes(e *echo.Echo, db *gorm.DB, validate *validator.Validate) {
	patientRepository :=repository.NewPatientRepository(db)
	patientService := services.NewPatientService(patientRepository, validate)
	PatientController := controllers.NewPatientController(patientService)

	patientGroup := e.Group("api/v1/patients")

	patientGroup.POST("/register", PatientController.RegisterPatientController)
	patientGroup.POST("/login", PatientController.LoginPatientController)

	patientGroup.Use(echoJwt.JWT([]byte(os.Getenv("SECRET_KEY"))))

	patientGroup.GET("/name/:name", PatientController.GetPatientByNameController, middleware.AuthMiddleware("Patient"))
	patientGroup.GET("/:id", PatientController.GetPatientController, middleware.AuthMiddleware("Patient"))
	patientGroup.GET("", PatientController.GetPatientsController, middleware.AuthMiddleware("Patient"))
	patientGroup.PUT("/:id", PatientController.UpdatePatientController, middleware.AuthMiddleware("Patient"))
	patientGroup.DELETE("/:id", PatientController.DeletePatientController, middleware.AuthMiddleware("Patient"))
}
