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

func RegistrationRoutes(e *echo.Echo, db *gorm.DB, validate *validator.Validate) {
	diagnosisRepository := repository.NewDiagnosisRepository(db)
	diagnosisService := services.NewDiagnosisService(diagnosisRepository, validate)

	registrationRepository := repository.NewRegistrationRepository(db)
	registrationService := services.NewRegistrationService(registrationRepository, validate)
	RegistrationController := controllers.NewRegistrationController(registrationService, diagnosisService)

	registrationGroup := e.Group("api/v1/registration")

	registrationGroup.Use(echoJwt.JWT([]byte(os.Getenv("SECRET_KEY"))))

	registrationGroup.POST("", RegistrationController.CreateRegistrationController, middleware.AuthMiddleware("Patient"))
	registrationGroup.GET("/:id", RegistrationController.GetRegistrationController, middleware.AuthMiddleware("Patient"))
	registrationGroup.GET("/patient", RegistrationController.GetRegistrationControllerByPatient, middleware.AuthMiddleware("Patient"))
	registrationGroup.GET("", RegistrationController.GetAllRegistrationController, middleware.AuthMiddleware("Doctor"))
	registrationGroup.PUT("/:id", RegistrationController.UpdateRegistrationController, middleware.AuthMiddleware("Patient"))
	registrationGroup.DELETE("/:id", RegistrationController.DeleteRegistrationController, middleware.AuthMiddleware("Patient"))
}