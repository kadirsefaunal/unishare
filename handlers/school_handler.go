package handlers

import (
	"net/http"
	"unishare/models"
	"unishare/services"

	"github.com/labstack/echo"
)

func SchoolCreate(c echo.Context) error {
	school := new(models.School)

	if err := c.Bind(school); err != nil {
		panic(err)
	}

	err := services.SchoolCreate(school)
	if err != nil {
		return c.JSON(http.StatusBadRequest, err.Error())
	}

	return c.JSON(http.StatusCreated, "{status: true}")
}
