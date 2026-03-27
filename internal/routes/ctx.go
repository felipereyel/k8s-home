package routes

import (
	"scaler/internal/services"

	"github.com/gofiber/fiber/v2"
)

type controllerHandler func(svcs *services.Services, c *fiber.Ctx) error

func bind(svcs *services.Services, handler controllerHandler) fiber.Handler {
	return func(c *fiber.Ctx) error {
		return handler(svcs, c)
	}
}
