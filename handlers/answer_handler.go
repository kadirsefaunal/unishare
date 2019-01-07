package handlers

import (
	"net/http"
	"unishare/models"
	"unishare/services"

	"github.com/labstack/echo"
)

func AnswerCreate(c echo.Context) error {
	token := c.Request().Header.Get("access-token")
	answer := new(models.Answer)

	if err := c.Bind(answer); err != nil {
		panic(err)
	}

	err := services.AnswerInsert(answer, token)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}

	return c.JSON(http.StatusCreated, "{status: true}")
}
