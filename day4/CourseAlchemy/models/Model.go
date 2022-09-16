package models

import (
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Name     string `json:"name"`
	Email    string `json:"email"`
	Password string `json:"password"`
}
type Books struct {
	gorm.Model
	Name      string `json:"name"`
	Author    string `json:"author"`
	Publisher string `json:"publisher"`
	Year      int    `json:"year"`
}
