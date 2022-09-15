package routes

import (
	"CourseAlchemy/controllers/auth"
	"CourseAlchemy/controllers/books"
	"CourseAlchemy/controllers/user"
	m "CourseAlchemy/middleware"
	"os"

	echoMiddleware "github.com/labstack/echo/v4/middleware"

	"github.com/labstack/echo/v4"
)

func New() *echo.Echo {

	e := echo.New()
	v1 := e.Group("/v1")
	// e.Use(echoMid.BasicAuth(m.BasicAuth))

	v1.GET("/extract", func(c echo.Context) error {
		result := map[string]interface{}{
			"message": "failed",
		}
		payload := m.ExtractTokenPayload(c)
		if payload != nil {
			result["message"] = "success"
			result["token"] = m.ExtractTokenPayload(c)
		}

		return c.JSON(200, result)
	})

	v1.POST("/login", auth.Login)

	bookGroup := v1.Group("/books")
	bookGroup.GET("/:id", books.GetBooks)
	bookGroup.GET("", books.ShowBooks)
	bookGroup.POST("", books.CreateBooks)
	bookGroup.PUT("/:id", books.UpdateBooks)
	bookGroup.DELETE("/:id", books.DeleteBooks)

	userGroup := v1.Group("/users")
	userGroupWithAuth := userGroup.Group("/")
	userGroupWithAuth.Use(echoMiddleware.JWTWithConfig(echoMiddleware.JWTConfig{
		SigningKey:  []byte(os.Getenv("JWT_SECRET")),
		TokenLookup: `header:Authorization`,
	}))
	userGroup.GET("/:id", user.GetUser)
	userGroup.GET("", user.ShowUser)
	userGroup.POST("", user.CreateUser)
	userGroupWithAuth.PUT(":id", user.UpdateUser)
	userGroupWithAuth.DELETE(":id", user.DeleteUser)

	return e

}
