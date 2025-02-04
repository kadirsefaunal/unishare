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

	class := school.Group("/:schoolID/class", middlewares.CheckToken)
	class.POST("", handlers.ClassCreate).Name = "create-class"
	class.DELETE("/:id", handlers.ClassDelete).Name = "delete-class"
	class.PUT("/:id", handlers.ClassUpdate).Name = "update-class"
	class.GET("/:id", handlers.ClassGet).Name = "get-class"
	class.GET("", handlers.ClassListGet).Name = "get-class-list"

	post := class.Group("/:classID/post", middlewares.CheckToken)
	post.POST("", handlers.PostCreate).Name = "create-post"
	post.GET("/:id", handlers.PostGet).Name = "get-post"
	post.PUT("/:id", handlers.PostUpdate).Name = "update-post"
	post.DELETE("/:id", handlers.PostDelete).Name = "delete-post"
	post.GET("", handlers.PostList).Name = "list-post"

	answer := post.Group("/:postID/answer", middlewares.CheckToken)
	answer.POST("", handlers.AnswerCreate).Name = "create-answer"
	answer.GET("/:id", handlers.AnswerGet).Name = "get-answer"
	answer.GET("", handlers.AnswerList).Name = "list-answer"
	answer.PUT("/:id", handlers.AnswerUpdate).Name = "update-answer"
	answer.DELETE("/:id", handlers.AnswerDelete).Name = "delete-answer"
}
