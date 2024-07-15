package main

import (
	"UserAPI/database"
	"UserAPI/route"
	"UserAPI/util"

	"github.com/gofiber/fiber/v2"
)

func main() {
	util.InitializeEnv()
	database.Initialize()
	app := fiber.New()
	route.SetupRoute(app)
	app.Listen(":8080")
}
