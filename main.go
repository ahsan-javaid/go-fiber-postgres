package main

import (
	"log"
	"os"

	"github.com/gofiber/fiber/v2"
	"github.com/joho/godotenv"

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

	config := &storage.Config{
		Host:     os.Getenv("host"),
		Port:     os.Getenv("port"),
		Password: os.Getenv("pass"),
		User:     os.Getenv("user"),
		SSLMode:  os.Getenv("ssl"),
		DBName:   os.Getenv("db"),
	}

	db, err := storage.NewConnection(config)

	if err != nil {
		log.Fatal(err)
	}

	err = models.MigrateBooks(db)

	if err != nil {
		log.Fatal(err)
	}

	r := controllers.Repository{
		DB: db,
	}

	r.SetupBookRoutes(app)

	app.Listen(":8080")

}