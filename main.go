package main

import (
	"log"
	"os"
	"os/signal"
	"syscall"
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

	// Listen from a different goroutine
	go func() {
		if err := app.Listen(fmt.Sprintf(":%v", config.Port)); err != nil {
			log.Panic(err)
		}
	}()

	c := make(chan os.Signal, 1)   // Create channel to signify a signal being sent
	signal.Notify(c, os.Interrupt, syscall.SIGTERM) // When an interrupt or termination signal is sent, notify the channel

	_ = <-c // This blocks the main thread until an interrupt is received
	fmt.Println("Gracefully shutting down...")
	_ = app.Shutdown()

	fmt.Println("Running cleanup tasks...")

	// cleanup tasks go here
	// db.Close()
	dbInstance, _ := db.DB()
	dbInstance.Close()
	fmt.Println("Fiber was successful shutdown.")

}
