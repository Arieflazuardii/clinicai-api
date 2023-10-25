package main

import (
	"clinicai-api/app/routes"
	"clinicai-api/drivers"
	"net/http"

	"github.com/go-playground/validator"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func main() {
	myApp := echo.New()
	validate := validator.New()
	DB := drivers.ConnectDB()

	myApp.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "Welcome to Clinic AI API Services")
	})

	routes.PatientRoutes(myApp, DB, validate)
	routes.DoctorRoutes(myApp, DB, validate)
	routes.ScheduleRoutes(myApp, DB, validate)


	myApp.Pre(middleware.RemoveTrailingSlash())
	myApp.Use(middleware.CORS())
	myApp.Use(middleware.LoggerWithConfig(
		middleware.LoggerConfig{
			Format: "method=${method}, uri=${uri}, status=${status}, time=${time_rfc3339}\n",
		},
	))
	

	myApp.Logger.Fatal(myApp.Start(":8080"))
}