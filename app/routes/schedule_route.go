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

func ScheduleRoutes(e *echo.Echo, db *gorm.DB, validate *validator.Validate) {
	scheduleRepository := repository.NewScheduleRepository(db)
	scheduleService := services.NewScheduleService(scheduleRepository, validate)
	ScheduleController := controllers.NewScheduleController(scheduleService)

	scheduleGroup := e.Group("api/v1/schedule")

	scheduleGroup.Use(echoJwt.JWT([]byte(os.Getenv("SECRET_KEY"))))

	scheduleGroup.POST("", ScheduleController.CreateScheduleController, middleware.AuthMiddleware("Doctor"))
	scheduleGroup.GET("/:id", ScheduleController.GetScheduleController, middleware.AuthMiddleware("Doctor"))
	scheduleGroup.GET("", ScheduleController.GetAllScheduleController, middleware.AuthMiddleware("Doctor"))
	scheduleGroup.PUT("/:id", ScheduleController.UpdateScheduleController, middleware.AuthMiddleware("Doctor"))
	scheduleGroup.DELETE("/:id", ScheduleController.DeleteScheduleController, middleware.AuthMiddleware("Doctor"))
}