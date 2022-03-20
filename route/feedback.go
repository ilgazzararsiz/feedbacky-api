package route

import (
	"feedbacky-api/controller"
	"github.com/gofiber/fiber/v2"
)

func Feedback(app *fiber.App) {
	app.Post("/CreateFeedback", controller.CreateFeedback)
	app.Get("/GetFeedback/:feedbackId", controller.GetFeedback)
}
