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

func AnswerGet(c echo.Context) error {
	answerID := c.Param("id")

	answer, err := services.AnswerGet(answerID)
	if err != nil {
		return c.JSON(http.StatusNotFound, err.Error())
	}

	return c.JSON(http.StatusOK, answer)
}

func AnswerList(c echo.Context) error {
	token := c.Request().Header.Get("access-token")

	answers, err := services.AnswerList(token)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}

	return c.JSON(http.StatusOK, answers)
}
