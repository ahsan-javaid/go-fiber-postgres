package models

import (
	"os"
	"time"
	"gorm.io/gorm"
	"github.com/golang-jwt/jwt/v4"
)


type User struct {
	ID          uint    `gorm:"primary key; autoIncrement" json:"id"`
	Name  			string  `json:"name"`
	Email       string  `json:"email"`
	Password    string  `json:"password"`
}

func (user *User) CreateToken() string {
	// Create the Claims
	claims := jwt.MapClaims{
		"name":  user.Name,
		"email": user.Email,
		"id": user.ID,
		"exp":   time.Now().Add(time.Hour * 72).Unix(),
	}

	// Create token
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	signed, _ := token.SignedString([]byte(os.Getenv("secret")))

	return signed
}

func MigrateUsers(db *gorm.DB) error {
	err := db.AutoMigrate(&User{})
	return err
}