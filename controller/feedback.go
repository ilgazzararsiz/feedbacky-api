package controller

import (
	"feedbacky-api/service"
	"github.com/gofiber/fiber/v2"
)

func CreateFeedback(c *fiber.Ctx) error {
	return service.CreateFeedback(c)
}

func GetFeedback(c *fiber.Ctx) error {
	return service.GetFeedback(c)
}
