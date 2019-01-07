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
	school.GET("/:id", handlers.SchoolGet).Name = "get-school"

	class := e.Group("/class", middlewares.CheckToken)
	class.POST("", handlers.ClassCreate).Name = "create-class"
	class.DELETE("/:id", handlers.ClassDelete).Name = "delete-class"
	class.PUT("/:id", handlers.ClassUpdate).Name = "update-class"
	class.GET("/:id", handlers.ClassGet).Name = "get-class"

	post := e.Group("/post", middlewares.CheckToken)
	post.POST("", handlers.PostCreate).Name = "create-post"
	post.GET("/:id", handlers.PostGet).Name = "get-post"
	post.PUT("/:id", handlers.PostUpdate).Name = "update-post"
	post.DELETE("/:id", handlers.PostDelete).Name = "delete-post"
	post.GET("", handlers.PostList).Name = "list-post"

	answer := e.Group("/answer", middlewares.CheckToken)
	answer.POST("", handlers.AnswerCreate).Name = "create-answer"
}
