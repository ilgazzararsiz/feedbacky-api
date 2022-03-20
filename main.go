package main

import (
	"feedbacky-api/appconfig"
	"feedbacky-api/route"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"os"
)

func main() {
	app := fiber.New()

	app.Use(cors.New())

	appconfig.ConnectDB()

	route.Feedback(app)

	port := os.Getenv("PORT")
	if port == "" {
		port = "5000"
	}

	app.Listen(":" + port)
}
