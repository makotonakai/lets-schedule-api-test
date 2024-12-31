package main

import (
	"github.com/makotonakai/lets-schedule-api-test/router"
)

func main() {

	e := router.Initialize()
	e.Logger.Fatal(e.Start(":8080"))
	
}
