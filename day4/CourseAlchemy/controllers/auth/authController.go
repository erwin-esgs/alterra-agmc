package auth

import (
	"CourseAlchemy/config"
	m "CourseAlchemy/middleware"
	"CourseAlchemy/models"
	"fmt"
	"net/http"

	"github.com/labstack/echo/v4"
)

var table = "users"

func Login(c echo.Context) error {
	result := map[string]interface{}{
		"message": "failed",
	}

	var requestBody = make(map[string]interface{})
	err := c.Bind(&requestBody)
	if err != nil {
		fmt.Println(err)
		return err
	}
	if requestBody["email"] != nil && requestBody["password"] != nil {
		modelStruct := models.User{}
		db := config.InitDB()
		err := db.Table(table).Where(models.User{Email: requestBody["email"].(string), Password: requestBody["password"].(string)}).First(&modelStruct).Error
		if err == nil {
			params := map[string]interface{}{
				"name":  modelStruct.Name,
				"email": modelStruct.Email,
			}
			token, _ := m.CreateToken(params)
			result["message"] = "success"
			// result["code"] = http.StatusOK
			result["token"] = token

		}
	}
	return c.JSON(http.StatusOK, result)
}
