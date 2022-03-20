package response

import "github.com/gofiber/fiber/v2"

type CreateFeedbackResponse struct {
	Status  int        `json:"status"`
	Message string     `json:"message"`
	Data    *fiber.Map `json:"data"`
}
