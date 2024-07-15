package route

import (
	"UserAPI/controller"
	"UserAPI/model"

	"github.com/gofiber/fiber/v2"
)

func signinHandler(c *fiber.Ctx) error {
	var user *model.User
	if err := c.BodyParser(&user); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": "Invalid request body",
		})
	}
	err := controller.CreateUser(user.Username, user.Password)
	if !err {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": "Username is duplicated",
		})
	} else {
		return c.Status(fiber.StatusOK).JSON(fiber.Map{
			"Username": user.Username,
			"message":  "Signin successful",
		})
	}
}
