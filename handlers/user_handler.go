package handlers

import (
	"net/http"
	"unishare/models"
	"unishare/services"

	"github.com/labstack/echo"
)

// UserCreate creates user and return information.
func UserCreate(c echo.Context) error {
	user := new(models.User)
	if err := c.Bind(user); err != nil {
		panic(err)
	}

	err := services.UserCreate(user)
	if err != nil {
		return c.JSON(http.StatusBadRequest, err.Error())
	}

	return c.JSON(http.StatusCreated, "{status: true}")
}

func UserLogin(c echo.Context) error {
	loginUser := new(models.LoginUser)
	if err := c.Bind(loginUser); err != nil {
		panic(err)
	}

	token, err := services.UserLogin(loginUser)
	if err != nil {
		return c.JSON(http.StatusBadRequest, err.Error())
	}

	return c.JSON(http.StatusOK, token)
}
