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

func DoctorRoutes(e *echo.Echo, db *gorm.DB, validate *validator.Validate) {
	doctorRepository :=repository.NewDoctorRepository(db)
	doctorService := services.NewDoctorService(doctorRepository, validate)
	DoctorController := controllers.NewDoctorController(doctorService)

	doctorGroup := e.Group("api/v1/doctors")

	doctorGroup.POST("/register", DoctorController.RegisterDoctorController)
	doctorGroup.POST("/login", DoctorController.LoginDoctorController)

	doctorGroup.Use(echoJwt.JWT([]byte(os.Getenv("SECRET_KEY"))))

	doctorGroup.GET("/:id", DoctorController.GetDoctorController, middleware.AuthMiddleware("Doctor"))
	doctorGroup.GET("", DoctorController.GetDoctorsController, middleware.AuthMiddleware("Doctor"))
	doctorGroup.GET("/name/:name", DoctorController.GetDoctorByNameController, middleware.AuthMiddleware("Doctor"))
	doctorGroup.PUT("/:id", DoctorController.UpdateDoctorController, middleware.AuthMiddleware("Doctor"))
	doctorGroup.DELETE("/:id", DoctorController.DeleteDoctorController, middleware.AuthMiddleware("Doctor"))
}