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

	info, err := services.SchoolCreate(school)
	if err != nil {
		return c.JSON(http.StatusBadRequest, err.Error())
	}

	return c.JSON(http.StatusCreated, info)
}

func SchoolDelete(c echo.Context) error {
	schoolID := c.Param("id")

	info, err := services.SchoolDelete(schoolID)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}

	return c.JSON(http.StatusOK, info)
}

func SchoolList(c echo.Context) error {
	schools, err := services.SchoolList()
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}

	return c.JSON(http.StatusOK, schools)
}

func SchoolGet(c echo.Context) error {
	schoolID := c.Param("id")

	school, err := services.SchoolGet(schoolID)
	if err != nil {
		return c.JSON(http.StatusNotFound, err.Error())
	}

	return c.JSON(http.StatusOK, school)
}
