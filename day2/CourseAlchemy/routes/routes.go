package routes

import (
	"CourseAlchemy/controllers/books"
	"CourseAlchemy/controllers/user"

	"github.com/labstack/echo/v4"
)

func New() *echo.Echo {
	e := echo.New()
	v1 := e.Group("/v1")

	bookGroup := v1.Group("/books")
	bookGroup.GET("/:id", books.GetBooks)
	bookGroup.GET("", books.ShowBooks)
	bookGroup.POST("", books.CreateBooks)
	bookGroup.PUT("/:id", books.UpdateBooks)
	bookGroup.DELETE("/:id", books.DeleteBooks)

	userGroup := v1.Group("/users")
	userGroup.GET("/:id", user.GetUser)
	userGroup.GET("", user.ShowUser)
	userGroup.POST("", user.CreateUser)
	userGroup.PUT("/:id", user.UpdateUser)
	userGroup.DELETE("/:id", user.DeleteUser)

	return e

}
