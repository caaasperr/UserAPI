package route

import (
	"UserAPI/controller"
	"UserAPI/model"
	"UserAPI/util"

	"github.com/gofiber/fiber/v2"
)

func loginHandler(c *fiber.Ctx) error {
	var user model.User
	if err := c.BodyParser(&user); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": "Invalid request body",
		})
	}
	if util.ComparePassword(user.Password, controller.GetUserById(user.Username).Password) {
		refreshToken, err := util.CreateRefreshToken(user.Username)
		if err != nil {
			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
				"message": "Could not create refresh token",
			})
		}
		c.Set("refreshToken", refreshToken)
		accessToken, err := util.CreateAccessToken(user.Username)
		if err != nil {
			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
				"message": "Could not create access token",
			})
		}
		c.Set("accessToken", accessToken)
		return c.Status(fiber.StatusOK).JSON(fiber.Map{
			"message": "Login success",
		})
	} else {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": "Incorrect username or password",
		})
	}
}
