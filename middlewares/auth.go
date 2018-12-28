package middlewares

import (
	"net/http"
	"unishare/services"

	"github.com/labstack/echo"
)

func CheckToken(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		token := c.Request().Header.Get("access-token")
		if "" == token {
			return c.JSON(http.StatusUnauthorized, "Token header not found!")
		}

		loggedIn := services.RedisCheckUser(token)

		if false == loggedIn {
			return c.JSON(http.StatusUnauthorized, "Unauthorized")
		}

		return next(c)
	}
}
