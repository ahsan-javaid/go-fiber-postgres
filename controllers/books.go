package controllers

import (
	"books-crud/common"
	"books-crud/models"
	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

type Repository struct {
	DB *gorm.DB
}

func (r *Repository) GetAllBook(context *fiber.Ctx) error {
	bookModels := &[]models.Book{}

	err := r.DB.Find(bookModels).Error

	if err != nil {
		return common.Http400(context, err.Error())
	}

	return common.Http200(context, bookModels)
}

func (r *Repository) CreateBook(context *fiber.Ctx) error {
	book := models.Book{}

	err := context.BodyParser(&book)

	if err != nil {
		return common.Http400(context, err.Error())
	}

	result := r.DB.Create(&book)

	if result.Error != nil {
		return common.Http400(context, err.Error())
	}

	return common.Http200(context, book)
}

func (r *Repository) DeleteBook(context *fiber.Ctx) error {
	bookModel := models.Book{}
	id := context.Params("id")
	if id == "" {
		return common.Http400(context, "id cannot be empty")
	}

	err := r.DB.Delete(bookModel, id)

	if err.Error != nil {
		return common.Http400(context, err.Error)
	}

	return common.Http200(context, "book deleted")
}

func (r *Repository) GetBookByID(context *fiber.Ctx) error {

	id := context.Params("id")
	bookModel := &models.Book{}

	if id == "" {
		return common.Http400(context, "id is required")
	}

	err := r.DB.Where("id = ?", id).First(bookModel).Error
	if err != nil {
		return common.Http400(context, err.Error())
	}

	return common.Http200(context, bookModel)
}

func (r *Repository) SetupBookRoutes(app *fiber.App) {
	api := app.Group("/api")
	api.Post("/books", r.CreateBook)
	api.Get("/books", r.GetAllBook)
	api.Delete("/books/:id", r.DeleteBook)
	api.Get("/books/:id", r.GetBookByID)
}
