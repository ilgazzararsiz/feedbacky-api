package controller

import (
	"context"
	"feedbacky-api/appconfig"
	"feedbacky-api/entity"
	"feedbacky-api/response"
	"net/http"
	"time"

	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

var feedbackCollection *mongo.Collection = appconfig.GetCollection(appconfig.DB, "feedbacky")
var validate = validator.New()

func CreateFeedback(c *fiber.Ctx) error {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	var feedback entity.Feedback
	defer cancel()

	if err := c.BodyParser(&feedback); err != nil {
			return c.Status(http.StatusBadRequest).JSON(response.ApiResponse{Status: http.StatusBadRequest, Message: "error", Data: &fiber.Map{"data": err.Error()}})
	}

	if validationErr := validate.Struct(&feedback); validationErr != nil {
			return c.Status(http.StatusBadRequest).JSON(response.ApiResponse{Status: http.StatusBadRequest, Message: "error", Data: &fiber.Map{"data": validationErr.Error()}})
	}

	newFeedback := entity.Feedback{
			Id:       			primitive.NewObjectID(),
			Content:  			feedback.Content,
			CreatedDate: 		feedback.CreatedDate,
			ApplicationId:	feedback.ApplicationId,
	}

	result, err := feedbackCollection.InsertOne(ctx, newFeedback)
	if err != nil {
			return c.Status(http.StatusInternalServerError).JSON(response.ApiResponse{Status: http.StatusInternalServerError, Message: "error", Data: &fiber.Map{"data": err.Error()}})
	}

	return c.Status(http.StatusCreated).JSON(response.ApiResponse{Status: http.StatusCreated, Message: "success", Data: &fiber.Map{"data": result}})
}