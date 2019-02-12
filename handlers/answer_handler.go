package handlers

import (
	"net/http"
	"unishare/models"
	"unishare/services"

	"github.com/labstack/echo"
)

func AnswerCreate(c echo.Context) error {
	postID := c.Param("postID")
	token := c.Request().Header.Get("access-token")

	answer := new(models.Answer)
	if err := c.Bind(answer); err != nil {
		panic(err)
	}

	info, err := services.AnswerInsert(answer, token, postID)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}

	return c.JSON(http.StatusCreated, info)
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
	questionID := c.Param("postID")

	answers, err := services.AnswerList(questionID)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}

	return c.JSON(http.StatusOK, answers)
}

func AnswerUpdate(c echo.Context) error {
	answerID := c.Param("id")

	answer, err := services.AnswerGet(answerID)
	if err != nil {
		return c.JSON(http.StatusNotFound, err.Error())
	}

	if err := c.Bind(answer); err != nil {
		panic(err)
	}

	info, err := services.AnswerUpdate(answer)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}

	return c.JSON(http.StatusOK, info)
}

func AnswerDelete(c echo.Context) error {
	answerID := c.Param("id")

	info, err := services.AnswerDelete(answerID)
	if err != nil {
		return c.JSON(http.StatusNotFound, err.Error())
	}

	return c.JSON(http.StatusOK, info)
}
