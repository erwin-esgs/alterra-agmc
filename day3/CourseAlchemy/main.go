package main

import (
	"CourseAlchemy/config"
	m "CourseAlchemy/middleware"
	"CourseAlchemy/models"
	"CourseAlchemy/routes"

	"github.com/joho/godotenv"
)

func main() {
	// ex, err := os.Executable()
	// if err != nil {
	// 	panic(fmt.Sprintf("Failed to get executable path: %s", err))
	// }
	// envPath := filepath.Join(filepath.Dir(ex), ".env")
	// err = godotenv.Load(envPath)
	// fmt.Println("main", os.Getenv("JWT_SECRET"))
	godotenv.Load(".env")
	config.InitDB().AutoMigrate(models.User{}, models.Books{})

	e := routes.New()
	m.LogMiddleware(e)
	e.Logger.Fatal(e.Start(":8000"))
}
