package route

import (
	"feedbacky-api/controller"
	"github.com/gofiber/fiber/v2"
)

func Feedback(app *fiber.App) {
	app.Post("/feedback", controller.CreateFeedback)
	app.Get("/feedback/:feedbackId", controller.GetFeedback)
}
