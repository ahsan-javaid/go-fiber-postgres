package main

import (
	"log"
	"os"
	"fmt"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/gofiber/fiber/v2/middleware/recover"
	"github.com/joho/godotenv"
	"books-crud/middleware"

	"books-crud/models"
	"books-crud/storage"
	"books-crud/controllers"
)

func main() {

	err := godotenv.Load(".env")

	if err != nil {
		log.Fatal(err)
	}

	app := fiber.New()

	app.Use(recover.New())
	app.Use(logger.New(logger.Config{
		Format: "[${time}] ${ip}  ${status} - ${latency} ${method} ${path}\n",
	}))


	config := &storage.Config{
		Port:     os.Getenv("PORT"),
		DBHost:     os.Getenv("DB_HOST"),
		DBPort:     os.Getenv("DB_PORT"),
		DBPassword: os.Getenv("DB_PASSWORD"),
		DBUser:     os.Getenv("DB_USER"),
		SSLMode:  os.Getenv("SSL"),
		DBName:   os.Getenv("DB_NAME"),
	}

	db, err := storage.NewConnection(config)

	if err != nil {
		log.Fatal(err)
	}

	err = models.MigrateBooks(db)

	if err != nil {
		log.Fatal(err)
	}

	err = models.MigrateUsers(db)

	if err != nil {
		log.Fatal(err)
	}

	r := controllers.Repository{
		DB: db,
	}

	middleware.Auth(app)
	r.SetupBookRoutes(app)
	r.SetupUserRoutes(app)

	app.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("Ok")
	})

	app.Listen(fmt.Sprintf(":%v", config.Port))


}
