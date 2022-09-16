package middleware

import (
	"CourseAlchemy/config"
	"CourseAlchemy/models"
	"fmt"
	"os"
	"strings"
	"time"

	"github.com/golang-jwt/jwt"
	"github.com/labstack/echo/v4"
)

func JWTAuth(c echo.Context) (bool, error) {
	payload := ExtractTokenPayload(c)
	if payload != nil {
		modelStruct := models.User{}
		db := config.InitDB()
		err := db.Table("users").Where(models.User{Email: payload["email"].(string)}).First(&modelStruct).Error
		if err == nil {
			return true, nil
		}
	}
	return false, nil
}

func CreateToken(params map[string]interface{}) (string, error) {
	claim := jwt.MapClaims{} // JWT payload
	claim["authorized"] = true
	if params["email"] != nil {
		claim["email"] = params["email"]
	}
	if params["name"] != nil {
		claim["name"] = params["name"]
	}
	claim["exp"] = time.Now().Add(time.Hour * 1).Unix() // Token exp 1 hour
	tokenPointer := jwt.NewWithClaims(jwt.SigningMethodHS256, claim)
	fmt.Println(claim)
	fmt.Println(tokenPointer)
	return tokenPointer.SignedString([]byte(os.Getenv("JWT_SECRET")))
}

func ExtractTokenPayload(e echo.Context) map[string]interface{} {
	reqToken := e.Request().Header.Get("Authorization")
	splitToken := strings.Split(reqToken, "Bearer ")
	reqToken = splitToken[1]

	token, err := jwt.Parse(reqToken, func(token *jwt.Token) (interface{}, error) {
		// check token signing method etc
		return []byte(os.Getenv("JWT_SECRET")), nil
	})
	if err == nil {
		if token.Valid {
			return token.Claims.(jwt.MapClaims)
		}
	}
	return nil
}
