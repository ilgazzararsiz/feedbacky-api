package service

import (
	"context"
	"feedbacky-api/appconfig"
	"feedbacky-api/entity"
	"feedbacky-api/response"
	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"net/http"
	"strconv"
	"time"
	"unicode/utf8"
)

const MAX_FEEDBACK_CONTENT_LENGTH = 2000
const DATE_LAYOUT = "02-01-2006 15:04:05"

var feedbackCollection = appconfig.GetCollection(appconfig.DB, "feedbacky")
var validate = validator.New()

func CreateFeedback(c *fiber.Ctx) error {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	var feedback entity.Feedback
	defer cancel()

	if err := c.BodyParser(&feedback); err != nil {
		return c.Status(http.StatusBadRequest).JSON(response.CreateFeedbackResponse{Status: http.StatusBadRequest, Message: "error", Data: &fiber.Map{"data": err.Error()}})
	}

	if validationErr := validate.Struct(&feedback); validationErr != nil {
		return c.Status(http.StatusBadRequest).JSON(response.CreateFeedbackResponse{Status: http.StatusBadRequest, Message: "error", Data: &fiber.Map{"data": validationErr.Error()}})
	}
	contentLength := utf8.RuneCountInString(feedback.Content)

	if contentLength > MAX_FEEDBACK_CONTENT_LENGTH {
		return c.Status(http.StatusBadRequest).JSON(response.CreateFeedbackResponse{Status: http.StatusBadRequest, Message: "Content length must lower than " + strconv.Itoa(MAX_FEEDBACK_CONTENT_LENGTH) + "."})
	}

	newFeedback := entity.Feedback{
		Content:       feedback.Content,
		CreatedDate:   time.Now(),
		ApplicationId: feedback.ApplicationId,
	}

	result, err := feedbackCollection.InsertOne(ctx, newFeedback)
	if err != nil {
		return c.Status(http.StatusInternalServerError).JSON(response.CreateFeedbackResponse{Status: http.StatusInternalServerError, Message: "error", Data: &fiber.Map{"data": err.Error()}})
	}

	return c.Status(http.StatusCreated).JSON(response.CreateFeedbackResponse{Status: http.StatusCreated, Message: "success", Data: &fiber.Map{"data": result}})
}

func GetFeedback(c *fiber.Ctx) error {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	feedbackId := c.Params("feedbackId")
	var feedback entity.Feedback
	defer cancel()

	objectId, _ := primitive.ObjectIDFromHex(feedbackId)

	err := feedbackCollection.FindOne(ctx, bson.M{"_id": objectId}).Decode(&feedback)
	if err != nil {
		return c.Status(http.StatusInternalServerError).JSON(response.CreateFeedbackResponse{Status: http.StatusInternalServerError, Message: "error", Data: &fiber.Map{"data": err.Error()}})
	}

	return c.Status(http.StatusOK).JSON(response.CreateFeedbackResponse{Status: http.StatusOK, Message: "success", Data: &fiber.Map{"data": feedback}})
}
