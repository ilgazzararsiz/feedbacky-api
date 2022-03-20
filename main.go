package main

import (
	"feedbacky-api/appconfig"
	"feedbacky-api/route"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
)

func main() {
	app := fiber.New()

	app.Use(cors.New())

	//run database
	appconfig.ConnectDB()

	//routes
	route.Feedback(app)

	app.Listen(":6001")
}
