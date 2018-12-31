package routes

import (
	"unishare/handlers"
	"unishare/middlewares"

	"github.com/labstack/echo"
)

// InitRoutes creates routes
func InitRoutes(e *echo.Echo) {
	user := e.Group("/user")
	user.POST("", handlers.UserCreate).Name = "create-user"

	e.POST("/login", handlers.UserLogin).Name = "login"

	school := e.Group("/school", middlewares.CheckToken)
	school.POST("", handlers.SchoolCreate).Name = "create-school"
	school.DELETE("/:id", handlers.SchoolDelete).Name = "delete-school"
	school.GET("", handlers.SchoolList).Name = "list-school"
}
