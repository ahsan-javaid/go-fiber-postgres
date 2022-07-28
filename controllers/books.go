package controllers

import (
	"books-crud/models"
	"fmt"
	"net/http"

	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

type Repository struct {
	DB *gorm.DB
}

type Book struct {
	Author    string `json:"author"`
	Title     string `json:"title"`
	Publisher string `json:"publisher"`
}

func (r *Repository) GetAllBook(context *fiber.Ctx) error {
	bookModels := &[]models.Books{}

	err := r.DB.Find(bookModels).Error

	if err != nil {
		context.Status(http.StatusBadRequest).JSON(&fiber.Map{"msg": "req failed"})

		return err
	}

	context.Status(http.StatusOK).JSON(&fiber.Map{"msg": "success", "data": bookModels})

	return nil
}

func (r *Repository) CreateBook(context *fiber.Ctx) error {
	book := Book{}

	err := context.BodyParser(&book)

	if err != nil {
		context.Status(http.StatusUnprocessableEntity).JSON(&fiber.Map{"msg": "req failer"})

		return err
	}

	err = r.DB.Create(&book).Error

	if err != nil {
		context.Status(http.StatusBadRequest).JSON(&fiber.Map{"msg": "cannot create book"})

		return err
	}

	context.Status(http.StatusOK).JSON(&fiber.Map{"msg": "book created"})

	return nil
}

func (r *Repository) DeleteBook(context *fiber.Ctx) error {
	bookModel := models.Books{}
	id := context.Params("id")
	if id == "" {
		context.Status(http.StatusInternalServerError).JSON(&fiber.Map{
			"message": "id cannot be empty",
		})
		return nil
	}

	err := r.DB.Delete(bookModel, id)

	if err.Error != nil {
		context.Status(http.StatusBadRequest).JSON(&fiber.Map{
			"message": "could not delete book",
		})
		return err.Error
	}
	context.Status(http.StatusOK).JSON(&fiber.Map{
		"message": "book delete successfully",
	})
	return nil
}

func (r *Repository) GetBookByID(context *fiber.Ctx) error {

	id := context.Params("id")
	bookModel := &models.Books{}
	if id == "" {
		context.Status(http.StatusInternalServerError).JSON(&fiber.Map{
			"message": "id cannot be empty",
		})
		return nil
	}

	fmt.Println("the ID is", id)

	err := r.DB.Where("id = ?", id).First(bookModel).Error
	if err != nil {
		context.Status(http.StatusBadRequest).JSON(
			&fiber.Map{"message": "could not get the book"})
		return err
	}
	context.Status(http.StatusOK).JSON(&fiber.Map{
		"message": "book id fetched successfully",
		"data":    bookModel,
	})
	return nil
}

func (r *Repository) SetupBookRoutes(app *fiber.App) {
	api := app.Group("/api")
	api.Post("/books", r.CreateBook)
	api.Get("/books", r.GetAllBook)
	api.Delete("/books/:id", r.DeleteBook)
	api.Get("/books/:id", r.GetBookByID)
}
