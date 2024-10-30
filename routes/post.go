package routes

import (
	"github.com/gofiber/fiber/v2"
)

func PostRoutes(router fiber.Router) {
    router.Post("/create", func(c *fiber.Ctx) error {
        return c.SendString("respond with a resource")
    })
}