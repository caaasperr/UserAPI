package route

import (
	"github.com/gofiber/fiber/v2"
)

func SetupRoute(app *fiber.App) {
	api := app.Group("/api")

	v1 := api.Group("/v1")

	v1.Get("/register", signinHandler)
	v1.Get("/login", loginHandler)
	v1.Get("/logout", logOutHandler)
	v1.Get("/refresh", refreshHandler)
	v1.Get("/delete", deleteHandler)
}
