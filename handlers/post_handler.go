package handlers

import (
	"net/http"
	"unishare/models"
	"unishare/services"

	"github.com/labstack/echo"
)

func PostCreate(c echo.Context) error {
	classID := c.Param("classID")
	token := c.Request().Header.Get("access-token")

	post := new(models.Post)
	if err := c.Bind(&post); err != nil {
		panic(err)
	}

	info, err := services.PostCreate(post, token, classID)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}

	return c.JSON(http.StatusCreated, info)
}

func PostGet(c echo.Context) error {
	postID := c.Param("id")

	post, err := services.PostGet(postID)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}

	return c.JSON(http.StatusOK, post)
}

func PostUpdate(c echo.Context) error {
	postID := c.Param("id")

	post, err := services.PostGet(postID)
	if err != nil {
		return c.JSON(http.StatusNotFound, err.Error())
	}

	if err := c.Bind(post); err != nil {
		panic(err)
	}

	info, err := services.PostUpdate(post)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}

	return c.JSON(http.StatusOK, info)
}

func PostDelete(c echo.Context) error {
	postID := c.Param("id")

	info, err := services.PostDelete(postID)
	if err != nil {
		return c.JSON(http.StatusNotFound, err.Error())
	}

	return c.JSON(http.StatusOK, info)
}

func PostList(c echo.Context) error {
	classID := c.Param("classID")

	posts, err := services.PostList(classID)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}

	return c.JSON(http.StatusOK, posts)
}
