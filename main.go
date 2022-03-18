package main

import (
    "feedbacky-api/appconfig"
    "feedbacky-api/route"
    "github.com/gofiber/fiber/v2" 
)

func main() {
    app := fiber.New()

    //run database
    appconfig.ConnectDB()

    //routes
    route.Feedback(app)

    app.Listen(":6000")
}