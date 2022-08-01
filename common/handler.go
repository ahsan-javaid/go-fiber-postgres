package common

import (
	"net/http"
	"github.com/gofiber/fiber/v2"
)

func Http200 (context *fiber.Ctx, data interface{}) error {
	context.Status(http.StatusOK).JSON(&fiber.Map{"msg": "success", "data": data})
	return nil
}

func Http400 (context *fiber.Ctx, data interface{}) error {
	context.Status(http.StatusBadRequest).JSON(&fiber.Map{"msg": data})
	return nil
}

func Http401 (context *fiber.Ctx, data interface{}) error {
	context.Status(http.StatusUnauthorized).JSON(&fiber.Map{"msg": data})
	return nil
}