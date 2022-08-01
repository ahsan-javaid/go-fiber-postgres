package controllers

import (
	"books-crud/common"
	"books-crud/models"

	"github.com/gofiber/fiber/v2"
)

func (r *Repository) Login(context *fiber.Ctx) error {
	body := models.User{}

	err := context.BodyParser(&body)

	if err != nil {
		return common.Http400(context, err.Error())
	}

	email := body.Email
	password := body.Password

	user := &models.User{}

	r.DB.Where(&models.User{Email: email, Password: password}).First(&user)

	if user.ID == 0 {
		return common.Http401(context, "Invalid email or password")
	}

	resp := map[string]interface{}{"token": user.CreateToken(), "user": user}

	return common.Http200(context, resp)
}

func (r *Repository) SetupUserRoutes(app *fiber.App) {
	api := app.Group("/api")
	api.Post("/users/login", r.Login)
}
