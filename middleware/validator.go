package middleware

import  (
	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
	"go-rest-api/types"
)

var Validator = validator.New()
func ValidateCreatePost(c *fiber.Ctx) error {
	var errors []*types.IError
	body := new(types.IPost)

	if err := c.BodyParser(&body); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"status": 400,
			"success": false,
			"error": "Invalid request payload",
		})
	}

	err := Validator.Struct(body)
	if err != nil {
		for _, err := range err.(validator.ValidationErrors) {
			var el types.IError
			el.Field = err.Field()
			el.Tag = err.Tag()
			el.Value = err.Param()
			errors = append(errors, &el)
		}
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{ "status": 400,
			"success": false,
			"error": errors,
		})
	}
	return c.Next()
}