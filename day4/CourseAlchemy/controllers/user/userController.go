package user

import (
	"CourseAlchemy/config"
	"CourseAlchemy/models"
	"fmt"
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
)

var table = "users"

func GetUser(c echo.Context) error {
	result := map[string]interface{}{
		"message": "failed",
	}
	par := c.Param("id")
	intPar, err := strconv.Atoi(par)
	modelName := models.User{}
	db := config.InitDB()
	err = db.Table(table).First(&modelName, intPar).Error
	if err == nil {
		result["message"] = "success"
		result["code"] = http.StatusOK
		result["data"] = modelName

	}
	return c.JSON(http.StatusOK, result)
}

func ShowUser(c echo.Context) error {
	db := config.InitDB()
	arrModel := []models.User{}
	db.Table(table).Find(&arrModel).Scan(&arrModel)
	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "success",
		"code":    http.StatusOK,
		"data":    arrModel,
	})
}

func CreateUser(c echo.Context) error {
	var jsonData = make(map[string]interface{})
	err := c.Bind(&jsonData)
	if err != nil {
		fmt.Println(err)
		return err

	}
	db := config.InitDB()
	modelName := models.User{Name: "",
		Email:    "",
		Password: ""}
	if jsonData["name"] != nil {
		modelName.Name = jsonData["name"].(string)
	}
	if jsonData["email"] != nil {
		modelName.Email = jsonData["name"].(string)
	}
	if jsonData["password"] != nil {
		modelName.Password = jsonData["password"].(string)
	}

	db.Table(table).Create(&modelName)
	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "success",
		"code":    http.StatusOK,
		"data":    modelName,
	})
}

func UpdateUser(c echo.Context) error {
	result := map[string]interface{}{
		"message": "failed",
	}
	par := c.Param("id")
	intPar, err := strconv.Atoi(par)
	db := config.InitDB()
	modelName := models.User{}
	db.Table(table).First(&modelName, intPar)
	if (modelName == models.User{}) {
		return c.JSON(http.StatusOK, result)
	}

	var jsonData = make(map[string]interface{})
	err = c.Bind(&jsonData)
	if err != nil {
		fmt.Println(err)
		return c.JSON(http.StatusOK, result)

	}
	if jsonData["name"] != nil {
		modelName.Name = jsonData["name"].(string)
	}
	if jsonData["email"] != nil {
		modelName.Email = jsonData["email"].(string)
	}
	if jsonData["password"] != nil {
		modelName.Password = jsonData["password"].(string)
	}
	saveResult := db.Table(table).Save(&modelName)

	if saveResult.RowsAffected > 0 {
		result["message"] = "success"
		result["code"] = http.StatusOK
		result["data"] = modelName
	}
	return c.JSON(http.StatusOK, result)
}

func DeleteUser(c echo.Context) error {
	result := map[string]interface{}{
		"message": "failed",
	}
	par := c.Param("id")
	modelName := models.User{}
	u64, err := strconv.ParseUint(par, 10, 64)
	if err != nil {
		fmt.Println(err)
	}
	modelName.ID = uint(u64)

	db := config.InitDB()
	bookModel := models.User{}
	saveResult := db.Table(table).Delete(&bookModel, uint(u64))
	if err != nil {
		return err
	}
	if saveResult.RowsAffected > 0 {
		result["message"] = "success"
		result["code"] = http.StatusOK
		result["data"] = modelName
	}
	return c.JSON(http.StatusOK, result)
}
