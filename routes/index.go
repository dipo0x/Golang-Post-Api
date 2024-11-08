package routes

import (
	"github.com/gofiber/fiber/v2"
    "go-rest-api/controllers"
)

func IndexRoutes(router fiber.Router) {
    router.Get("/", controller.Index)
}