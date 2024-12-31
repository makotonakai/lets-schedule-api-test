package controllers

import (
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/makotonakai/lets-schedule-api-test/config"
	"github.com/makotonakai/lets-schedule-api-test/models"
)

var maxIndex = 1
var mockDB = map[int]*models.User{}


func HealthCheck(c echo.Context) error {
	return c.JSON(http.StatusOK, "Accessible")
}

func CreateUser(c echo.Context) error {

	newUser := models.User{}
	err := c.Bind(&newUser)

	if err != nil {
		return c.JSON(http.StatusBadRequest, config.ErrFailedToBindUser)
	}

	newUser.Id = maxIndex
	mockDB[maxIndex] = &newUser
	maxIndex += 1

	return c.JSON(http.StatusCreated, newUser)

}
