package main

import (
	"CourseAlchemy/config"
	"CourseAlchemy/models"
	"CourseAlchemy/routes"

	"github.com/joho/godotenv"
)

func main() {
	godotenv.Load(".env")
	config.InitDB().AutoMigrate(models.User{}, models.Books{})
	e := routes.New()
	e.Logger.Fatal(e.Start(":8000"))
}
