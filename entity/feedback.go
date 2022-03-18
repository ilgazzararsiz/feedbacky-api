package entity

import(
 "go.mongodb.org/mongo-driver/bson/primitive"
 "time"
)

type Feedback struct {
	Id 							primitive.ObjectID 	`json:"id,omitempty"`
	Content 				string 							`json:"content,omitempty" validate:"required"`
	CreatedDate 		*time.Time 					`json:"createdDate,omitempty"`
	ApplicationId		string 							`json:"applicationId,omitempty"`
}
