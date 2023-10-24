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

func DoctorRoutes(e *echo.Echo, db *gorm.DB, validate *validator.Validate) {
	doctorRepository :=repository.NewDoctorRepository(db)
	doctorService := services.NewDoctorService(doctorRepository, validate)
	DoctorController := controllers.NewDoctorController(doctorService)

	doctorGroup := e.Group("api/v1/doctors")

	doctorGroup.POST("/register", DoctorController.RegisterDoctorController)
	doctorGroup.POST("/login", DoctorController.LoginDoctorController)

	doctorGroup.Use(echoJwt.JWT([]byte(os.Getenv("SECRET_KEY"))))

	doctorGroup.GET("/:id", DoctorController.GetDoctorController)
	doctorGroup.GET("", DoctorController.GetDoctorsController)
	doctorGroup.GET("/name/:name", DoctorController.GetDoctorByNameController)
	doctorGroup.PUT("/:id", DoctorController.UpdateDoctorController)
	doctorGroup.DELETE("/:id", DoctorController.DeleteDoctorController)
}