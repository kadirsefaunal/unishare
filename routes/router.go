package routes

import (
	"unishare/handlers"

	"github.com/labstack/echo"
)

// InitRoutes creates routes
func InitRoutes(e *echo.Echo) {
	user := e.Group("/user")
	user.POST("", handlers.UserCreate).Name = "create-user"
	
	e.POST("/login", handlers.UserLogin).Name = "login"
}
