package route

import (
	"UserAPI/util"
	"os"

	"github.com/dgrijalva/jwt-go"
	"github.com/gofiber/fiber/v2"
)

func refreshHandler(c *fiber.Ctx) error {
	refreshToken := c.Get("refreshToken")
	if refreshToken == "" {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"message": "Missing refresh token",
		})
	}
	token, err := jwt.Parse(refreshToken, func(t *jwt.Token) (interface{}, error) {
		if _, ok := t.Method.(*jwt.SigningMethodHMAC); !ok {
			return fiber.ErrUnauthorized, nil
		}
		return []byte(os.Getenv("Secret-Key")), nil
	})
	if err != nil || !token.Valid {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"error": "Invalid refresh token",
		})
	}
	claims, ok := token.Claims.(jwt.MapClaims)
	if !ok || !token.Valid {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"error": "Invalid refresh token claims",
		})
	}
	username := claims["Username"].(string)
	newAccessToken, err := util.CreateAccessToken(username)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": "Could not create access token",
		})
	}
	c.Set("accessToken", newAccessToken)
	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"message": "Refreshed access token",
	})
}
