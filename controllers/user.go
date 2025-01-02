package controllers

import (
	"fmt"
	"time"
	"net/http"

	gomail "gopkg.in/gomail.v2"

	"github.com/golang-jwt/jwt"
	"github.com/labstack/echo/v4"

	"github.com/makotonakai/lets-schedule-api-test/config"
	"github.com/makotonakai/lets-schedule-api-test/models"
)

var maxIndex = 1
var mockDB = map[string]*models.User{}

type JWTCustomClaims struct {
	Id       int    `json:"uid"`
	UserName string `json:"name"`
	jwt.StandardClaims
}

type Credential struct {
	Id       int    `json:"id"`
	UserName string `json:"user_name"`
	Token string `json:"token"`
}

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
	mockDB[newUser.UserName] = &newUser
	maxIndex += 1

	return c.JSON(http.StatusCreated, newUser)

}

func Login(c echo.Context) error {

	u_ := models.User{}
	err := c.Bind(&u_)

	if err != nil {
		return c.JSON(http.StatusBadRequest, config.ErrFailedToBindUser)
	}

	pu := mockDB[u_.UserName]
	u := *pu

	if u.Id == 0 || u_.UserName != u.UserName || u_.Password != u.Password {
		return c.JSON(http.StatusUnauthorized, config.ErrLoginFailed)
	}

	claims := JWTCustomClaims{
		Id:       u.Id,
		UserName: u.UserName,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().Add(time.Hour).Unix(),
			Issuer:    "test",
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	signingKey := []byte("secret")
	t, err := token.SignedString(signingKey)

	return c.JSON(http.StatusOK, Credential{
		Id:        u.Id,
		UserName: u.UserName,
		Token:     t,
	})

}

func SendEmail(c echo.Context) error {

	ea := models.EmailAddress{}
	err := c.Bind(&ea)

	if err != nil {
		return c.JSON(http.StatusBadRequest, err.Error())
	}

	ui, err := models.GetUserIdFromEmailAddress(mockDB, ea.EmailAddress)
	if err != nil {
		return c.JSON(http.StatusBadRequest, err.Error())
	}

	m := gomail.NewMessage()
	m.SetHeader("From", "from@email.com")
	m.SetHeader("To", ea.EmailAddress)
	m.SetHeader("Subject", "Please reset your password for Let's Schedule")
	m.SetBody("text/plain", fmt.Sprintf("http://localhost:3000/%d/reset-password", ui))

	d := gomail.Dialer{Host: "localhost", Port: 25}
	err = d.DialAndSend(m)
	if err != nil {
		return c.JSON(http.StatusBadRequest, "There is some problem with sending your email")
	}

	return c.JSON(http.StatusOK, "Email was sent successfully!")
}

