package handlers

import (
	"net/http"
	"unishare/models"
	"unishare/services"

	"github.com/labstack/echo"
)

func ClassCreate(c echo.Context) error {
	schoolID := c.Param("schoolID")

	class := new(models.Class)
	if err := c.Bind(&class); err != nil {
		panic(err)
	}

	info, err := services.ClassCreate(schoolID, class)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}

	return c.JSON(http.StatusCreated, info)
}

func ClassDelete(c echo.Context) error {
	classID := c.Param("id")

	info, err := services.ClassDelete(classID)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}

	return c.JSON(http.StatusOK, info)
}

func ClassUpdate(c echo.Context) error {
	classID := c.Param("id")

	class, err := services.ClassGet(classID)
	if err != nil {
		return c.JSON(http.StatusNotFound, err.Error())
	}

	if err := c.Bind(class); err != nil {
		panic(err)
	}

	info, err := services.ClassUpdate(class)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}

	return c.JSON(http.StatusOK, info)
}

func ClassGet(c echo.Context) error {
	classID := c.Param("id")

	class, err := services.ClassGet(classID)
	if err != nil {
		return c.JSON(http.StatusNotFound, err.Error())
	}

	return c.JSON(http.StatusOK, class)
}

func ClassListGet(c echo.Context) error {
	schoolID := c.Param("schoolID")

	classes, err := services.ClassList(schoolID)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}

	return c.JSON(http.StatusOK, classes)
}
