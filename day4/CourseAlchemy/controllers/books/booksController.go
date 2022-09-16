package books

import (
	"CourseAlchemy/config"
	"CourseAlchemy/models"
	"fmt"
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
)

var table = "books"

func GetBooks(c echo.Context) error {
	result := map[string]interface{}{
		"message": "failed",
	}
	par := c.Param("id")
	intPar, err := strconv.Atoi(par)
	modelName := models.Books{}
	db := config.InitDB()
	err = db.Table(table).First(&modelName, intPar).Error
	if err == nil {
		result["message"] = "success"
		result["code"] = http.StatusOK
		result["data"] = modelName

	}
	return c.JSON(http.StatusOK, result)
}

func ShowBooks(c echo.Context) error {
	db := config.InitDB()
	arrModel := []models.Books{}
	db.Table(table).Find(&arrModel).Scan(&arrModel)
	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "success",
		"code":    http.StatusOK,
		"data":    arrModel,
	})
}

func CreateBooks(c echo.Context) error {
	var jsonData = make(map[string]interface{})
	err := c.Bind(&jsonData)
	if err != nil {
		fmt.Println(err)
		return err

	}
	db := config.InitDB()
	modelName := models.Books{Name: "",
		Author:    "",
		Publisher: "",
		Year:      0,
	}
	if jsonData["name"] != nil {
		modelName.Name = jsonData["name"].(string)
	}
	if jsonData["email"] != nil {
		modelName.Author = jsonData["name"].(string)
	}
	if jsonData["password"] != nil {
		modelName.Publisher = jsonData["password"].(string)
	}
	if jsonData["year"] != nil {
		modelName.Year = int(jsonData["year"].(float64))
	}

	db.Table(table).Create(&modelName)
	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "success",
		"code":    http.StatusOK,
		"data":    modelName,
	})
}

func UpdateBooks(c echo.Context) error {
	result := map[string]interface{}{
		"message": "failed",
	}
	par := c.Param("id")
	intPar, _ := strconv.Atoi(par)
	if intPar > 0 {
		db := config.InitDB()
		modelName := models.Books{}
		db.Table(table).First(&modelName, intPar)
		if (modelName == models.Books{}) {
			return c.JSON(http.StatusOK, result)
		}

		var jsonData = make(map[string]interface{})
		err := c.Bind(&jsonData)
		if err != nil {
			fmt.Println(err)
			return c.JSON(http.StatusOK, result)

		}
		if jsonData["name"] != nil {
			modelName.Name = jsonData["name"].(string)
		}
		if jsonData["author"] != nil {
			modelName.Author = jsonData["author"].(string)
		}
		if jsonData["publisher"] != nil {
			modelName.Publisher = jsonData["publisher"].(string)
		}
		if jsonData["year"] != nil {
			modelName.Year = int(jsonData["year"].(float64))
		}
		saveResult := db.Table(table).Save(&modelName)

		if saveResult.RowsAffected > 0 {
			result["message"] = "success"
			result["code"] = http.StatusOK
			result["data"] = modelName
		}
	}

	return c.JSON(http.StatusOK, result)
}

func DeleteBooks(c echo.Context) error {
	result := map[string]interface{}{
		"message": "failed",
	}
	par := c.Param("id")
	modelName := models.Books{}
	u64, err := strconv.ParseUint(par, 10, 64)
	if err != nil {
		fmt.Println(err)
	}
	modelName.ID = uint(u64)
	if modelName.ID > 0 {
		db := config.InitDB()
		bookModel := models.Books{}
		saveResult := db.Table(table).Delete(&bookModel, uint(u64))
		if err != nil {
			return err
		}
		if saveResult.RowsAffected > 0 {
			result["message"] = "success"
			result["code"] = http.StatusOK
			result["data"] = modelName
		}
	}

	return c.JSON(http.StatusOK, result)
}
