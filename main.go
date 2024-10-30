package main

import (
	"github.com/gofiber/fiber/v2"
	// "go-rest-api/modules/post/post.route"
)

func main () {

	app := fiber.New()
	app.Get("/testApi", func(ctx *fiber.Ctx) error {
		return ctx.Status(200).JSON(fiber.Map{
			"success": true,
			"message": "Go fiber first app",
		})
	})
	app.Listen(":3000")
}