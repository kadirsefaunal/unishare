package handlers

import (
	"net/http"
	"unishare/models"
	"unishare/services"

	"github.com/labstack/echo"
)

func PostCreate(c echo.Context) error {
	token := c.Request().Header.Get("access-token")
	post := new(models.Post)
	if err := c.Bind(&post); err != nil {
		panic(err)
	}

	err := services.PostCreate(post, token)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}

	return c.JSON(http.StatusCreated, "{status: true}")
}

func PostGet(c echo.Context) error {
	postID := c.Param("id")

	post, err := services.PostGet(postID)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}

	return c.JSON(http.StatusOK, post)
}
