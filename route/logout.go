package route

import (
	"UserAPI/util"

	"github.com/gofiber/fiber/v2"
)

func logOutHandler(c *fiber.Ctx) error {
	if util.VerifyAccessToken(c.Get("accessToken")) {
		c.Set("refreshToken", "")
		c.Set("accessToken", "")
		return c.Status(fiber.StatusOK).JSON(fiber.Map{
			"message": "Logout success",
		})
	} else {
		return c.Status(fiber.StatusNotAcceptable).JSON(fiber.Map{
			"message": "Couldn't verify access token",
		})
	}
}
