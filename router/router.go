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
	e.Use(middleware.BodyDump(bodyDumpHandler))
	e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins: []string{"*"},
		AllowMethods: []string{echo.GET, echo.HEAD, echo.PUT, echo.PATCH, echo.POST, echo.DELETE},
	}))
	e.Use(session.Middleware(sessions.NewCookieStore([]byte("secret"))))

	api := e.Group("/api")
	api.GET("/", controllers.HealthCheck)
	api.POST("/signup", controllers.CreateUser)

}
