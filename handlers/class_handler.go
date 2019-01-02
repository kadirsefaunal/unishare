package handlers

import (
	"net/http"
	"unishare/models"
	"unishare/services"

	"github.com/labstack/echo"
)

func ClassCreate(c echo.Context) error {
	class := new(models.Class)
	if err := c.Bind(&class); err != nil {
		panic(err)
	}

	err := services.ClassCreate(class)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}

	return c.JSON(http.StatusCreated, "{status: true}")
}

func ClassDelete(c echo.Context) error {
	classID := c.Param("id")

	err := services.ClassDelete(classID)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}

	return c.JSON(http.StatusOK, "{status: true}")
}
