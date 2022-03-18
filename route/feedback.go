package route

import (
	"github.com/gofiber/fiber/v2"
	"feedbacky-api/controller"
)

func Feedback(app *fiber.App) {
	app.Post("/CreateFeedback", controller.CreateFeedback)
}