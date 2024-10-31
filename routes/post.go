package routes

import (
	"github.com/gofiber/fiber/v2"
    "go-rest-api/controllers"
    "go-rest-api/middleware"
)

func PostRoutes(router fiber.Router) {
    router.Post("/create", middleware.ValidateCreatePost, controller.CreatePost)
    router.Get("/create", controller.CreatePost)
}