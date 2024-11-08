package routes

import (
	"github.com/gofiber/fiber/v2"
    "go-rest-api/controllers"
    "go-rest-api/middleware"
)

func PostRoutes(router fiber.Router) {
    router.Post("/create", middleware.ValidateCreatePost, controller.CreatePost)
    router.Get("/get/", controller.GetPosts)
    router.Patch("/update/:id", controller.EditPost)
    router.Delete("/delete/:id", controller.DeletePost)
}