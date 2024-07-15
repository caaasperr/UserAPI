package route

import (
	"UserAPI/controller"
	"UserAPI/util"

	"github.com/gofiber/fiber/v2"
)

func deleteHandler(c *fiber.Ctx) error {
	accessToken := c.Get("accessToken")
	if util.VerifyAccessToken(accessToken) {
		controller.DeleteUser(util.GetUsernameFromToken(accessToken))
		c.Set("refreshToken", "")
		c.Set("accessToken", "")
		return c.Status(fiber.StatusOK).JSON(fiber.Map{
			"message": "Deleted success",
		})
	} else {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": "Couldn't verify access token",
		})
	}
}
