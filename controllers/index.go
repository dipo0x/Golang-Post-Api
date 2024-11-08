package controller

import (
	"github.com/gofiber/fiber/v2"
)

func Index(c *fiber.Ctx) error {
    return c.Status(fiber.StatusOK).JSON(fiber.Map{
        "status": 200,
        "success": true,
        "data": "Go rest api server is up and running ",
    })
}