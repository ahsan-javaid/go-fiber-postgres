package models

import (
	"os"
	"time"

	"github.com/golang-jwt/jwt/v4"
	"gorm.io/gorm"
	"github.com/go-playground/validator/v10"
)

type User struct {
	ID       uint   `gorm:"primary key; autoIncrement" json:"id"`
	Name     string `json:"name" validate:"required"`
	Email    string `json:"email" validate:"required,email"`
	Password string `json:"password" validate:"required,min=6"`
}

type ErrorResponse struct {
	FailedField string
	Tag         string
	Value       string
	Reason      string
}

var validate = validator.New()

func (user *User) ValidateUser() []*ErrorResponse {
	var errors []*ErrorResponse
	err := validate.Struct(user)
	if err != nil {
		for _, err := range err.(validator.ValidationErrors) {
			var element ErrorResponse
			element.FailedField = err.StructNamespace()
			element.Tag = err.Tag()
			element.Value = err.Param()
			element.Reason = err.Error()
			errors = append(errors, &element)
		}
	}
	return errors
}

func (user *User) CreateToken() string {
	// Create the Claims
	claims := jwt.MapClaims{
		"name":  user.Name,
		"email": user.Email,
		"id":    user.ID,
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
