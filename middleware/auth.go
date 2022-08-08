package middleware

import (
	"os"
	jwtware "github.com/gofiber/jwt/v3"
	"github.com/gofiber/fiber/v2"
)

func Auth(app *fiber.App) {
	// https://github.com/gofiber/jwt

	/*
	JWT returns a JSON Web Token (JWT) auth middleware. 
	For valid token, it sets the user in Ctx.Locals and calls next handler. 
	For invalid token, it returns "401 - Unauthorized" error. For missing token, it returns "400 - Bad Request" error.
	*/
	app.Use("/api/*",jwtware.New(jwtware.Config{
		SigningKey: []byte(os.Getenv("secret")),
		ErrorHandler: jwtError, // Add custom error handler
	}))
}

func jwtError(c *fiber.Ctx, err error) error {
	if err.Error() == "Missing or malformed JWT" {
		return c.Status(fiber.StatusBadRequest).
			JSON(fiber.Map{"status": "error", "message": "Missing or malformed JWT", "data": nil})
	}
	return c.Status(fiber.StatusUnauthorized).
		JSON(fiber.Map{"status": "error", "message": "Invalid or expired JWT", "data": nil})
}