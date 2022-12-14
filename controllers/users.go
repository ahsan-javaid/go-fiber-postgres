package controllers

import (
	"books-crud/common"
	"books-crud/models"
	"books-crud/utils"

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

	r.DB.Where(&models.User{Email: email}).First(&user)

	if user.ID == 0 || utils.CheckPasswordHash(password, user.Password) == false {
		return common.Http401(context, "Invalid email or password")
	}

	resp := map[string]interface{}{"token": user.CreateToken(), "user": user}

	return common.Http200(context, resp)
}

func (r *Repository) Signup(context *fiber.Ctx) error {
	body := models.User{}

	err := context.BodyParser(&body)

	if err != nil {
		return common.Http400(context, err.Error())
	}

	errors := body.ValidateUser()

	if len(errors) > 0 {
		return common.Http400(context, errors)
	}

	user := &models.User{}

	r.DB.Where(&models.User{Email: body.Email}).First(&user)

	if user.ID != 0 {
		return common.Http400(context, "Email already exists")
	}

	body.Password, _ = utils.HashPassword(body.Password)

	result := r.DB.Create(&body)

	if result.Error != nil {
		return common.Http400(context, err.Error())
	}

	resp := map[string]interface{}{"token": body.CreateToken(), "user": body}

	return common.Http200(context, resp)
}

func (r *Repository) SetupUserRoutes(app *fiber.App) {
	// public routes
	app.Post("/users/login", r.Login)
	app.Post("/users/signup", r.Signup)
}
