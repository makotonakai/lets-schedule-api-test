package main

import (
	"net/http"
	"github.com/labstack/echo/v4"

	"github.com/makotonakai/lets-schedule-api-test/router"
)

func main() {

	e := router.Initialize()
	e.Logger.Fatal(e.Start(":8080"))

}
