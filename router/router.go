package router

import (

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"

	"github.com/makotonakai/lets-schedule-api-test/controllers"
)

func Initialize() *echo.Echo {

	e := echo.New()
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())
	e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins: []string{"*"},
		AllowMethods: []string{echo.GET, echo.HEAD, echo.PUT, echo.PATCH, echo.POST, echo.DELETE},
	}))

	api := e.Group("/api")
	api.GET("/", controllers.HealthCheck)
	api.POST("/signup", controllers.CreateUser)
	api.POST("/login", controllers.Login)
	api.POST("/send-email", controllers.SendEmail)

	return e

}
