package routes

import (
	"github.com/gofiber/fiber/v2"
)

func healthzHandler(c *fiber.Ctx) error {
	return c.SendStatus(fiber.StatusOK)
}
